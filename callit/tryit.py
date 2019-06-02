#!/usr/bin/env python3

"""
Example of using simple library created from go using
go build -buildmode=c-shared -o <library filename> <source filename>
"""

from ctypes import *

# Import the library created from go
lib = cdll.LoadLibrary("./go_so.so")

# Below important - otherwise lib.sin returns 0
# Not needed for lib.add though
lib.sin.argtypes = [c_double]
lib.sin.restype = c_double 


# Try lib.add
x = lib.add(1, 2)
print(x)

# Try lib.sin
y = lib.sin(1)

