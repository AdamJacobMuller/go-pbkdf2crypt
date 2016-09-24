#!/usr/bin/python
import pbkdf2
import random
import string
import sys
import sh


go = sh.Command("go")

go("build", "-o", "go-pbkdf2crypt", "go-pbkdf2crypt.go")

godfc = sh.Command('./go-pbkdf2crypt')


def go_validate_pbkdf2(password, crypted):
    r = godfc(password, crypted).strip()
    if r == "true":
        return True
    else:
        return False


def randomthings():
    return ''.join(random.SystemRandom().choice(string.ascii_uppercase + string.digits + string.punctuation) for _ in range(random.randint(1, 100)))

for x in range(0, int(sys.argv[1])):
    password = randomthings()
    crypted = pbkdf2.crypt(password)
    if not go_validate_pbkdf2(password, crypted):
        print "\n%s %s" % (password, crypted)
    else:
        sys.stdout.write(".")
        sys.stdout.flush()
