package tools

import "testing"

func Test_getQrCodeBase64(t *testing.T) {
	data, _ := GetQrCodeBase64("doter1995")
	t.Logf("密钥: %s", data)
}
