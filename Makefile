# This version-strategy uses git tags to set the version string
VERSION ?= $(shell git describe --tags --always --dirty)
#
# This version-strategy uses a manual value to set the version string
#VERSION ?= 1.2.3

version: # @HELP outputs the version string
version:
	@echo $(VERSION)

fmt: # @HELP format code
	find . -type f | grep .go\$$ | grep -v vendor | xargs -I {} gofmt -w {}

test: # @HELP run all tests
test: fmt
	go test ./...
	
help: # @HELP show commands
help:
	@echo
	@echo "TARGETS:"
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)    \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '