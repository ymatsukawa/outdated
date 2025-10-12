package crypto

import (
	"crypto/sha256"
)

func DeriveKey32B(rawKey []byte) []byte {
	k := sha256.Sum256(rawKey)
	return k[:]
}
