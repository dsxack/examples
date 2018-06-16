from ctypes import *

lib = cdll.LoadLibrary("./sum.a")

print "Sum(12,99) = %d" % lib.Sum(12, 99)

