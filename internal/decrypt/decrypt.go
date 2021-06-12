// origin https://gist.github.com/brettscott/2ac58ab7cb1c66e2b4a32d6c1c3908a7
package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"github.com/mergermarket/go-pkcs7"
)

// Decrypt decrypts cipher text string into plain text string
func Decrypt(encrypted string, cipherKey string) (string, error) {
	key := []byte(cipherKey)
	cipherText, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(cipherText) < aes.BlockSize {
		panic("cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		panic("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	return fmt.Sprintf("%s", cipherText), nil
}
