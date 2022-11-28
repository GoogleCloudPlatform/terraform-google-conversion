test:
	GO111MODULE=on go test -short ./...

test-integration:
	go version
	terraform --version
	go test ./e2etest

.PHONY: test
