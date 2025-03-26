<!-- Space: EPT -->
<!-- Parent: Projects -->
<!-- Title: numkey -->

# NumKey

*Numerical Encoding for Short Codes or E.164 LVN*

Also includes:

- *prefixkey*  To encode number prefixes or E.164 LVN as uint64.
- *countrykey* To encode ISO 3166 alpha-2 country code as uint16.

[![check](https://github.com/Vonage/numkey/actions/workflows/check.yaml/badge.svg)](https://github.com/Vonage/numkey/actions/workflows/check.yaml)


* **category**    Library
* **copyright**   2019-2024 Vonage
* **license**     [MIT](https://github.com/Vonage/numkey/blob/main/LICENSE)
* **link**        https://github.com/Vonage/numkey


-----------------------------------------------------------------

## TOC

* [Description](#description)
* [Encoding](#encoding)
    * [Properties](#properties)
* [Quick Start](#quickstart)
* [C Library](#clib)
* [CGO Library](#cgolib)
* [GO Library](#golib)
* [Python Module](#pythonlib)
* [Javascript library](#jslib)
* [Java library](#java)

-----------------------------------------------------------------

<a name="description"></a>
## Description

This library provides functions to encode and decode *Numbers* into a 64 bit unsigned integer.

This document use the word "*Number*" to indicate any of the following:

* **DID** (*Direct Inward Dial*)  : [E.164](https://en.wikipedia.org/wiki/E.164) Telephone Number
* **LVN** (*Long Virtual Number*) : Another definition of DID
* **VLN** (*Virtual Long Number*) : Another definition of DID
* **SC**  (*Short Code*)          : Short number usually containing 4-8 digits

Short Codes (SC) numbers are not universally unique and requires the addition of the country prefix to be uniquely identified.
In other words, the same SC can be used in different countries.

Long numbers do not require a country code to be uniquely identified as the country information is already encoded in the number.
The ITU-T recommendation [E.164](https://en.wikipedia.org/wiki/E.164) establish max 15 digits for a number. 

<a name="encoding"></a>
## Encoding

Modern computers and operating systems are optimized to manipulate 64 bit words, so they are extremely more efficient than using arbitrary strings.

The individual components of a number identifier (country + number) can be easily encoded in 64 bit unsigned integer.

Numbers can be encoded in a reversible way, so we can instantly map a *Number* to a *NumKey* and a *NumKey* to a *Number* without the need of a lookup table.

The encoding is as below:

* 5 + 5 bit to represent each letter of the ISO 3166 alpha-2 country code.
	* Uppercase ASCII characters are transleted back by 64dec, so 1=A and 26=Z.
* 50 bit to encode the number.
* 4 bit to indicate the total number or digits in the number (max 2^4 =16)



The NumKey is composed of 3 sections arranged in 64 bit:


```
           0   4 5                                                            59 60 63
           |   | |                                                             | |  |
           01234 567 89012345 67890123 45678901 2 3456789 01234567 89012345 6789 0123
5 bit COUNTRY >| |<                       50 bit NUMBER                       >| |< 4 bit LENGHT
```

Example of NumKey encoding:

```
               | COUNTRY    | NUMBER                                          | LEN |
---------------+------+-----+-------------------------------------------------+-----+
       Number  |   I     T  | 123456                                          |  6  |
---------------+---- -+-----+-------------------------------------------------+-----+
    NumKey bin | 10011 01000 0000000000000000000000000000000011110001001000000 0110 |
---------------+--------------------------------------------------------------------+
    NumKey hex | 4D000000001E2406                                                   |
    NumKey dec | 5548434740922426374                                                |
---------------+---+----------------------------------------------------------------+
```


* **`COUNTRY FIRST LETTER`** : 5 bit to represent the first letter of the ISO 3166 alpha-2 country code (A=1, ..., Z=26).

    ```
         0   4
         |   |
         11111000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
         |   |
         MSB LSB

         Binary mask: F800000000000000 hex
    ```

* **`COUNTRY SECOND LETTER`** : 5 bit to represent the second letter of the ISO 3166 alpha-2 country code (A=1, ..., Z=26).

    ```
              0    4
              |    |
         00000111 11000000 00000000 00000000 00000000 00000000 00000000 00000000
              |    |
              MSB  LSB

         Binary mask: 0x07C0000000000000 hex
    ```

* **`NUMBER`** : 50 bit to store the number.

    ```
         0         10                                                      59
         |          |                                                      |
         00000000 00111111 11111111 11111111 11111111 11111111 11111111 11110000
                    |                                                      |
                    MSB                                                    LSB

         Binary mask: 0x003FFFFFFFFFFFF0 hex
         Max value:   1125899906842624 (Safe to represent any 15 digit number as per E.164)
    ```

* **`LENGHT`** : 4 bit to store the number of digits.

    ```
         0                                                                 60  63
         |                                                                  |  |
         00000000 00000000 00000000 00000000 00000000 00000000 00000000 00001111
                                                                            |  |
                                                                          MSB  LSB

         Binary mask: 0x000000000000000F hex
    ```

### Long non-standard numbers

If a non-standard number is longer than the E.164 supported 15 digits, than the LENGHT is set to zero and the number is truncated to include only the last 15 digits.
These keys are not directly reversible and require a lookup-table.


<a name="properties"></a>
### Properties

* Each NumKey code is unique for a given Number.
* It can be quickly encoded and decoded on-the-fly.
* Sorting by NumKey is equivalent of sorting by country and number.
* The 64 bit NumKey can be exported as a 16 character hexadecimal string.
* Sorting the hexadecimal representation of NumKey in alphabetical order is equivalent of sorting the NumKey numerically.
* Comparing two Numbers by NumKey only requires comparing two 64 bit numbers, a very well optimized operation in current computer architectures.
* NumKey can be used as a main database key to index data by "Number". This simplify common searching, merging and filtering operations.
* All types of database joins between two data sets (inner, left, right and full) can be easily performed using the NumKey as index.
* Less data storage, less memory usage and increased performances.
* Enable the usage ok key-value systems.
* Enable the use of columnar data formats like *Apache Arrow* with the ability to perform fast binary searches.

----------

<a name="quickstart"></a>
## Quick Start

This project includes a Makefile that allows you to test and build the project in a Linux-compatible system with simple commands.

To see all available options, from the project root type:

```
make help
```

To build all the NumKey versions inside a Docker container (requires Docker):

```
make dbuild
```

An arbitrary make target can be executed inside a [Docker](https://www.docker.com/) container by specifying the `MAKETARGET` parameter:

```
MAKETARGET='build' make dbuild
```
The list of make targets can be obtained by typing ```make```


The base Docker building environment is defined in the following Dockerfile:

```
resources/Docker/Dockerfile.dev
```

To build and test only a specific language version, `cd` into the language directory and use the `make` command.
For example:

```
cd c
make test
```

----------

<a name="clib"></a>
## C Library

* [C Usage Examples](https://github.com/Vonage/numkey/blob/main/c/test/test_example.c)

The reference implementation of this library is written in header-only C programming language in a way that is also compatible with C++.

This project includes a Makefile that allows you to test and build the project in a Linux-compatible system with simple commands.  
All the artifacts and reports produced using this Makefile are stored in the *target* folder.  

* To see all available options: `make help`
* To build everything: `make all`

### Example command-Line tool

The code inside the `c/nk` folder is used to generate the `nk` command line tool.  
This tools requires the positional arguments `COUNTRY`, `NUMBER` and returns the NumKey in hexadecimal representation.


<a name="cgolib"></a>
## Go Library C wrapper

* [Go Usage Examples](https://github.com/Vonage/numkey/blob/main/cgo/example/main.go)

A go wrapper is located in the `cgo` directory.  
Use the "`make cgo`" command to test the GO wrapper and generate reports.


<a name="golib"></a>
## Go Library (golang)

* [Go Usage Examples](https://github.com/Vonage/numkey/blob/main/go/example/main.go)

A native Go implementation is located in the `go` directory.  
Use the "`make go`" command to test the GO version and generate reports.


<a name="pythonlib"></a>
## Python Module

* [Python Usage Examples](https://github.com/Vonage/numkey/blob/main/python/test/example.py)

The python module is located in the `python` directory.
Use the "`make python`" command to test the Python wrapper and generate reports.


<a name="jslib"></a>
## Javascript library

The javascript module is located in the `javascript` directory.
Use the "`make javascript`" command to test and minify the Javascript implementation.


<a name="java"></a>
## Java library

The java module is located in the `java` directory.
Use the "`make java`" command to build and test the Java implementation.
