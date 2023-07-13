SCOPE_SLACK := scope-slack
CFLAGS := -O2 -g -Wall -Werror $(CFLAGS)

GO ?= $(shell which go 2>&1)
ifeq (,$(GO))
$(error "error: \`go\` not in PATH; install or set GO to it's path")
endif

# Define a variable to store the list of Go files
GO_FILES := $(shell find . -name "*.go" ! -name "*bpfel*.go" -type f)

all: build
build: scope-slack

scope-slack:
	$(GO) build -tags "netgo" -ldflags="-extldflags=-static" -o bin/scope-slack ./cmd/scope-slack
	chmod +x bin/${SCOPE_SLACK}

fmt:
	@for file in $(GO_FILES); do \
		$(GO) fmt $$file; \
	done

vet:
	@for file in $(GO_FILES); do \
		$(GO) vet $$file; \
	done


help:
	@echo "Available targets:"
	@echo "  all             - Default target, builds the scope-slack binary"
	@echo "  scope-slack     - Builds the scope-slack binary"
	@echo "  fmt             - Formats Go source files"
	@echo "  vet             - Runs Go vet on source files"


.PHONY: all build fmt help scope-slack vet
