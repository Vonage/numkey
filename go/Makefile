# MAKEFILE
#
# @author      Nicola Asuni <info@tecnick.com>
# @link        https://github.com/nexmoinc/numkey
# ------------------------------------------------------------------------------

# Use bash as shell (Note: Ubuntu now uses dash which doesn't support PIPESTATUS).
SHELL=/bin/bash

# CVS path (path to the parent dir containing the project)
CVSPATH=github.com/nexmoinc

# Project vendor
VENDOR=nexmoinc

# Project name
PROJECT=numkey

# Project version
VERSION=$(shell cat ../VERSION)

# Project release number (packaging build number)
RELEASE=$(shell cat ../RELEASE)

# Name of RPM or DEB package
PKGNAME=${VENDOR}-${PROJECT}

# Current directory
CURRENTDIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

# GO lang path
ifeq ($(GOPATH),)
	# extract the GOPATH
	GOPATH=$(firstword $(subst /src/, ,$(CURRENTDIR)))
endif

# Add the GO binary dir in the PATH
export PATH := $(GOPATH)/bin:$(PATH)

# Common commands
GO := GOPATH=$(GOPATH) GOPRIVATE=$(CVSPATH) go
GOFMT := gofmt
GOTEST := GOPATH=$(GOPATH) gotest

# Directory containing the source code
SRCDIR=./src

# List of packages
GOPKGS := $(shell $(GO) list $(SRCDIR)/...)

# Enable junit report when not in LOCAL mode
ifeq ($(strip $(DEVMODE)),LOCAL)
	TESTEXTRACMD=&& $(GO) tool cover -func=target/report/coverage.out
else
	TESTEXTRACMD=2>&1 | tee >(PATH=$(GOPATH)/bin:$(PATH) go-junit-report > target/test/report.xml); test $${PIPESTATUS[0]} -eq 0
endif

# --- MAKE TARGETS ---

# Display general help about this command
.PHONY: help
help:
	@echo ""
	@echo "$(PROJECT) Makefile."
	@echo "GOPATH=$(GOPATH)"
	@echo "The following commands are available:"
	@echo ""
	@echo "    make qa       : Run all tests and static analysis tools"
	@echo "    make test     : Run unit tests"
	@echo "    make coverage : Generate the coverage report"
	@echo ""
	@echo "    make format   : Format the source code"
	@echo "    make generate : Generate go code automatically"
	@echo "    make linter   : Check code against multiple linters"
	@echo "    make docs     : Generate static documentation"
	@echo ""
	@echo "    make deps     : Get dependencies"
	@echo "    make clean    : Remove any build artifact"
	@echo ""

# Alias for help target
all: clean deps qa

# Create the trget directories if missing
.PHONY: ensuretarget
ensuretarget:
	@mkdir -p target/test
	@mkdir -p target/report
	@mkdir -p target/docs

# Generate test mocks
.PHONY: generate
generate:
	$(GO) generate $(GOPKGS)

# Run the unit tests
.PHONY: test
test: ensuretarget
	@echo -e "\n\n>>> START: Unit Tests <<<\n\n"
	$(GOTEST) \
	-tags=unit \
	-covermode=atomic \
	-bench=. \
	-race \
	-failfast \
	-coverprofile=target/report/coverage.out \
	-v $(GOPKGS) $(TESTEXTRACMD)
	@echo -e "\n\n>>> END: Unit Tests <<<\n\n"

# Format the source code
.PHONY: format
format:
	@find ${SRCDIR} -type f -name "*.go" -exec $(GOFMT) -s -w {} \;

# Generate the coverage report
.PHONY: coverage
coverage:
	@mkdir -p target/report
	$(GO) tool cover -html=target/report/coverage.out -o target/report/coverage.html

.PHONY: linter
linter:
	./target/tools/golangci-lint run $(SRCDIR)/...

# Generate source docs
.PHONY: docs
docs: ensuretarget
	GOPATH=$(GOPATH) godoc -url "http://localhost:6060/pkg/${CVSPATH}/${PROJECT}/go/src/" > target/docs/index.html
	sed -i 's|="/lib/godoc/|="|g' target/docs/index.html
	cp $(GOPATH)/src/golang.org/x/tools/godoc/static/*.js target/docs/
	cp $(GOPATH)/src/golang.org/x/tools/godoc/static/*.css target/docs/
	sed -i "s|\"/src/${CVSPATH}/${PROJECT}/|https://${CVSPATH}/${PROJECT}/blob/master/|g" target/docs/index.html

# Run all tests and static analysis tools
.PHONY: qa
qa: test coverage linter

# Get the test dependencies
.PHONY: deps
deps: ensuretarget
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./target/tools v1.27.0
	(GO111MODULE=off $(GO) get github.com/jstemmer/go-junit-report)
	(GO111MODULE=off $(GO) get github.com/rakyll/gotest)
	(GO111MODULE=off $(GO) get github.com/golang/mock/mockgen)

# Remove any build artifact
.PHONY: clean
clean:
	rm -rf ./target
	$(GO) clean -i ./...

# build example
.PHONY: example
example:
	$(GO) build -o ./target/example ./example \
	&& ./target/example
