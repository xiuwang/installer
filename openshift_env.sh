#!/bin/bash
export OPENSHIFT_INSTALL_BASE_DOMAIN="devcluster.openshift.com"
export OPENSHIFT_INSTALL_CLUSTER_NAME="xiuwang"
export OPENSHIFT_INSTALL_EMAIL_ADDRESS="xiuwang@redhat.com"
export OPENSHIFT_INSTALL_PASSWORD="redhat"
export OPENSHIFT_INSTALL_PLATFORM="aws"
export OPENSHIFT_INSTALL_SSH_PUB_KEY_PATH=~/.ssh/libra.pub
export OPENSHIFT_INSTALL_PULL_SECRET_PATH=~/.secrets/openshift_pull_secret.json
export AWS_PROFILE="openshift-dev"  #(whatever you set up in your ~/.aws/config file)
export OPENSHIFT_INSTALL_AWS_REGION=$(aws configure get region)
#export OPENSHIFT_INSTALL_RELEASE_IMAGE_OVERRIDE=quay.io/openshift-release-dev/ocp-release:4.0.0-5
