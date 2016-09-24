package main

import (
	"fmt"
	"github.com/AdamJacobMuller/go-pbkdf2crypt"
	"os"
)

func main() {
	r, _ := pbkdf2crypt.Validate(os.Args[1], os.Args[2])
	if r {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}
}
