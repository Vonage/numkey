# MAKEFILE
#
# @author      Nicola Asuni <nicola.asuni@vonage.com>
# @link        https://github.com/Vonage/numkey
# ------------------------------------------------------------------------------

SHELL=/bin/bash
.SHELLFLAGS=-o pipefail -c

# Display general help about this command
.PHONY: help
help:
	@echo ""
	@echo "NumKey Java Makefile."
	@echo "The following commands are available:"
	@echo ""
	@echo "    make build   : Build the library"
	@echo "    make clean   : Remove any build artifact"
	@echo "    make test    : Run the unit tests against source code"
	@echo ""

all: clean build test

# Build the library
.PHONY: build
build:
	./gradlew jar

# Remove any build artifact
.PHONY: clean
clean:
	./gradlew clean

# Run the unit tests
.PHONY: test
test:
	./gradlew test
