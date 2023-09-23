package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Characters to include in the password
const (
	letterChars           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars            = "0123456789"
	specialChars          = "!@#$%^&*()_+-=[]{}|;:'<>,.?/~"
	defaultPasswordLength = 12
)

func generateRandomPassword(length int) (string, error) {
	// Define the character set to use
	allChars := letterChars + digitChars + specialChars

	// Create a buffer to store random bytes
	password := make([]byte, length)

	// Generate cryptographically secure random bytes
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}
		password[i] = allChars[randomIndex.Int64()]
	}

	return string(password), nil
}

func main() {
	passwordLength := defaultPasswordLength // Change the desired password length here
	limit := 5
	for i := 0; i < limit; i++ {
		randomPassword, err := generateRandomPassword(passwordLength)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(i+1, ". Random Password ", randomPassword)
	}

}
