# These shell flags are REQUIRED for an early exit in case any program called by make errors!
.SHELLFLAGS=-euo pipefail -c
SHELL := /bin/bash

.PHONY: go-check go-test

CACHE := $(PWD)/.cache

##@ Global

clean: ## Clean directory.
	@rm -fr $(CACHE)

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Golang

export CGO_ENABLED := 0
export GOBIN := $(CACHE)/bin
export PATH := $(GOBIN):$(PATH)

GO_PKGS := $(shell go list go/...)

go-init: ## Runs go mod tidy && go mod install.
	@cd go/ && go mod tidy
	@cat go/tools.go | grep _ | cut -d " " -f 2 | tr -d \" | xargs -I % go install %@latest

go-check: go-init ## Run goimports and golangci-lint.
	@$(GOBIN)/goimports -w -l $(shell find . -type f -name "*.go")
	@$(GOBIN)/golangci-lint run --timeout=10m --verbose

go-test: go-init ## Run go unit tests.
	@go test -count=1 -race $(PKGS)
