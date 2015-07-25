GO_FILES = $(shell find . -type f -name '*.go')
PACKAGES = $(shell ls src/)
GOPATH = $(shell pwd):$(shell pwd)/vendor
VENDORS_PATH = $(shell pwd)/vendor
all: setup build

build: $(GO_FILES)
	@GOPATH=$(GOPATH) go build -o bin/zookup

run: build
	./bin/zookup -talk

clean:
	rm -rf pkg/* bin/*

package: setup build
	@cp -r bin package

setup:
	@brew install zookeeper
	@brew install bzr
	@GOPATH=$(VENDORS_PATH) CGO_CFLAGS='-I/usr/local/include/zookeeper' CGO_LDFLAGS='/usr/local/lib/libzookeeper_mt.a' go get launchpad.net/gozk
