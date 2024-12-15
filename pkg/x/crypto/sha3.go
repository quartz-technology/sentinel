package xcrypto

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func Keccak256Hash(data string) (string, error) {
	hash := sha3.NewLegacyKeccak256()
	_, err := hash.Write([]byte(data))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("0x%x", hash.Sum(nil)), nil
}
