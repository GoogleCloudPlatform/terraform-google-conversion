# Contributing

If you want to contribute to Terraform Google Conversion, check out the
[contribution guidelines](../../CONTRIBUTING.md)

## Table of Contents

- [Contributing](#contributing)
  - [Table of Contents](#table-of-contents)
  - [Testing](#testing)
    - [Unit Test](#unit-test)
    - [Integration Test](#integration-test)
    - [Add a new resource](#add-a-new-resource)

## Testing

The read test(one of the unit tests) and integration test installs the provider
binary and creates a terraform dev override file tf-dev-override.tfrc.
[Dev override](https://googlecloudplatform.github.io/magic-modules/develop/run-tests/#optional-test-manually)
is to make sure the test runs against the provider version specified in the go
module. You can place your own config to point to a different go binary
destination within tf-dev-override.tfrc. If the file already exists, it will not
be overwritten.

It is better to run both the unit test and integration tests inside a Docker
container to match the CI/CD pipeline.

**Note:** Integration tests require a test project.

```bash
# Integration tests (interacts with real APIs)

# obtain credentials
gcloud auth application-default login
export TEST_PROJECT=my-project-id
export TEST_CREDENTIALS=~/.config/gcloud/application_default_credentials.json

# Spin up a docker container to run tests.
make run-docker
```

### Unit Test

The read test(tfplan2cai/test/read_test.go) runs terraform binary to generate
`.tfplan.json` files. It requires gcloud credential during that process.

```bash
# Inside the docker container,
# run this to setup dev override and run unit tests.
make test
```

### Integration Test

```bash
# Inside the docker container,
# run this to setup dev override and run integration tests.
make test-integration
```

### Add a new resource

See [Add a new resource](./add_new_resource.md)
