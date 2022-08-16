package tools

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/itchyny/base58-go"
	"golang.org/x/crypto/bcrypt"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

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
