package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func EncryptAES256GCM(plaintext []byte, key []byte) ([]byte, error) {
	k := DeriveKey32B(key)
	block, err := aes.NewCipher(k[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, plaintext, nil)

	return cipherText, nil
}
