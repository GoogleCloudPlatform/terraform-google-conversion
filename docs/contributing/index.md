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

The [read test](../../tfplan2cai/test/read_test.go) and integration test apply
[dev override](https://googlecloudplatform.github.io/magic-modules/develop/run-tests/#optional-test-manually)
to make sure the test runs against the provider version specified in the go
module.

The dev override configuration installs the provider binary into GOPATH/bin and
creates a terraform dev override file tf-dev-override.tfrc. You can create your
tf-dev-override.tfrc with your go binary destination. If the file already
exists, it will not be overwritten.

It is better to run both the unit test and integration tests inside the Docker
container to match the CI/CD pipeline, where the docker image has a
pre-installed terraform binary.

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

Unit tests include [read test](../../tfplan2cai/test/read_test.go), which runs
terraform binary to generate `.tfplan.json` files.

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
