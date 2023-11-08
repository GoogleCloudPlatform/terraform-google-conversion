#!/bin/bash
set -ex

TF_CONFIG_FILE="tf-dev-override.tfrc"

go clean --modcache
go list -json  -m github.com/hashicorp/terraform-provider-google-beta
REPLACE_DIR=`go list -json  -m github.com/hashicorp/terraform-provider-google-beta | jq -r '.Dir // empty'`
VERSION=`go list -json  -m github.com/hashicorp/terraform-provider-google-beta | jq -r .Version`

if [ ! -z "$REPLACE_DIR" ]
then
  pushd $REPLACE_DIR
    go install
  popd
else
  go install github.com/hashicorp/terraform-provider-google-beta@$VERSION
fi

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
