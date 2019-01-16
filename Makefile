export GO111MODULE=on

ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BIN_DIR = $(ROOT_DIR)/bin
APP_DIR = /go/src/github.com/ulule/helloapp

branch = $(shell git rev-parse --abbrev-ref HEAD)
commit = $(shell git log --pretty=format:'%h' -n 1)
now = $(shell date "+%Y-%m-%d %T UTC%z")
compiler = $(shell go version)

build:
	@(echo "-> Compiling helloapp binary")
	@(mkdir -p $(BIN_DIR))
	@(go build -mod=vendor -ldflags "\
		-X 'github.com/ulule/helloapp.Branch=$(branch)' \
		-X 'github.com/ulule/helloapp.Revision=$(commit)' \
		-X 'github.com/ulule/helloapp.BuildTime=$(now)' \
		-X 'github.com/ulule/helloapp.Compiler=$(compiler)'" -o $(BIN_DIR)/helloapp ./cmd/main.go)
	@(echo "-> helloapp binary created")

run:
	bin/helloapp

live:
	@modd

build-static:
	@(echo "-> Creating statically linked binary...")
	@(mkdir -p $(BIN_DIR))
	@(CGO_ENABLED=0 go build -mod=vendor -ldflags "\
		-X 'github.com/ulule/helloapp.Branch=$(branch)' \
		-X 'github.com/ulule/helloapp.Revision=$(commit)' \
		-X 'github.com/ulule/helloapp.BuildTime=$(now)' \
		-X 'github.com/ulule/helloapp.Compiler=$(compiler)'" -a -installsuffix cgo -o $(BIN_DIR)/helloapp ./cmd/main.go)

dump-vendor:
	@(go mod vendor)

docker-build:
	@(echo "-> Preparing builder...")
	@(docker build -t helloapp-builder -f Dockerfile.build .)
	@(mkdir -p $(BIN_DIR))
	@(echo "-> Running helloapp builder...")
	@(docker run --rm -v $(BIN_DIR):$(APP_DIR)/bin helloapp-builder)
