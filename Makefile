test:
	GO111MODULE=on go test -short ./...

test-integration:
	go version
	terraform --version
	TEST_PROJECT=${TEST_PROJECT} go test ./e2etest

.PHONY: test
