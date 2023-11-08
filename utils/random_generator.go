package utils

import (
	"fmt"
	"math/rand"
)

var alphabeticCharset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomAlphabeticString(size int) string {
	b := make([]byte, size)
	n := len(alphabeticCharset)
	for i := range b {
		b[i] = alphabeticCharset[rand.Intn(n)]
	}

	return string(b)
}

var alphaNumericCharset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomAlphanumericString(size int) string {
	b := make([]byte, size)
	n := len(alphaNumericCharset)
	for i := range b {
		b[i] = alphaNumericCharset[rand.Intn(n)]
	}

	return string(b)
}

func GenerateRandomName() string {
	return generateRandomAlphabeticString(3 + rand.Intn(10))
}

func GenerateRandomEmail() string {
	return fmt.Sprintf(generateRandomAlphanumericString(10) + "@iiitl.ac.in")
}

func GenerateRandomPassword() string {
	return fmt.Sprintf(generateRandomAlphanumericString(8+rand.Intn(10)) + "aA1!")
}
