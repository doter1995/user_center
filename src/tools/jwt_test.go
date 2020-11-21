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
	token, err := GenerateToken(c)
	assert.NoError(t, err)
	t.Logf("密钥: %s", token)
	c, err = VerifyToken(token)
	assert.NoError(t, err)
	assert.Equal(t, c.Username, "username")
	assert.Equal(t, c.Status, 1)
}
