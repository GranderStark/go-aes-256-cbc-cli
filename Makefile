# Origin: https://github.com/innogames/graphite-ch-optimizer/blob/master/Makefile
# MIT License
NAME = aes-256-cbc-cli
VERSION = $(shell git describe --always --tags --abbrev=0 2>/dev/null | sed 's/^v//;s/\([^-]*-g\)/c\1/;s/-/./g')
GIT_COMMIT = $(shell git rev-parse HEAD)
DATE = $(shell date +%F)
VENDOR = "Lev Subbotin <subveles@gmail.com>"
URL = https://github.com/GranderStark/$(NAME)
define DESC =
'Simple aes encryption-decryption ommand line tool'
endef
GO_FILES = $(shell find ./ -name '*.go')
GO_BUILD = go build -ldflags "-X 'main.version=$(VERSION)' -X 'main.gitCommit=$(GIT_COMMIT)' -X 'main.buildDate=$(DATE)'"
PKG_FILES = build/$(NAME)_$(VERSION)_amd64.deb build/$(NAME)-$(VERSION)-1.x86_64.rpm
export CGO_ENABLED = 0
export GOOS = $(shell uname -s | awk '{print tolower($0)}')
export GOARCH = amd64

.PHONY: clean all version test

all: build

version:
	@echo $(VERSION)

clean:
	rm -rf build
	rm -rf $(NAME)

rebuild: clean all

test:
	go vet ./...
	go test -v ./...

build: $(NAME)/$(NAME)

$(NAME)/config.yml: $(NAME)/$(NAME)
	./$(NAME)/$(NAME) default-config > $@

$(NAME)/$(NAME): $(GO_FILES)
	$(GO_BUILD) -o $@ ./cmd/$(NAME)

build/$(NAME): $(GO_FILES)
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_BUILD) -o $@ ./cmd/$(NAME)

packages: $(PKG_FILES)

.ONESHELL:
build/pkg: build/$(NAME)
	cd build
	mkdir -p pkg/etc/$(NAME)
	mkdir -p pkg/usr/bin
	cp -l $(NAME) pkg/usr/bin/

deb: $(word 1, $(PKG_FILES))

rpm: $(word 2, $(PKG_FILES))

# Set TYPE to package suffix w/o dot
$(PKG_FILES): TYPE = $(subst .,,$(suffix $@))
$(PKG_FILES): build/pkg
	fpm --verbose \
		-s dir \
		-a x86_64 \
		-t $(TYPE) \
		--vendor $(VENDOR) \
		-m $(VENDOR) \
		--url $(URL) \
		--description $(DESC) \
		--license MIT \
		-n $(NAME) \
		-v $(VERSION) \
		-p build \
		build/pkg/=/
