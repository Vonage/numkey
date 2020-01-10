# MAKEFILE
#
# @author      Nicola Asuni <nicola.asuni@vonage.com>
# @link        https://github.com/nexmoinc/numkey
# ------------------------------------------------------------------------------

# --- MAKE TARGETS ---

# Display general help about this command
.PHONY: help
help:
	@echo ""
	@echo "NumKey Java Makefile."
	@echo "The following commands are available:"
	@echo ""
	@echo "    make build   : Build the library"
	@echo "    make test    : Run the unit tests against source code"
	@echo "    make clean   : Remove any build artifact"
	@echo ""

all: clean build test

# Build the library
.PHONY: build
build:
	./gradlew jar

# Run the unit tests
.PHONY: test
test:
	./gradlew test

# Remove any build artifact
.PHONY: clean
clean:
	./gradlew clean
