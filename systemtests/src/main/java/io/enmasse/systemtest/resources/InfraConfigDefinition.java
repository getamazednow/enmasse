/*
 * Copyright 2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */
package io.enmasse.systemtest.resources;

import java.util.ArrayList;
import java.util.List;

import io.enmasse.systemtest.AddressSpaceType;
import io.vertx.core.json.JsonObject;

public class InfraConfigDefinition {

    private String name;
    private AddressSpaceType type;
    private List<InfraSpecComponent> infraComponents;
    private String version;

    public InfraConfigDefinition(String name, AddressSpaceType type, List<InfraSpecComponent> infraComponents, String version) {
        this.name = name;
        this.type = type;
        this.infraComponents = infraComponents;
        this.version = version;
    }

    public String getName() {
        return name;
    }

    public AddressSpaceType getType() {
        return type;
    }

    public List<InfraSpecComponent> getAddressResources() {
        return infraComponents;
    }

    public String getVersion() {
        return version;
    }

    public JsonObject toJson() {
        JsonObject config = new JsonObject();
        config.put("apiVersion", "admin.enmasse.io/v1beta1");
        config.put("kind", type.equals(AddressSpaceType.STANDARD) ? "StandardInfraConfig" : "BrokeredInfraConfig");

        JsonObject definitionMetadata = new JsonObject(); // <metadata>
        definitionMetadata.put("name", this.getName());
        config.put("metadata", definitionMetadata);// </metadata>

        JsonObject spec = new JsonObject(); // <spec>
        for (InfraSpecComponent res : this.getAddressResources()) {
            JsonObject component = res.toJson();
            spec.put(res.getType(), component);
        }
        spec.put("version", version);
        config.put("spec", spec); // </requiredResources>
        return config;
    }


    public static InfraConfigDefinition fromJson(JsonObject infraDefinition) {
        JsonObject metadataDef = infraDefinition.getJsonObject("metadata");

        JsonObject spec = infraDefinition.getJsonObject("spec");
        List<InfraSpecComponent> components = new ArrayList<>();

        spec.stream().forEach(entry -> {
            InfraSpecComponent component = null;
            switch (entry.getKey()) {
                case InfraSpecComponent.ADMIN_INFRA_RESOURCE:
                    component = AdminInfraSpec.fromJson((JsonObject) entry.getValue());
                    break;
                case InfraSpecComponent.BROKER_INFRA_RESOURCE:
                    component = BrokerInfraSpec.fromJson((JsonObject) entry.getValue());
                    break;
                case InfraSpecComponent.ROUTER_INFRA_RESOURCE:
                    component = RouterInfraSpec.fromJson((JsonObject) entry.getValue());
                    break;
            }
            if(component!=null) {
                components.add(component);
            }
        });

        return new InfraConfigDefinition(
                metadataDef.getString("name"),
                infraDefinition.getString("kind").equals("StandardInfraConfig") ? AddressSpaceType.STANDARD : AddressSpaceType.BROKERED,
                components, spec.getString("version"));
    }
}
