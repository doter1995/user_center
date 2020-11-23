package tools

import (
	"crypto/md5"
	"encoding/hex"
)

//GetMD5Code 获取MD5Code
func GetMD5Code(password string, slat string) string {
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(slat))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}
