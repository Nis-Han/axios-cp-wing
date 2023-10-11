package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"github.com/nerd500/axios-cp-wing/internal/database"
)

func GenerateSalt() string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(salt)
}

func HashPassword(password, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CheckPassword(user database.User, password string) bool {
	return user.HashedPassword == HashPassword(password, user.Salt)
}

func GenerateAuthToken(length int) string {
	tokenBytes := make([]byte, length)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(tokenBytes)
}
