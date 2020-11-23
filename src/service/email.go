package service

import (
	"fmt"
	"github.com/doter1995/user_center/src/config"
	"github.com/doter1995/user_center/src/model"
	"github.com/doter1995/user_center/src/tools"
	"net/smtp"
)

func send(to []string, msg string) (bool, error) {
	c := config.Config.Email
	auth := smtp.PlainAuth("", c.Account, c.Password, c.Host)
	err := smtp.SendMail(c.Smtp, auth, c.Account, to, []byte(msg))
	if err != nil {
		tools.Logger.Err(err)
		return false, err
	}
	return true, nil
}

func sendRegisterEmail(u model.User, authCode string) {
	to := []string{u.Email}
	str := fmt.Sprintf("otpauth://totp/%s?secret=%s", u.Username, authCode)
	bsData, err := tools.GetQrCodeBase64(str)
	if err != nil {
		tools.Logger.Err(err).Msg("创建二维码base64失败")
		bsData = ""
	}
	msg := "To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"恭喜你已经注册成功.\r\n" +
		"以下是你authCode，请妥善保存，并使用googleAuth保存后作为你登陆验证码生成器的密钥.\r\n\n" +
		authCode + "\r\n\n" +
		"或使用googleAuth扫描以下二维码自动添加" +
		`<img src="data:image/jpeg;base64,` + bsData + `">`
	isSend, err := send(to, msg)
	if !isSend {
		tools.Logger.Err(err)
	}
}
