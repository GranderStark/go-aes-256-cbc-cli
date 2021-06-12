package mainActions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var (
		cipherKey string = ""
		fromStdin string = ""
		toStdout  bool
		fromFile  string = ""
		toFile    string = ""
		encrypted string = ""
		decrypted string = ""
	)
	cipherKey = "abcdefghijklmnopqrstuvwxyz012345"
	fromStdin = "123123"
	toStdout = false
	fromFile = ""
	toFile = ""
	encrypted, _ = RunEncrypt(cipherKey, fromStdin, toStdout, fromFile, toFile)

	decrypted, _ = RunDecrypt(cipherKey, encrypted, toStdout, fromFile, toFile)

	assert.Equal(t, decrypted, fromStdin)

}

func Benchmark(b *testing.B) {
	var (
		cipherKey string = ""
		fromStdin string = ""
		toStdout  bool
		fromFile  string = ""
		toFile    string = ""
		encrypted string = ""
	)
	cipherKey = "abcdefghijklmnopqrstuvwxyz012345"
	fromStdin = "123123"
	for i := 0; i < b.N; i++ {
		encrypted, _ = RunEncrypt(cipherKey, fromStdin, toStdout, fromFile, toFile)
		_, _ = RunDecrypt(cipherKey, encrypted, toStdout, fromFile, toFile)
	}
}
