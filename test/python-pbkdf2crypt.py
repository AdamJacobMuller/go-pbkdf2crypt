#!/usr/bin/python
import pbkdf2
import sys

print "%s" % pbkdf2.crypt(sys.argv[1], sys.argv[2])
