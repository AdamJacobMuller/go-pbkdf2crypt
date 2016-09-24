package main

import (
	"fmt"
	"github.com/AdamJacobMuller/go-pbkdf2crypt"
	"os"
)

func main() {
	fmt.Printf("%#v", os.Args)
	r, e := pbkdf2crypt.Validate(os.Args[1], os.Args[2])
	fmt.Printf("r = %s\n", r)
	fmt.Printf("e = %s\n", e)
}
