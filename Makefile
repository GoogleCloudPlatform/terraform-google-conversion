test:
	GO111MODULE=on go test -short ./...

test-integration:
	go version
	terraform --version
	TEST_CREDENTIALS=${GOOGLE_APPLICATION_CREDENTIALS} TEST_PROJECT=${PROJECT_ID} go test ./e2etest

.PHONY: test
