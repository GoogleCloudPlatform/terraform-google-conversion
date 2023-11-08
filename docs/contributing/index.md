# Contributing

If you want to contribute to Terraform Google Conversion, check out the [contribution guidelines](../../CONTRIBUTING.md)

## Table of Contents

- [Contributing](#contributing)
  - [Table of Contents](#table-of-contents)
  - [Testing](#testing)
    - [Unit Test](#unit-test)
    - [Integration Test](#integration-test)
    - [Add a new resource](#add-a-new-resource)


## Testing

### Unit Test
```bash
# Unit tests
make test
```

### Integration Test

**Note:** Integration tests require a test project.

It is better to run the integration tests inside a Docker container to match the CI/CD pipeline.

The integration test installs the provider binary and creates a terraform dev override file tf-dev-override.tfrc. [Dev override](https://googlecloudplatform.github.io/magic-modules/develop/run-tests/#optional-test-manually) is to make sure the test runs against the provider version specified in the go module. You can place your own config within tf-dev-override.tfrc. If the file already exists, it will not be overwritten.

```bash
# Integration tests (interacts with real APIs)

# obtain credentials
gcloud auth application-default login
export TEST_PROJECT=my-project-id
export TEST_CREDENTIALS=~/.config/gcloud/application_default_credentials.json

# Spin up a docker container to run tests.
make run-docker

# Inside the docker container,
# run this to setup dev override and run integration tests.
make test-integration
```


### Add a new resource

See [Add a new resource](./add_new_resource.md)
