#!/bin/bash

# This script is for deploying EnMasse into OpenShift. The target of
# installation can be an existing OpenShift deployment or an all-in-one
# container can be started.
#
# In either case, access to the `oc` command is required.
#
# example usage:
#
#    $ enmasse-deploy.sh -c 10.0.1.100 -o enmasse.10.0.1.100.xip.io
#
# this will deploy EnMasse into the OpenShift cluster running at 10.0.1.100
# and set the EnMasse webui route url to enmasse.10.0.1.100.xip.io.
# further it will use the user `developer` and project `myproject`, asking
# for a login when appropriate.
# for further parameters please see the help text.

if which oc &> /dev/null
then :
else
    echo "Cannot find oc command, please check path to ensure it is installed"
    exit 1
fi

ENMASSE_TEMPLATE_MASTER_URL=https://raw.githubusercontent.com/EnMasseProject/openshift-configuration/master/generated
TEMPLATE_NAME=enmasse
TEMPLATE_PARAMS=""

DEFAULT_OPENSHIFT_USER=developer
DEFAULT_OPENSHIFT_PROJECT=myproject

while getopts c:dk:o:p:s:t:u:h opt; do
    case $opt in
        c)
            OS_CLUSTER=$OPTARG
            ;;
        d)
            OS_ALLINONE=true
            ;;
        k)
            SERVER_KEY=$OPTARG
            ;;
        o)
            TEMPLATE_PARAMS="MESSAGING_HOSTNAME=$OPTARG $TEMPLATE_PARAMS"
            ;;
        p)
            PROJECT=$OPTARG
            ;;
        s)
            SERVER_CERT=$OPTARG
            ;;
        t)
            ALT_TEMPLATE=$OPTARG
            ;;
        u)
            OS_USER=$OPTARG
            USER_REQUESTED=true
            ;;
        h)
            echo "usage: enmasse-deploy.sh [options]"
            echo
            echo "deploy the EnMasse suite into a running OpenShift cluster"
            echo
            echo "optional arguments:"
            echo "  -h             show this help message"
            echo "  -c CLUSTER     OpenShift cluster url to login against (default: https://localhost:8443)"
            echo "  -d             create an all-in-one docker OpenShift on localhost"
            echo "  -k KEY         Server key file (default: none)"
            echo "  -o HOSTNAME    Custom hostname for messaging endpoint (default: use autogenerated from template)"
            echo "  -p PROJECT     OpenShift project name to install EnMasse into (default: $DEFAULT_OPENSHIFT_PROJECT)"
            echo "  -s CERT        Server certificate file (default: none)"
            echo "  -t TEMPLATE    An alternative opan OpenShift template file to deploy EnMasse (default: curl'd from upstream)"
            echo "  -u USER        OpenShift user to run commands as (default: $DEFAULT_OPENSHIFT_USER)"
            echo
            exit
            ;;
        \?)
            echo "Invalid option: -$OPTARG" >&2
            exit
            ;;
    esac
done

if [ -z "$OS_USER" ]
then
    echo "user not set, using default value"
    OS_USER=$DEFAULT_OPENSHIFT_USER
fi

if [ -z "$PROJECT" ]
then
    echo "project name not set, using default value"
    PROJECT=$DEFAULT_OPENSHIFT_PROJECT
fi

if [ -n "$OS_ALLINONE" ]
then
    if [ -n "$OS_CLUSTER" ]
    then
        echo "Error: You have requested an all-in-one deployment AND specified a cluster address."
        echo "Please choose one of these options and restart."
        exit 1
    fi
    if [ -n "$USER_REQUESTED" ]
    then
        echo "Error: You have requested an all-in-one deployment AND specified an OpenShift user."
        echo "Please choose either all-in-one or a cluster deployment if you need to use a specific user."
        exit 1
    fi
    sudo oc cluster up
fi


oc login $OS_CLUSTER -u $OS_USER

AVAILABLE_PROJECTS=`oc projects -q`

for proj in $AVAILABLE_PROJECTS
do
    if [ "$proj" == "$PROJECT" ]; then
        oc project $proj
        break
    fi
done

CURRENT_PROJECT=`oc project -q`
if [ "$CURRENT_PROJECT" != "$PROJECT" ]; then
    oc new-project $PROJECT
fi

oc create sa enmasse-service-account -n $PROJECT
oc policy add-role-to-user view system:serviceaccount:${PROJECT}:default
oc policy add-role-to-user edit system:serviceaccount:${PROJECT}:enmasse-service-account


if [ -n "$SERVER_KEY" ] && [ -n "$SERVER_CERT" ]
then
    TEMPLATE_NAME=tls-enmasse
    oc secret new qdrouterd-certs ${SERVER_CERT} ${SERVER_KEY}
    oc secret add serviceaccount/default secrets/qdrouterd-certs --for=mount
    # secret for MQTT certificates
    oc secret new mqtt-certs ${SERVER_CERT} ${SERVER_KEY}
    oc secret add serviceaccount/default secrets/mqtt-certs --for=mount
fi

if [ -n "$ALT_TEMPLATE" ]
then
    ENMASSE_TEMPLATE=$ALT_TEMPLATE
    oc process -f $ENMASSE_TEMPLATE $TEMPLATE_PARAMS | oc create -n $PROJECT -f -
else
    ENMASSE_TEMPLATE=${ENMASSE_TEMPLATE_MASTER_URL}/${TEMPLATE_NAME}-template.yaml
    oc create -f $ENMASSE_TEMPLATE
    if [ -n "$TEMPLATE_PARAMS" ]
    then
        oc new-app -n $PROJECT --template=$TEMPLATE_NAME -p $TEMPLATE_PARAMS
    else
        oc new-app -n $PROJECT --template=$TEMPLATE_NAME
    fi
fi
