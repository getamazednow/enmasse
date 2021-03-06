/*
 * Copyright 2016 Red Hat Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
'use strict';

var util = require("util");
var events = require("events");
var qdr = require("./qdr.js");
var myutils = require("./utils.js");
var router_config = require("./router_config.js");
var log = require("./log.js").logger();

const ID_QUALIFIER = 'ragent'

function has_inter_router_role (record) {
    return record.role === 'inter-router';
}

function matches_qualifier (record) {
    return record.name && record.name.indexOf(ID_QUALIFIER) === 0;
}

function to_host_port (record) {
    return record.host + ':' + record.port;
}

/**
 * A KnownRouter instance represents routers this process knows about
 * but is not directly connected to (or responsible for).
 */
var KnownRouter = function (container_id, listeners) {
    this.container_id = container_id;
    this.listeners = listeners;
};

/**
 * A ConnectedRouter represents a router this process is connected to
 * and is therefore resonsible for configuring.
 */
var ConnectedRouter = function (connection) {
    events.EventEmitter.call(this);
    this.container_id = connection.container_id;
    this.listeners = undefined;
    this.connectors = undefined;
    this.fully_connected = false;
    this.addresses = {};
    this.initial_provisioning_completed = false;
    this.addresses_synchronized = false;
    this.realise_address_definitions = myutils.serialize(this._realise_address_definitions.bind(this));

    this.router_mgmt = new qdr.Router(connection);
    this.router_mgmt.name = 'mgmt';
    var self = this;
    connection.on('receiver_open', function () {
        self.emit('ready', self);
    });
};

util.inherits(ConnectedRouter, events.EventEmitter);

function remove_current(set) {
    if (this.listeners) {
        this.listeners.forEach(function (l) { delete set[l]; });
    }
}

function has_listener(host_port) {
    return this.listeners && this.listeners.indexOf(host_port) !== -1;
}

ConnectedRouter.prototype.has_connector_to = function (router) {
    return this.connectors && this.connectors.some(has_listener.bind(router));
};

ConnectedRouter.prototype.expects_connector_to = function (router) {
    return router.container_id < this.container_id && router.listeners && router.listeners.length > 0;
};

ConnectedRouter.prototype.is_missing_connector_to = function (router) {
    return this.expects_connector_to(router) && !this.has_connector_to(router);
};

ConnectedRouter.prototype.is_ready_for_connectivity_check = function () {
    return this.initial_provisioning_completed && this.connectors !== undefined;
}

ConnectedRouter.prototype.check_connectors = function (routers) {
    var missing = [];
    var stale = myutils.index(this.connectors);
    for (var r in routers) {
        var router = routers[r];
        if (router === this) continue;
        if (this.is_missing_connector_to(router)) {
            missing.push(router.listeners[0]);
        }
        remove_current.call(router, stale);
    }
    stale = Object.keys(stale);

    var num_connectors = 1;
    if (process.env.ROUTER_NUM_CONNECTORS) {
        num_connectors = process.env.ROUTER_NUM_CONNECTORS;
    }
    var do_create = this.forall_connectors.bind(this, num_connectors, this.create_connector.bind(this));
    var do_delete = this.forall_connectors.bind(this, num_connectors, this.delete_connector.bind(this));
    var work = missing.map(do_create).concat(stale.filter(matches_qualifier).map(do_delete));
    var self = this;
    if (work.length) {
        this.fully_connected = false;
        log.info('[%s] checking connectors on router, missing=%j, stale=%j', this.container_id, missing, stale);
        //if made changes, requery when they are complete
        Promise.all(work).then(function () {
            self.retrieve_connectors();
        }).catch(function (error) {
            log.warn('[%s] error on updating connectors: %s', self.container_id, error);
            self.retrieve_connectors();
        });
        //prevent any updates to connectors until we have re-retrieved
        //them from router after updates:
        this.connectors = undefined;
    } else {
        log.info('[%s] fully connected', this.container_id);
        this.fully_connected = true;
        this.emit('synchronized');
    }
};

ConnectedRouter.prototype.verify_addresses = function (expected) {
    if (!expected || this.actual === undefined) {
        return false;
    }

    for (var i = 0; i < expected.length; i++) {
        var address = expected[i];
        if (address["store_and_forward"] && !address["multicast"]) {
            if (this.actual[address.name] === undefined) {
                return false;
            }
        }
    }
    return true;
}

