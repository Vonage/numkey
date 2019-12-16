// NumKey Python wrapper
//
// pynumkey.h
//
// @category   Libraries
// @author     Nicola Asuni <nicola.asuni@vonage.com>
// @copyright  2019 Vonage
// @license    see LICENSE file
// @link       https://github.com/nexmoinc/numkey

#define PY_SSIZE_T_CLEAN
#include <Python.h>

// NUMKEY
static PyObject *py_numkey(PyObject *self, PyObject *args, PyObject *keywds);
static PyObject *py_decode_numkey(PyObject *self, PyObject *args, PyObject *keywds);
static PyObject *py_compare_numkey_country(PyObject *self, PyObject *args, PyObject *keywds);
static PyObject *py_numkey_hex(PyObject *self, PyObject *args, PyObject *keywds);
static PyObject *py_parse_numkey_hex(PyObject *self, PyObject *args, PyObject *keywds);

PyMODINIT_FUNC initnumkey(void);

// NUMKEY

#define PYNUMKEY_DOCSTRING "Returns a 64 bit number key based on COUNTRY and NUMBER.\n"\
"\n"\
"Parameters\n"\
"----------\n"\
"country : str or bytes\n"                                             \
"    Country. A ISO 3166 alpha-2 country code.\n"                      \
"number : str or bytes\n"                                              \
"    Number. String containing the Short code or E.164 LVN number.\n"  \
"\n"\
"Returns\n"\
"-------\n"\
"int:\n"\
"    NumKey 64 bit code.\n"\
"\n"\
"Example\n"\
"-------\n"\
">>> numkey(country=b'XJ', number=b'762942138198343')\n"               \
"14027409114588157055"

#define PYDECODENUMKEY_DOCSTRING "Decode a NumKey code and returns the components.\n"\
"\n"\
"Parameters\n"\
"----------\n"\
"nk : int\n"                                                           \
"    NumKey code.\n"\
"\n"\
"Returns\n"\
"-------\n"\
"tuple string: \n"                                                     \
"    - COUNTRY string\n"                                               \
"    - NUMBER string\n"                                                \
"\n"\
"Example\n"\
"-------\n"\
">>> decode_numkey(14027409114588157055)\n"                            \
"(b'XJ', b'762942138198343')"

#define PYCOMPARENUMKEYCOUNTRY_DOCSTRING "Compares two NumKeys by country only.\n"\
"\n"\
"Parameters\n"\
"----------\n"\
"nka : int\n"                                                          \
"    The first NumKey to be compared.\n"\
"nkb : int\n"                                                          \
"    The second NumKey to be compared.\n"\
"\n"\
"Returns\n"\
"-------\n"\
"int :\n"\
"    -1 if the first country is smaller than the second, 0 if they are equal and 1 if the first is greater than the second.\n"\
"\n"\
"Example\n"\
"-------\n"\
">>> compare_numkey_country(0xd6a23089b8e15cdf, 0xd6a2300000000000)\n" \
"0"

#define PYNUMKEYHEX_DOCSTRING "Returns NumKey hexadecimal string (16 characters).\n"\
"\n"\
"Parameters\n"\
"----------\n"\
"nk : int\n"                                                           \
"    NumKey code.\n"\
"\n"\
"Returns\n"\
"-------\n"\
"bytes:\n"\
"    NumKey hexadecimal string.\n"\
"\n"\
"Example\n"\
"-------\n"\
">>> numkey_hex(14027409114588157055)\n"                               \
"b'c2ab5e44f21a947f'"

#define PYPARSENUMKEYSTRING_DOCSTRING "Parses a NumKey hexadecimal string and returns the code.\n"\
"\n"\
"Parameters\n"\
"----------\n"\
"ns : str or bytes\n"                                                  \
"    NumKey hexadecimal string (it must contain 16 hexadecimal characters).\n"\
"\n"\
"Returns\n"\
"-------\n"\
"int :\n"\
"    NumKey 64 bit code.\n"\
"\n"\
"Example\n"\
"-------\n"\
">>> parse_numkey_hex(b'c2ab5e44f21a947f')\n"                          \
"14027409114588157055"


#if defined(__SUNPRO_C) || defined(__hpux) || defined(_AIX)
#define inline
#endif

#ifdef __linux
#define inline __inline
#endif
