# MAKEFILE
#
# @author      Nicola Asuni
# @link        https://github.com/Vonage/numkey
# ------------------------------------------------------------------------------

SHELL=/bin/bash
.SHELLFLAGS=-o pipefail -c

# Project name
PROJECT=numkey

# Project version
VERSION=$(shell cat ../VERSION)

# Project release number (packaging build number)
RELEASE=$(shell cat ../RELEASE)

# Path for the selected Python version
PYTHON=$(shell which python3)

# --- MAKE TARGETS ---

# Display general help about this command
.PHONY: help
help:
	@echo ""
	@echo "${PROJECT} Python Makefile."
	@echo "The following commands are available:"
	@echo ""
	@echo "    make build      : Build the library"
	@echo "    make clean      : Remove any build artifact"
	@echo "    make doc        : Generate documentation"
	@echo "    make format     : Format the source code"
	@echo "    make test       : Test the library"
	@echo "    make venv       : Create a virtual environment"
	@echo "    make version    : Set the version from VERSION file"
	@echo ""

all: clean venv build test

# Build and test the package
.PHONY: build
build: version
	source venv/bin/activate \
	&& python -m build --sdist --wheel

# Remove any build artifact
.PHONY: clean
clean:
	rm -rf venv target c Dockerfile htmlcov build dist .pytest_cache .cache .benchmarks ./test/*.so ./test/__pycache__ ./numkey/__pycache__ ./numkey.egg-info
	find . -type f -name '*.pyc' -exec rm -f {} \;

# Generate source code documentation
.PHONY: doc
doc:
	pydoc3 -p 1234 $(PROJECT)

# Format the source code
.PHONY: format
format:
	astyle --style=allman --recursive --suffix=none 'numkey/*.h'
	astyle --style=allman --recursive --suffix=none 'numkey/*.c'
	black setup.py

# Test the package
.PHONY: test
test:
	source venv/bin/activate \
	&& pytest

# Create virtual environment
.PHONY: venv
venv: clean version
	rm -rf venv
	virtualenv -p $(PYTHON) venv
	source venv/bin/activate \
	&& python --version \
	&& pip install --upgrade pip \
	&& pip install --upgrade --pre wheel build black pytest setuptools \
	&& pip install --use-pep517 --no-build-isolation -e .[test]

# Set the version from VERSION file
.PHONY: version
version:
	sed -i "s/version=.*,$$/version=\"$(VERSION).$(RELEASE)\",/" setup.py
