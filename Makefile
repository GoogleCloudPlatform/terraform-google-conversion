build_dir=bin
TF_CONFIG_FILE=tf-dev-override.tfrc

build:
	GO111MODULE=on go build -o ./${build_dir}/tfplan2cai ./cmd/tfplan2cai
	cd v6;\
		GO111MODULE=on go build -o ../${build_dir}/tgc .

test:
	go version
	terraform --version
	./config-tf-dev-override.sh
	TF_CLI_CONFIG_FILE="$${PWD}/${TF_CONFIG_FILE}" GO111MODULE=on go test -short ./...
	cd v6;\
		TF_CLI_CONFIG_FILE="$${PWD}/${TF_CONFIG_FILE}" GO111MODULE=on go test -short ./...

test-integration:
	go version
	terraform --version
	./config-tf-dev-override.sh
	TF_CLI_CONFIG_FILE="$${PWD}/${TF_CONFIG_FILE}" go test -run=CLI ./...

test-go-licenses:
	cd .. && go version && go install github.com/google/go-licenses@latest
	$$(go env GOPATH)/bin/go-licenses check ./... --ignore github.com/dnaeon/go-vcr

run-docker:
	docker run -it \
	-v `pwd`:/terraform-google-conversion \
	-v ${GOOGLE_APPLICATION_CREDENTIALS}:/terraform-google-conversion/credentials.json \
	-w /terraform-google-conversion \
	--entrypoint=/bin/bash \
	--env TEST_PROJECT=${PROJECT_ID} \
	--env GOOGLE_APPLICATION_CREDENTIALS=/terraform-google-conversion/credentials.json \
	gcr.io/graphite-docker-images/go-plus;

release:
	./release.sh ${VERSION}

.PHONY: build test test-integration test-go-licenses run-docker release
