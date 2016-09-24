package pbkdf2crypt

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"strconv"
	"strings"
)

func d(b []byte, n string) {
	fmt.Printf("%s (%d) = %#v (%s)\n", n, len(b), b, b)
}

func Validate(password string, hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	if len(parts) != 5 {
		return false, errors.New("pbkdf2crypt should have 5 parts")
	}
	if parts[1] != "p5k2" {
		return false, errors.New("hash is not a pbkdf2crypt")
	}
	var iterations int
	var salt []byte
	if parts[2] == "" {
		iterations = 400
		salt = []byte(fmt.Sprintf("$p5k2$$%s", parts[3]))
	} else {
		iterations64, _ := strconv.ParseInt(parts[2], 10, 32)
		iterations = int(iterations64)
		salt = []byte(fmt.Sprintf("$p5k2$%x$%s", iterations, parts[3]))
	}

	rk := pbkdf2.Key([]byte(password), salt, iterations, 24, sha1.New)

	buf := new(bytes.Buffer)
	pythonEncoding := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789./"

	encoder := base64.NewEncoder(base64.NewEncoding(pythonEncoding), buf)
	encoder.Write(rk)
	encoder.Close()

	result := buf.String()
	if result == parts[4] {
		return true, nil
	} else {
		return false, nil
	}
	return false, nil
}
