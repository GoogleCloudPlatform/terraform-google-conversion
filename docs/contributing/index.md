# Contributing

If you want to contribute to Terraform Google Conversion, check out the [contribution guidelines](../../CONTRIBUTING.md)

## Table of Contents

- [Contributing](#contributing)
  - [Table of Contents](#table-of-contents)
  - [Testing](#testing)
    - [Docker](#docker)
    - [Add a new resource](#add-a-new-resource)


## Testing

**Note:** Integration tests require a test project.

```
# Unit tests
make test

# Integration tests (interacts with real APIs)
gcloud auth application-default login
export TEST_PROJECT=my-project-id
export TEST_CREDENTIALS=~/.config/gcloud/application_default_credentials.json
make test-integration

# Specific integration test
make build
go test -v -run=<test name or prefix> ./tfplan2cai/test
```

### Docker
It is better to run the integration tests inside a Docker container to match the CI/CD pipeline.
First, build the Docker container:

```
make build-docker
```

For obtaining a credentials file, run
```
gcloud auth application-default login
```

Start the Docker container:

```
export TEST_PROJECT=my-project-id
export GOOGLE_APPLICATION_CREDENTIALS=$(pwd)/credentials.json
make run-docker
# install terraform, check Makefile for the default version of terraform.
/install-terraform.sh
```

Finally, run the integration tests inside the container:
```
make test-integration
```

### Add a new resource

See [Add a new resource](./add_new_resource.md)
