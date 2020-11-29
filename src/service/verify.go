package service

import (
	"github.com/doter1995/user_center/src/config"
	"github.com/doter1995/user_center/src/model"
	"github.com/doter1995/user_center/src/tools"
	"time"
)

func VerifyMD5(str string, code string) bool {
	slate := config.Config.Security.Token.Slate
	return tools.GetMD5Code(str, slate) == code
}
func NewToken(u model.User) (string, error) {
	token := config.Config.Security.Token
	c := tools.Claims{
		Username: u.Username,
		Status:   u.Status,
	}
	tokenStr, err := tools.GenerateToken(c, token)
	if err != nil {
		return "", err
	}
	if err := redisRecordToken(u.Username, tokenStr); err != nil {
		return "", err
	}
	return tokenStr, nil
}

func redisRecordToken(name string, token string) error {
	st := tools.Rdb.SetEX(tools.RCtx, name, token, time.Hour*2)
	return st.Err()
}

func RedisGetToken(name string) (string, error) {
	st := tools.Rdb.Get(tools.RCtx, name)
	return st.Result()
}

func RedisClearToken(name string) error {
	return tools.Rdb.Del(tools.RCtx, name).Err()
}
