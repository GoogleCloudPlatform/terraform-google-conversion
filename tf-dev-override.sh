#!/bin/bash
set -ex

TF_CONFIG_FILE="tf-dev-override.tfrc"

# required for go install terraform-provider-google-beta
go mod download github.com/hashicorp/terraform-plugin-mux

go install github.com/hashicorp/terraform-provider-google-beta

# create terraform configuration file
if ! [ -f $TF_CONFIG_FILE ];then
  cat <<EOF > $TF_CONFIG_FILE
  provider_installation {
    # Developer overrides will stop Terraform from downloading the listed
    # providers their origin provider registries.
    dev_overrides {
        "hashicorp/google-beta" = "$GOPATH/bin"
    }
    # For all other providers, install them directly from their origin provider
    # registries as normal. If you omit this, Terraform will _only_ use
    # the dev_overrides block, and so no other providers will be available.
    # Without this, show "Failed to query available provider packages"
    # at terraform init
    direct{}
  }
EOF
fi
