GO ?= go
PROJECT := github.com/robertgzr/xap
BUILDTAGS ?= static_build netgo osusergo

VERSION ?= $(shell git describe --tag --always)
BUILD_INFO ?= $(shell date +%FT%T)

LDFLAGS ?= -s -X main.version=$(VERSION) -X main.buildInfo=$(BUILD_INFO)
GOFLAGS := -ldflags '$(LDFLAGS)' -tags "$(BUILDTAGS)"

binaries: xap xap-radio

xap:
	$(GO) build $(GOFLAGS) -o bin/$@ $(PROJECT)

xap-radio:
	$(GO) build $(GOFLAGS) -o bin/$@ $(PROJECT)/plugins/xap-radio

install:
	install bin/xap -t $(GOPATH)/bin
	install bin/xap-radio -t $(GOPATH)/bin

clean:
	$(RM) -r bin/

.PHONY: \
    binaries \
    clean \
    install
