# MAKEFILE
#
# @author      Nicola Asuni
# @link        https://github.com/Vonage/numkey
# ------------------------------------------------------------------------------

SHELL=/bin/bash
.SHELLFLAGS=-o pipefail -c

# Project owner
OWNER=Vonage

# Project vendor
VENDOR=${OWNER}

# Lowercase VENDOR name for Docker
LCVENDOR=$(shell echo "${VENDOR}" | tr '[:upper:]' '[:lower:]')

# CVS path (path to the parent dir containing the project)
CVSPATH=github.com/${VENDOR}

# Project name
PROJECT=numkey

# Project version
VERSION=$(shell cat ../VERSION)

# Project release number (packaging build number)
RELEASE=$(shell cat ../RELEASE)

# Current directory
CURRENTDIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

# Name of RPM or DEB package
PKGNAME=${LCVENDOR}-${PROJECT}-headers

# Path for binary files (where the executable files will be installed)
LIBPATH=usr/include/$(VENDOR)/$(PROJECT)/

# Path path for documentation
DOCPATH=usr/share/doc/$(PKGNAME)/

# Installation path for the binary files
PATHINSTLIB=$(DESTDIR)/$(LIBPATH)

# Installation path for documentation
PATHINSTDOC=$(DESTDIR)/$(DOCPATH)

# RPM Packaging path (where RPMs will be stored)
PATHRPMPKG=$(CURRENTDIR)/target/RPM

# DEB Packaging path (where DEBs will be stored)
PATHDEBPKG=$(CURRENTDIR)/target/DEB

# --- MAKE TARGETS ---

# Display general help about this command
.PHONY: 
help:
	@echo ""
	@echo "$(PROJECT) C Makefile."
	@echo "The following commands are available:"
	@echo ""
	@echo "  make build     : Build the library"
	@echo "  make clean     : Remove any build artifact"
	@echo "  make deb       : Build a DEB package"
	@echo "  make doc       : Generate source code documentation"
	@echo "  make format    : Format the source code"
	@echo "  make install   : Install the library"
	@echo "  make rpm       : Build an RPM package"
	@echo "  make test      : Run the unit tests"
	@echo "  make tidy      : Check the code using clang-tidy"
	@echo "  make uninstall : Uninstall the library"
	@echo "  make version   : Set version from VERSION file"
	@echo ""

.PHONY: all
all: clean version format test tidy build rpm deb doc

