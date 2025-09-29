SHELL := /bin/bash
BUF_VERSION ?= v1.32.2
GOBIN ?= $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN := $(shell go env GOPATH)/bin
endif
BUF_BIN := $(GOBIN)/buf

.PHONY: install build cleanup format lint

install: $(BUF_BIN)

$(BUF_BIN):
	@echo "Installing buf $(BUF_VERSION) to $(GOBIN)"
	@GO111MODULE=on GOBIN=$(GOBIN) go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)

build: $(BUF_BIN)
	./protogen.sh

format: $(BUF_BIN)
	buf format -w proto

lint: $(BUF_BIN)
	buf lint proto

cleanup:
	rm -rf gen
	@echo "Removed generated artifacts."
