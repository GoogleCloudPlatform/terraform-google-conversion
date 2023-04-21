TERRAFORM_VERSION=0.12.31

build:
	go build ./...

test:
	GO111MODULE=on go test -short ./...

test-integration:
	go version
	terraform --version
	go test -run=CLI ./...

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
