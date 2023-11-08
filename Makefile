TERRAFORM_VERSION=0.12.31

build_dir=./bin

build:
	GO111MODULE=on go build -o ${build_dir}/tfplan2cai ./cmd/tfplan2cai

test:
	GO111MODULE=on go test -short ./...

test-integration:
	env
	go version
	terraform --version
	go test -v -run=CLI ./...

test-go-licenses:
	cd .. && go version && go install github.com/google/go-licenses@latest
	$$(go env GOPATH)/bin/go-licenses check ./... --ignore github.com/dnaeon/go-vcr

build-docker:
	docker build --build-arg TERRAFORM_VERSION=$(TERRAFORM_VERSION) -f ./Dockerfile -t terraform-google-conversion .

run-docker:
	docker run -it -v `pwd`:/terraform-google-conversion -v ${GOOGLE_APPLICATION_CREDENTIALS}:/terraform-google-conversion/credentials.json --entrypoint=/bin/bash --env TEST_PROJECT=${PROJECT_ID} --env GOOGLE_APPLICATION_CREDENTIALS=/terraform-google-conversion/credentials.json terraform-google-conversion;

release:
	./release.sh ${VERSION}

.PHONY: test test-integration test-go-licenses build-docker run-docker release