ConnectedRouter.prototype.forall_connectors = function (num_connectors, connector_operation, host_port) {
    var futures = [];
    for (var i = 0; i < num_connectors; i++) {
        var connector_name = ID_QUALIFIER + '-' + host_port + '-' + i;
        futures.push(connector_operation(host_port, connector_name));
    }
    return Promise.all(futures);
}

ConnectedRouter.prototype.create_connector = function (host_port, connector_name) {
    log.info('[%s] creating connector %s to %s', this.container_id, connector_name, host_port);
    var parts = host_port.split(':');
    return this.create_entity('connector', connector_name, {role:'inter-router', host:parts[0], port:parts[1],
                                                            sslProfile:'ssl_internal_details', verifyHostName:'no'});
};

ConnectedRouter.prototype.delete_connector = function (host_port, connector_name) {
    log.info('[%s] deleting connector %s to %s', this.container_id, connector_name, host_port);
    return this.delete_entity('connector', connector_name);
};

ConnectedRouter.prototype.retrieve_listeners = function () {
    var self = this;
    return this.query('listener', {attributeNames:['identity', 'name', 'host', 'port', 'role']}).then(function (results) {
        self.listeners = results.filter(has_inter_router_role).map(to_host_port);
        log.debug('[%s] retrieved listeners: %j', self.container_id, self.listeners);
        self.emit('listeners_updated', self);
    }).catch(function (error) {
        log.warn('[%s] failed to retrieve listeners: %s', self.container_id, error);
        return self.retrieve_listeners();
    });
};

ConnectedRouter.prototype.retrieve_connectors = function () {
    var self = this;
    return this.query('connector', {attributeNames:['identity', 'name', 'host', 'port', 'role']}).then(function (results) {
        self.connectors = results.filter(has_inter_router_role).map(to_host_port);
        log.debug('[%s] retrieved connectors: %j', self.container_id, self.connectors);
        self.emit('connectors_updated', self);
    }).catch(function (error) {
        log.warn('[%s] failed to retrieve connectors: %s', self.container_id, error);
        return self.retrieve_connectors();
    });
};

ConnectedRouter.prototype.query = function (type, options) {
    return this.router_mgmt.query(type, options);
};

ConnectedRouter.prototype.create_entity = function (type, name, attributes) {
    return this.router_mgmt.create_entity(type, name, attributes);
};

ConnectedRouter.prototype.delete_entity = function (type, name) {
    return this.router_mgmt.delete_entity(type, name);
};

ConnectedRouter.prototype.sync_addresses = function (desired) {
    this.desired = desired;
    this.realise_address_definitions();
};

ConnectedRouter.prototype._realise_address_definitions = function () {
    var self = this;
    return router_config.realise_address_definitions(this.desired, this.router_mgmt).then(function (result) {
        self.actual = result;
        log.info('[%s] addresses synchronized', self.container_id);
        if (self.initial_provisioning_completed !== true) {
            self.initial_provisioning_completed = true;
            self.emit('provisioned', self);
        }
        self.emit('synchronized', self);
    }).catch(function (error) {
        console.error('[%s] error while synchronizing addresses: %s', self.container_id, error);
        log.error('[%s] error while synchronizing addresses: %s', self.container_id, error);
    });
};

ConnectedRouter.prototype.is_synchronized = function () {
    if (this.actual === undefined) return false;
    for (var k in this.desired) {
        var desired = this.desired[k];
        var actual = this.actual[desired.address];
        if (actual === undefined) {
            log.info('[%s] not synchronized, missing %s %s', this.container_id, desired.type, desired.address);
            return false;
        } else if (actual.type !== desired.type) {
            log.info('[%s] not synchronized,  %s of wrong type expected %s got %s', this.container_id, desired.address, desired.type, actual.type);
            return false;
        }
    }
    if (Object.keys(this.actual).length !== this.desired.length) {
        log.info('[%s] not synchronized, have extra addresses', this.container_id);
        return false;
    }
    return true;
}

module.exports = {
    connected: function (conn) { return new ConnectedRouter(conn); },
    known: function (container_id, listeners) { return new KnownRouter(container_id, listeners); }
};
