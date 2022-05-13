// NumKey Python wrapper
//
// pynumkey.c
//
// @category   Libraries
// @author     Nicola Asuni <nicola.asuni@vonage.com>
// @copyright  2019-2022 Vonage
// @license    see LICENSE file
// @link       https://github.com/nexmoinc/numkey

#define MODULE_NAME "numkey"
#define PY_SSIZE_T_CLEAN  //!< Make "s#" use Py_ssize_t rather than int.

#include <Python.h>
#include "../../c/src/numkey/hex.h"
#include "../../c/src/numkey/numkey.h"
#include "../../c/src/numkey/prefixkey.h"
#include "pynumkey.h"

#ifndef Py_UNUSED // This is already defined for Python 3.4 onwards
#ifdef __GNUC__
#define Py_UNUSED(name) _unused_ ## name __attribute__((unused))
#else
#define Py_UNUSED(name) _unused_ ## name
#endif
#endif

static PyObject* py_numkey(PyObject *Py_UNUSED(ignored), PyObject *args, PyObject *keywds)
{
    const char *country, *number;
    Py_ssize_t sizecountry, sizenumber;
    static char *kwlist[] = {"country", "number", NULL};
    if (!PyArg_ParseTupleAndKeywords(args, keywds, "s#s#", kwlist, &country, &sizecountry, &number, &sizenumber))
        return NULL;
    if (sizecountry != 2 || sizenumber < 1)
    {
        return Py_BuildValue("K", 0);
    }
    uint64_t h = numkey(country, number, (size_t)sizenumber);
    return Py_BuildValue("K", h);
}

static PyObject* py_decode_numkey(PyObject *Py_UNUSED(ignored), PyObject *args, PyObject *keywds)
{
    uint64_t nk;
    static char *kwlist[] = {"nk", NULL};
    if (!PyArg_ParseTupleAndKeywords(args, keywds, "K", kwlist, &nk))
        return NULL;
    numkey_t h = {0};
    decode_numkey(nk, &h);
    PyObject *result = PyTuple_New(3);
    PyTuple_SetItem(result, 0, Py_BuildValue("y", h.country));
    PyTuple_SetItem(result, 1, Py_BuildValue("y", h.number));
    return result;
}

static PyObject* py_compare_numkey_country(PyObject *Py_UNUSED(ignored), PyObject *args, PyObject *keywds)
{
    uint64_t nka, nkb;
    static char *kwlist[] = {"nka", "nkb", NULL};
    if (!PyArg_ParseTupleAndKeywords(args, keywds, "KK", kwlist, &nka, &nkb))
        return NULL;
    int cmp = compare_numkey_country(nka, nkb);
    return Py_BuildValue("i", cmp);
}

static PyObject* py_numkey_hex(PyObject *Py_UNUSED(ignored), PyObject *args, PyObject *keywds)
{
    uint64_t code;
    static char *kwlist[] = {"nk", NULL};
    if (!PyArg_ParseTupleAndKeywords(args, keywds, "K", kwlist, &code))
        return NULL;
    char str[17];
    numkey_hex(code, str);
    return PyBytes_FromString(str);
}

static PyObject* py_parse_numkey_hex(PyObject *Py_UNUSED(ignored), PyObject *args, PyObject *keywds)
{
    const char *ns;
    Py_ssize_t sizens;
    static char *kwlist[] = {"ns", NULL};
    if (!PyArg_ParseTupleAndKeywords(args, keywds, "s#", kwlist, &ns, &sizens))
        return NULL;
    uint64_t h = 0;
    if (sizens == 16)
    {
        h = parse_numkey_hex(ns);
    }
    return Py_BuildValue("K", h);
}

static PyObject* py_prefixkey(PyObject *Py_UNUSED(ignored), PyObject *args, PyObject *keywds)
{
    const char *number;
    Py_ssize_t sizenumber;
    static char *kwlist[] = {"number", NULL};
    if (!PyArg_ParseTupleAndKeywords(args, keywds, "s#", kwlist, &number, &sizenumber))
        return NULL;
    if (sizenumber < 1)
    {
        return Py_BuildValue("K", 0);
    }
    uint64_t h = prefixkey(number, (size_t)sizenumber);
    return Py_BuildValue("K", h);
}

// ---

static PyMethodDef PyNumKeyMethods[] =
{
    // NUMKEY
    {"numkey", (PyCFunction)py_numkey, METH_VARARGS|METH_KEYWORDS, PYNUMKEY_DOCSTRING},
    {"decode_numkey", (PyCFunction)py_decode_numkey, METH_VARARGS|METH_KEYWORDS, PYDECODENUMKEY_DOCSTRING},
    {"compare_numkey_country", (PyCFunction)py_compare_numkey_country, METH_VARARGS|METH_KEYWORDS, PYCOMPARENUMKEYCOUNTRY_DOCSTRING},
    {"numkey_hex", (PyCFunction)py_numkey_hex, METH_VARARGS|METH_KEYWORDS, PYNUMKEYHEX_DOCSTRING},
    {"parse_numkey_hex", (PyCFunction)py_parse_numkey_hex, METH_VARARGS|METH_KEYWORDS, PYPARSENUMKEYSTRING_DOCSTRING},
    {"prefixkey", (PyCFunction)py_prefixkey, METH_VARARGS|METH_KEYWORDS, PYPREFIXKEY_DOCSTRING},
    {NULL, NULL, 0, NULL}
};

static const char modulename[] = MODULE_NAME;

struct module_state
{
    PyObject *error;
};

#if PY_MAJOR_VERSION >= 3
#define GETSTATE(m) ((struct module_state*)PyModule_GetState(m))
#else
#define GETSTATE(m) (&_state)
static struct module_state _state;
#endif

#if PY_MAJOR_VERSION >= 3
static int numkey_traverse(PyObject *m, visitproc visit, void *arg)
{
    Py_VISIT(GETSTATE(m)->error);
    return 0;
}

static int numkey_clear(PyObject *m)
{
    Py_CLEAR(GETSTATE(m)->error);
    return 0;
}

static struct PyModuleDef moduledef =
{
    PyModuleDef_HEAD_INIT,
    modulename,
    NULL,
    sizeof(struct module_state),
    PyNumKeyMethods,
    NULL,
    numkey_traverse,
    numkey_clear,
    NULL
};

#define INITERROR return NULL

PyObject* PyInit_numkey(void)
#else
#define INITERROR return

void initnumkey(void)
#endif
{
#if PY_MAJOR_VERSION >= 3
    PyObject *module = PyModule_Create(&moduledef);
#else
    PyObject *module = Py_InitModule(modulename, PyNumKeyMethods);
#endif
    struct module_state *st = NULL;
    if (module == NULL)
    {
        INITERROR;
    }
    st = GETSTATE(module);
    st->error = PyErr_NewException(MODULE_NAME ".Error", NULL, NULL);
    if (st->error == NULL)
    {
        Py_DECREF(module);
        INITERROR;
    }
#if PY_MAJOR_VERSION >= 3
    return module;
#endif
}
