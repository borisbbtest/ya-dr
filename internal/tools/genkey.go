package tools

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
	uuid "github.com/nu7hatch/gouuid"
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

func GenerateKeyGen(initialLink string) string {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	urlHashBytes := sha256Of(initialLink + u.String())
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
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
