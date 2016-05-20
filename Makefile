GO_FILES = $(shell find . -type f -name '*.go')
PACKAGES = $(shell ls src/)
GOPATH = $(shell pwd):$(shell pwd)/vendor
VENDORS_PATH = $(shell pwd)/vendor
all: setup build

build: $(GO_FILES)
	@GOPATH=$(GOPATH) go build -o bin/finance-ui

clean:
	rm -rf pkg/* bin/*

setup:
	@GOPATH=$(VENDORS_PATH) go get github.com/gorilla/mux
	@GOPATH=$(VENDORS_PATH) go get gopkg.in/throttled/throttled.v2/store/memstore
	@GOPATH=$(VENDORS_PATH) go get gopkg.in/throttled/throttled.v2
