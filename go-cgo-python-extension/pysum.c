#include <Python.h>
#include "./libsum.h"

static PyObject *
method_sum(PyObject *self, PyObject *args)
{
    const int a;
    const int b;

    /* Parse the input tuple */
    if (!PyArg_ParseTuple(args, "ii", &a, &b))
        return NULL;

    PyObject *ret = Py_BuildValue("i", Sum(a, b));
    return ret;
}

static PyMethodDef PysumMethods[] = {
    {"sum", method_sum, METH_VARARGS, NULL},
    {NULL, NULL, 0, NULL}
};

static PyModuleDef pysummodule = {
    PyModuleDef_HEAD_INIT,
    "pysum",            /* name of module */
    NULL,                 /* module documentation, may be NULL */
    -1,                 /* size of per-interpreter state of the module,
                         or -1 if the module keeps state in global variables. */
    PysumMethods
};

PyMODINIT_FUNC PyInit_pysum(void)
{
    return PyModule_Create(&pysummodule);
}
