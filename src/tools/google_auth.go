package tools

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

//GetSecret 获取密钥
func GetSecret() string {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, un())
	sha1Code := base32.StdEncoding.EncodeToString(hmacSha1(buf.Bytes(), nil))
	return strings.ToUpper(sha1Code)
}

//GetCode 获取Code
func GetCode(secret string) (string, error) {
	secretUpper := strings.ToUpper(secret)
	secretKey, err := base32.StdEncoding.DecodeString(secretUpper)
	if err != nil {
		return "", err
	}
	number := oneTimePassword(secretKey, toBytes(time.Now().Unix()/30))
	return fmt.Sprintf("%06d", number), nil
}

//VerifyCode 校验Code
func VerifyOTPCode(secret, code string) bool {
	_code, err := GetCode(secret)
	if err != nil {
		return false
	}
	return _code == code
}

//oneTimePassword 获取当前验证码
func oneTimePassword(key []byte, data []byte) uint32 {
	hash := hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := toUint32(hashParts)
	return number % 1000000
}

//toUint32 转Uint32
func toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}
func un() int64 {
	return time.Now().Unix() / 1000 / 30
}

func hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}
