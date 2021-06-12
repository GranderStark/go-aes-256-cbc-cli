package mainActions

import (
	"github.com/GranderStark/go-aes-256-cbc-cli/internal/decrypt"
	"github.com/GranderStark/go-aes-256-cbc-cli/internal/encrypt"
	"github.com/GranderStark/go-aes-256-cbc-cli/internal/utils"
)

func RunEncrypt(cipherKey string, fromStdin string, toStdout bool, fromFile string, toFile string) (string, error) {
	var (
		encryptValue     string = ""
		encryptionResult string = ""
		err              error  = nil
		valueFromFile    string = ""
		wroteStdout      bool   = false
	)

	encryptValue = utils.GetFromStdin(fromStdin)

	valueFromFile = utils.GetFromFile(fromFile)
	if valueFromFile != "" {
		encryptValue = valueFromFile
	}

	encryptionResult, err = encrypt.Encrypt(encryptValue, cipherKey)
	utils.Check(err)

	wroteStdout = utils.WriteToStdout(toStdout, encryptionResult)
	if wroteStdout == true {
		return encryptionResult, nil
	}
	utils.WriteToFile(toFile, encryptionResult)

	return encryptionResult, nil
}

func RunDecrypt(cipherKey string, fromStdin string, toStdout bool, fromFile string, toFile string) (string, error) {
	var (
		decryptValue     string = ""
		decryptionResult string = ""
		err              error  = nil
		valueFromFile    string = ""
		wroteStdout      bool   = false
	)
	decryptValue = utils.GetFromStdin(fromStdin)
	valueFromFile = utils.GetFromFile(fromFile)
	if valueFromFile != "" {
		decryptValue = valueFromFile
	}

	decryptionResult, err = decrypt.Decrypt(decryptValue, cipherKey)
	utils.Check(err)

	wroteStdout = utils.WriteToStdout(toStdout, decryptionResult)
	if wroteStdout == true {
		return decryptionResult, nil
	}

	utils.WriteToFile(toFile, decryptionResult)

	return decryptionResult, nil
}
