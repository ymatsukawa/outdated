package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func DecryptAES256GCM(cipherText []byte, key []byte) ([]byte, error) {
	k := DeriveKey32B(key)
	block, err := aes.NewCipher(k[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nSize := gcm.NonceSize()
	nonce, encryptedData := cipherText[:nSize], cipherText[nSize:]

	plainText, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, fmt.Errorf("key may be wrong.\n\tdetails `%w`", err)
	}

	return plainText, nil
}
