//
// https://www.golangprograms.com/hash-based-message-authentication-code-hmac.html
//

package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

var secretKey = "4234kxzjcjj3nxnxbcvsjfj"

// Generate a salt string with 16 bytes of crypto/rand data.
func generateSalt() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func main() {

	// This is the sample message
	message := "This is a sample message that will be Encrypted with HMAC-SHA512."

	// We need to create a new Salt value.
	salt := generateSalt()

	// Print the message
	fmt.Println("Message: " + message)

	fmt.Println("--  This is the Salt value: ")

	// Print the Salt value
	fmt.Println("\nSalt: " + salt)

	// New returns a new HMAC hash using the given hash.
	//   func hmac.New(h func() hash.Hash, key []byte) hash.Hash
	hash := hmac.New(sha256.New, []byte(secretKey)) // SHA256
	fmt.Println(hash)

	// writes the contents of the string s to w, which accepts a slice of bytes.
	io.WriteString(hash, message+salt)

	fmt.Println("--  This is the calculated Hash value: ")
	fmt.Printf("\nHMAC-Sha256: %x", hash.Sum(nil))

	// New returns a new HMAC hash using the given hash.   SHA512
	hash = hmac.New(sha512.New, []byte(secretKey))

	//
	io.WriteString(hash, message+salt)

	//
	fmt.Printf("\n\nHMAC-sha512: %x", hash.Sum(nil))
}
