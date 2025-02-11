package hashutil

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5Hash 计算字符串的 MD5 哈希值
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// SHA256Hash 计算字符串的 SHA-256 哈希值
func SHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
