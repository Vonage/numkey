#!/usr/bin/env python3

from codecs import open
from os.path import dirname, join
from subprocess import call
from setuptools import setup, find_packages, Extension, Command


def read(fname):
    return open(join(dirname(__file__), fname)).read()


class RunTests(Command):
    """Run all tests."""

    description = "run tests"
    user_options = []

    def initialize_options(self):
        pass

    def finalize_options(self):
        pass

    def run(self):
        """Run all tests!"""
        errno = call(["py.test", "--verbose"])
        raise SystemExit(errno)


setup(
    name="numkey",
    version="1.6.43.1",
    keywords=("numkey E.164 shortcode lvn did encoding"),
    description="NumKey Bindings for Python",
    long_description=read("../README.md"),
    author="Nicola Asuni",
    author_email="api.eng.cet@vonage.com ",
    url="https://github.com/Vonage/numkey",
    license="RESERVED",
    platforms="Linux",
    packages=find_packages(exclude=["doc", "test*"]),
    ext_modules=[
        Extension(
            "numkey",
            ["numkey/pynumkey.c"],
            include_dirs=["../c/src/numkey", "numkey"],
            extra_compile_args=[
                "-O3",
                "-pedantic",
                "-std=c99",
                "-Wall",
                "-Wextra",
                "-Wno-strict-prototypes",
                "-Wunused-value",
                "-Wcast-align",
                "-Wundef",
                "-Wformat",
                "-Wformat-security",
                "-Wshadow",
                "-Wno-format-overflow",
                "-I../c/src/numkey",
            ],
        )
    ],
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "Intended Audience :: Developers",
        "Programming Language :: C",
        "Programming Language :: Python",
    ],
    extras_require={
        "test": [
            "coverage",
            "pytest",
            "pytest-benchmark",
            "pytest-cov",
            "pycodestyle",
            "black",
        ]
    },
    cmdclass={"test": RunTests},
)
