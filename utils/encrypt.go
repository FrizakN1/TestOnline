package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"
)

const encryptKey = "SomeSecretKey"

func Encrypt(value string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(value + encryptKey))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func GenerateBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateString(n int) (string, error) {
	b, err := GenerateBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateHash(value string) (string, error) {
	str, err := GenerateString(16)

	if err != nil {
		return "", err
	}

	hash := sha256.New()
	_, err = hash.Write([]byte(time.Now().String() + str + value + encryptKey))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
