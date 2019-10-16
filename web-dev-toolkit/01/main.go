package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Printf("%x\n", hashString("helloworld"))
	fmt.Printf("%x\n", hashString("helloworl"))
}

func hashString(s string) string {
	h := hmac.New(sha256.New, []byte("secret")) // secret key
	h.Write([]byte(s))
	return string(h.Sum(nil))
}
