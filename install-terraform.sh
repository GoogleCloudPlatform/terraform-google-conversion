#!/bin/bash
set -e

TERRAFORM_BINARY=/terraform/$TERRAFORM_VERSION
if ! test -f "$TERRAFORM_BINARY"; then
    echo "Download terraform binary"
    wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip
    mkdir /terraform
    mv terraform /terraform/${TERRAFORM_VERSION}
fi
echo setting terraform to version ${TERRAFORM_VERSION}
set x-
mv /terraform/${TERRAFORM_VERSION} /bin/terraform
set x+
terraform version
