package main

import (
	"testing"
)

func TestGenerateRandomPassword(t *testing.T) {
	length := 12

	t.Run("Generated password has the correct length", func(t *testing.T) {
		password, err := generateRandomPassword(length)
		if err != nil {
			t.Fatalf("Error generating password: %v", err)
		}
		if len(password) != length {
			t.Errorf("Expected password of length %d, but got length %d", length, len(password))
		}
	})

	t.Run("Generated password contains at least one letter", func(t *testing.T) {
		password, err := generateRandomPassword(length)
		if err != nil {
			t.Fatalf("Error generating password: %v", err)
		}
		hasLetter := false
		for _, char := range password {
			if isLetter(char) {
				hasLetter = true
				break
			}
		}
		if !hasLetter {
			t.Errorf("Generated password does not contain any letters")
		}
	})

	t.Run("Generated password contains at least one digit", func(t *testing.T) {
		password, err := generateRandomPassword(length)
		if err != nil {
			t.Fatalf("Error generating password: %v", err)
		}
		hasDigit := false
		for _, char := range password {
			if isDigit(char) {
				hasDigit = true
				break
			}
		}
		if !hasDigit {
			t.Errorf("Generated password does not contain any digits")
		}
	})

	t.Run("Generated password contains at least one special character", func(t *testing.T) {
		password, err := generateRandomPassword(length)
		if err != nil {
			t.Fatalf("Error generating password: %v", err)
		}
		hasSpecial := false
		for _, char := range password {
			if isSpecial(char) {
				hasSpecial = true
				break
			}
		}
		if !hasSpecial {
			t.Errorf("Generated password does not contain any special characters")
		}
	})
}

func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isSpecial(char rune) bool {
	specialChars := "!@#$%^&*()_+-=[]{}|;:'<>,.?/~"
	for _, s := range specialChars {
		if char == s {
			return true
		}
	}
	return false
}
