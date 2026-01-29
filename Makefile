SRC_DIR := .
BUILD_DIR := ./bin

GOLANGCI_LINT_VERSION := v2.8.0

all: check lint test build

build: clean
	go build -o $(BUILD_DIR)/api $(SRC_DIR)/cmd/api/*.go
	go build -o $(BUILD_DIR)/audit $(SRC_DIR)/cmd/audit/*.go

check:
	go fmt ./...
	go fix ./...
	go vet ./...

lint:
	docker run --rm -v $(shell pwd):/app:cached \
		-v $(shell go env GOCACHE):/cache/go \
		-v $(shell go env GOPATH)/pkg:/go/pkg \
		-e GOCACHE=/cache/go \
		-e GOLANGCI_LINT_CACHE=/cache/go \
		-w /app golangci/golangci-lint:${GOLANGCI_LINT_VERSION} \
		golangci-lint run --config .golangci.yml -v

test:
	go test $(SRC_DIR)/...

clean:
	rm -rf $(BUILD_DIR)

docker-build: docker-build-api docker-build-audit

docker-build-api:
	docker build -f Dockerfile-api -t api:latest .

docker-build-audit:
	docker build -f Dockerfile-audit -t audit:latest .
