package util

import (
	"fmt"
	"os"
	"testing"
)

func TestGetFileSha256Hex(t *testing.T) {
	// openssl dgst -sha256 22.png
	// 8b37381b4b8d960265ca7c1a4f2bc2cef00057bde636d512c2848e90bfe84cac
	file, _ := os.Open("/Users/silence/Downloads/goProject/src/tmp/22.png")

	// 8b37381b4b8d960265ca7c1a4f2bc2cef00057bde636d512c2848e90bfe84cac
	sha256Hash := GetFileSha256Hash(file)
	fmt.Println(sha256Hash, len(sha256Hash))
}