# Build the library
.PHONY: build
build:
	@mkdir -p target/build
	@echo -e "\n\n*** BUILD RELEASE ***\n"
	rm -rf target/build/*
	cd target/build && \
	cmake -DCMAKE_C_FLAGS=$(CMAKE_C_FLAGS) \
	-DCMAKE_TOOLCHAIN_FILE=$(CMAKE_TOOLCHAIN_FILE) \
	-DCMAKE_BUILD_TYPE=Release \
	-DCMAKE_INSTALL_PREFIX=$(CMAKE_INSTALL_PATH) \
	-DBUILD_SHARED_LIB=$(VH_BUILD_SHARED_LIB) \
	-DBUILD_DOXYGEN=$(VH_BUILD_DOXYGEN) \
	../.. | tee cmake.log ; test $${PIPESTATUS[0]} -eq 0 && \
	export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./ && \
	make | tee make.log ; test $${PIPESTATUS[0]} -eq 0
	cd target/build && \
	export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./ && \
	env CTEST_OUTPUT_ON_FAILURE=1 make test | tee build.log ; test $${PIPESTATUS[0]} -eq 0

# Remove any build artifact
.PHONY: clean
clean:
	rm -rf target

# Build the DEB package for Debian-like Linux distributions
.PHONY: deb
deb:
	rm -rf $(PATHDEBPKG)
	make install DESTDIR=$(PATHDEBPKG)/$(PKGNAME)-$(VERSION)
	rm -f $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/$(DOCPATH)LICENSE
	tar -zcvf $(PATHDEBPKG)/$(PKGNAME)_$(VERSION).orig.tar.gz -C $(PATHDEBPKG)/ $(PKGNAME)-$(VERSION)
	cp -rf ./resources/debian $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian
	find $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/ -type f -exec sed -i "s/~#DATE#~/`date -R`/" {} \;
	find $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/ -type f -exec sed -i "s/~#PKGNAME#~/$(PKGNAME)/" {} \;
	find $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/ -type f -exec sed -i "s/~#VERSION#~/$(VERSION)/" {} \;
	find $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/ -type f -exec sed -i "s/~#RELEASE#~/$(RELEASE)/" {} \;
	echo $(LIBPATH) > $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).dirs
	echo "$(LIBPATH)* $(LIBPATH)" > $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/install
	echo $(DOCPATH) >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).dirs
	echo "$(DOCPATH)* $(DOCPATH)" >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/install
	echo "new-package-should-close-itp-bug" >> $(PATHDEBPKG)/$(PKGNAME)-$(VERSION)/debian/$(PKGNAME).lintian-overrides
	cd $(PATHDEBPKG)/$(PKGNAME)-$(VERSION) && debuild -us -uc

# Generate source code documentation
.PHONY: doc
doc:
	cd target/build && \
	make doc | tee doc.log ; test $${PIPESTATUS[0]} -eq 0

# Format the source code
.PHONY: format
format:
	astyle --style=allman --recursive --suffix=none 'src/numkey/*.h'
	astyle --style=allman --recursive --suffix=none 'test/*.c'
	astyle --style=allman --recursive --suffix=none 'nk/*.c'

# Install this application
.PHONY: install
install: uninstall
	mkdir -p $(PATHINSTLIB)
	cp src/numkey/*.h $(PATHINSTLIB)
	find $(PATHINSTLIB) -type f -exec chmod 644 {} \;
	mkdir -p $(PATHINSTDOC)
	cp -f ../LICENSE $(PATHINSTDOC)
	cp -f ../README.md $(PATHINSTDOC)
	cp -f ../VERSION $(PATHINSTDOC)
	cp -f ../RELEASE $(PATHINSTDOC)
	chmod -R 644 $(PATHINSTDOC)*

# Build the RPM package for RedHat-like Linux distributions
.PHONY: rpm
rpm:
	rm -rf $(PATHRPMPKG)
	rpmbuild \
	--define "_topdir $(PATHRPMPKG)" \
	--define "_vendor $(VENDOR)" \
	--define "_owner $(OWNER)" \
	--define "_project $(PROJECT)" \
	--define "_package $(PKGNAME)" \
	--define "_version $(VERSION)" \
	--define "_release $(RELEASE)" \
	--define "_current_directory $(CURRENTDIR)" \
	--define "_libpath /$(LIBPATH)" \
	--define "_docpath /$(DOCPATH)" \
	-bb resources/rpm/rpm.spec

# Build and run the unit tests
.PHONY: test
test: testcpp
	@mkdir -p target/test/test
	@echo -e "\n\n*** BUILD TEST ***\n"
	rm -rf target/test/*
	mkdir -p target/test/coverage
	cd target/test && \
	cmake -DCMAKE_C_FLAGS=$(CMAKE_C_FLAGS) \
	-DCMAKE_TOOLCHAIN_FILE=$(CMAKE_TOOLCHAIN_FILE) \
	-DCMAKE_BUILD_TYPE=Coverage \
	-DCMAKE_INSTALL_PREFIX=$(CMAKE_INSTALL_PATH) \
	-DBUILD_SHARED_LIB=$(VH_BUILD_SHARED_LIB) \
	-DBUILD_DOXYGEN=$(VH_BUILD_DOXYGEN) \
	../.. | tee cmake.log ; test $${PIPESTATUS[0]} -eq 0 && \
	export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./ && \
	make | tee make.log ; test $${PIPESTATUS[0]} -eq 0 && \
	lcov --zerocounters --directory . && \
	lcov --capture --initial --directory . --output-file coverage/numkey.cov && \
	env CTEST_OUTPUT_ON_FAILURE=1 make test | tee test.log ; test $${PIPESTATUS[0]} -eq 0 && \
	lcov --no-checksum --directory . --capture --output-file coverage/numkey.info && \
	lcov --remove coverage/numkey.info '*/test/*' '*/usr/include/x86_64-linux-gnu/bits*' --output-file coverage/numkey.info && \
	genhtml -o coverage -t "NumKey Test Coverage" coverage/numkey.info
ifeq ($(VH_BUILD_DOXYGEN),ON)
	cd target && \
	make doc | tee doc.log ; test $${PIPESTATUS[0]} -eq 0
endif

# Test C code compatibility with C++
.PHONY: testcpp
testcpp:
	find ./src/numkey -type f -name '*.h' -exec gcc -c -pedantic -Werror -Wall -Wextra -Wcast-align -Wundef -Wformat -Wformat-security -std=c++17 -x c++ -o /dev/null {} \;

# use clang-tidy
.PHONY: tidy
tidy:
	clang-tidy -checks='*,-clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling,-readability-function-cognitive-complexity,-altera-struct-pack-align,-altera-id-dependent-backward-branch,-bugprone-easily-swappable-parameters,-altera-unroll-loops,-readability-isolate-declaration,-llvmlibc-restrict-system-libc-headers,-readability-identifier-length,-cppcoreguidelines-avoid-magic-numbers,-readability-magic-numbers,-llvm-header-guard,-llvm-include-order,-android-cloexec-open,-hicpp-no-assembler,-hicpp-signed-bitwise,-clang-analyzer-alpha.*' -header-filter=.* -p . src/numkey/*.h nk/*.c
	clang-tidy -checks='*,-concurrency-mt-unsafe,-clang-analyzer-security.insecureAPI.DeprecatedOrUnsafeBufferHandling,-readability-function-cognitive-complexity,-altera-struct-pack-align,-altera-id-dependent-backward-branch,-bugprone-easily-swappable-parameters,-altera-unroll-loops,-readability-isolate-declaration,-llvmlibc-restrict-system-libc-headers,-readability-identifier-length,-cppcoreguidelines-avoid-magic-numbers,-readability-magic-numbers,-llvm-header-guard,-llvm-include-order,-android-cloexec-open,-hicpp-no-assembler,-hicpp-signed-bitwise,-clang-analyzer-alpha.*' -header-filter=.* -p . test/*.c

# Remove all installed files (excluding configuration files)
.PHONY: uninstall
uninstall:
	rm -rf $(PATHINSTLIB)
	rm -rf $(PATHINSTDOC)

# Set the version from VERSION file
.PHONY: version
version:
	sed -i "s/PROJECT_NUMBER.*=.*$$/PROJECT_NUMBER = $(VERSION)/" doc/Doxyfile
