package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// 获取文件的sha256hex值
func GetFileSha256Hex(f *os.File) (hexs string) {
	hasher := sha256.New()
	io.Copy(hasher, f)
	return hex.EncodeToString(hasher.Sum(nil))
}
