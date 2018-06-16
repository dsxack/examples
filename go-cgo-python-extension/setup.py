import os
from setuptools import Extension, setup

extension = Extension(
    "pysum",
    sources = ["pysum.c", ],
    libraries = ["sum"],
    library_dirs = [os.getcwd()]
)

setup(
    ext_modules=[extension],
)
