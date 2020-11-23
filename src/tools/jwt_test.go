package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateToken(t *testing.T) {
	c := Claims{
		Username: "username",
		Status:   1,
	}
	config := struct {
		Auth  string
		Slate string
	}{
		Auth:  "doter1995",
		Slate: "test",
	}
	token, err := GenerateToken(c, config)
	assert.NoError(t, err)
	t.Logf("密钥: %s", token)
	c, err = VerifyToken(token, config)
	assert.NoError(t, err)
	assert.Equal(t, c.Username, "username")
	assert.Equal(t, c.Status, 1)
}
