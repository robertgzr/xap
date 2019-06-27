GO ?= go
PROJECT := github.com/robertgzr/xap
BUILDTAGS ?= static_build netgo osusergo

VERSION ?= $(shell git describe --tag --always)
DATE ?= $(shell date +%FT%T)
COMMIT ?= $(shell git rev-parse HEAD)

LDFLAGS ?= -s -X main.version=$(VERSION) -X main.date=$(DATE) -X main.commit=$(COMMIT)
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
	$(RM) -r bin/ dist/

.PHONY: \
    binaries \
    clean \
    install
