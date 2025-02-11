package internal

import "crypto/md5"

// CalculateMD5Hash 用于计算 MD5 哈希（仅内部使用）
func CalculateMD5Hash(data []byte) [16]byte {
	return md5.Sum(data)
}
