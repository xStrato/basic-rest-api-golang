package services

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Encrypt(s *string) {
	encrypted := sha256.Sum256([]byte(*s))
	*s = fmt.Sprintf("%x", encrypted)
}
