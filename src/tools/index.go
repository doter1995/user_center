package tools

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
)

func GetQrCodeBase64(msg string) (string, error) {
	data, err := qrcode.Encode(msg, qrcode.Medium, 256)
	return base64.StdEncoding.EncodeToString(data), err
}
