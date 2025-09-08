//SHA256

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func mainSha() {
	h := sha256.New()
	h.Write([]byte("hello, world!"))
	fmt.Printf("%x\n", h.Sum(nil))
}

// Hmac
func mainHmac() {
	key := []byte("password")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte("hello, world!"))
	fmt.Printf("%x\n", mac.Sum(nil))
}
