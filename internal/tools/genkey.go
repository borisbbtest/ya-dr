package tools

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GetID() (res []byte, err error) {
	src := []byte("Ключ от сердца")

	key, err := GenerateRandom(aes.BlockSize)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	res = make([]byte, aes.BlockSize)
	aesblock.Encrypt(res, src)
	return
}
