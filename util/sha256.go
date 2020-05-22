package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// 获取文件的sha256hex值
func GetFileSha256Hash(f *os.File) (hexs string) {
	hasher := sha256.New()
	io.Copy(hasher, f)
	return hex.EncodeToString(hasher.Sum(nil))
}

//SHA256生成哈希值
func GetStringSha256Hash(message string) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write([]byte(message))
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode
}
