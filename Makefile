test:
	GO111MODULE=on go test -short ./...

test-integration:
	go version
	terraform --version
	go test -run=CLI ./...

test-go-licenses:
	cd .. && go version && go install github.com/google/go-licenses@latest
	$$(go env GOPATH)/bin/go-licenses check ./...

run-docker:
	docker run -it -v `pwd`:/terraform-validator -v ${GOOGLE_APPLICATION_CREDENTIALS}:/terraform-validator/credentials.json --entrypoint=/bin/bash --env TEST_PROJECT=${PROJECT_ID} --env TEST_CREDENTIALS=./credentials.json terraform-validator;

.PHONY: test test-integration  test-go-licenses
