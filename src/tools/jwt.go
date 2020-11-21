package tools

import (
	"errors"
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

//Claims Token内容
type Claims struct {
	Username string `json:"username"`
	Status   int    `json:"status"`
	AuthCode string `json:"authCode"`
	jwt.StandardClaims
}

//GenerateToken 生成Token
func GenerateToken(c Claims) (string, error) {
	c.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(), // 过期时间
		IssuedAt:  time.Now().Unix(),                    // 颁发时间
		Issuer:    "doter1995",                          //颁发人
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte("test"))
	if err != nil {
		return "", errors.New("jwt: token signing failed: " + err.Error())
	}

	return tokenString, nil
}

//VerifyToken 校验
func VerifyToken(tokenString string) (Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("jwt: unexpected signing method")
			}
			return []byte("test"), nil
		})
	if err != nil {
		return Claims{}, errors.New("jwt: ParseWithClaims failed: " + err.Error())
	}
	if !token.Valid {
		return Claims{}, errors.New("jwt: token is not valid")
	}

	c, ok := token.Claims.(*Claims)
	if !ok {
		return Claims{}, errors.New("jwt: failed to get token claims")
	}

	if c.Username == "" {
		return Claims{}, errors.New("jwt: UserID claim is not valid")
	}

	if c.IssuedAt == 0 {
		return Claims{}, errors.New("jwt: IssuedAt claim is not valid")
	}
	return *c, nil
}
