# MAKEFILE
#
# @author      Nicola Asuni <info@tecnick.com>
# @link        https://github.com/nexmoinc/numkey
# ------------------------------------------------------------------------------

# Use bash as shell (Note: Ubuntu now uses dash which doesn't support PIPESTATUS).
SHELL=/bin/bash

# Project name
PROJECT=numkey

# Project version
VERSION=$(shell cat ../VERSION)

# Project release number (packaging build number)
RELEASE=$(shell cat ../RELEASE)

# Current directory
CURRENTDIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

# --- MAKE TARGETS ---

# Display general help about this command
.PHONY: help
help:
	@echo ""
	@echo "${PROJECT} Python Makefile."
	@echo "The following commands are available:"
	@echo ""
	@echo "    make build   : Build and test the library"
	@echo "    make format  : Format the source code"
	@echo "    make clean   : Remove any build artifact"
	@echo ""

all: clean build


# Set the version from VERSION file
.PHONY: version
version:
	sed -i "s/version=.*,$$/version=\"$(VERSION).$(RELEASE)\",/" setup.py

# Build and test the package
.PHONY: build
build: clean version
	rm -rf venv
	virtualenv -p /usr/bin/python3 venv
	source venv/bin/activate && python --version && pip install --pre black && pip install -e .[test] && python setup.py sdist bdist_wheel install test

# Generate source code documentation
.PHONY: doc
doc:
	pydoc -p 1234 $(PROJECT)

# Format the source code
.PHONY: format
format:
	astyle --style=allman --recursive --suffix=none 'numkey/*.h'
	astyle --style=allman --recursive --suffix=none 'numkey/*.c'
	black setup.py

# Remove any build artifact
.PHONY: clean
clean:
	rm -rf venv target c Dockerfile htmlcov build dist .pytest_cache .cache .benchmarks ./test/*.so ./test/__pycache__ ./numkey/__pycache__ ./numkey.egg-info
	find . -type f -name '*.pyc' -exec rm -f {} \;
