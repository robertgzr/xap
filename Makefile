GO ?= go
PROJECT := github.com/robertgzr/xap
BUILDTAGS ?= static_build netgo osusergo

PREFIX ?= ${DESTDIR}/usr/local
BINDIR ?= ${PREFIX}/bin
MANDIR ?= ${PREFIX}/share/man

COMMIT_NO ?= $(shell git rev-parse HEAD 2> /dev/null || true)
GIT_COMMIT ?= $(if $(shell git status --porcelain --untracked-files=no),${COMMIT_NO}-dirty,${COMMIT_NO})
VERSION ?= $(GIT_COMMIT)
BUILD_INFO ?= $(shell date +%F_%T)

LDFLAGS ?= -s
XAP_LDFLAGS ?= $(LDFLAGS) -X $(PROJECT)/command.version=$(VERSION) -X $(PROJECT)/command.buildInfo=$(BUILD_INFO)
XAPRADIO_LDFLAGS ?= $(LDFLAGS) -X $(PROJECT)/plugins/xap-rad-io.version=$(VERSION) -X $(PROJECT)/plugins/xap-rad-io.buildInfo=$(BUILD_INFO)

binaries: xap xap-rad-io

xap:
	$(GO) build -ldflags '$(XAP_LDFLAGS)' -tags "$(BUILDTAGS)" -o bin/$@ $(PROJECT)

xap-rad-io:
	$(GO) build -ldflags '$(XAPRADIO_LDFLAGS)' -tags "$(BUILDTAGS)" -o bin/$@ $(PROJECT)/plugins/xap-rad-io


MANPAGES ?=
$(MANPAGES):
	@echo "not yet :("

docs: $(MANPAGES)

install: \
    install.bin \
    install.man

install.bin: install.xap install.xap-rad-io
install.xap: xap
	install ${SELINUXOPT} -d -m 755 $(BINDIR)
	install ${SELINUXOPT} -m 755 bin/xap $(BINDIR)/xap
	test -z "${SELINUXOPT}" || chcon --verbose --reference=$(BINDIR)/xap bin/xap
install.xap-rad-io: xap-rad-io
	install ${SELINUXOPT} -d -m 755 $(BINDIR)
	install ${SELINUXOPT} -m 755 bin/xap-rad-io $(BINDIR)/xap-rad-io
	test -z "${SELINUXOPT}" || chcon --verbose --reference=$(BINDIR)/xap-rad-io bin/xap-rad-io

install.man: docs
	@echo "onoo, not done yet :("

uninstall:
	rm -f $(BINDIR)/xap
	rm -f $(BINDIR)/xap-rad-io
	for i in $(filter %.1,$(MANPAGES)); do \
		rm -f $(MANDIR)/man1/$$(basename $${i}); \
	done; \
	for i in $(filter %.5,$(MANPAGES)); do \
		rm -f $(MANDIR)/man5/$$(basename $${i}); \
	done

clean:
	$(RM) -r bin/

.PHONY: \
    binaries \
    clean \
    docs \
    uninstall
