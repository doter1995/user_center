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
	var template =
`To: %s
MIME-Version: 1.0
Subject:注册成功
Content-Type: text/html; charset=utf-8
Content-Transfer-Encoding: quoted-printable

<h1>恭喜你已经注册成功<h1></br><p>以下是你authCode，请妥善保存，并使用Google Authenticator保存后作为你登陆验证码生成器的密钥.</p></br></br>
<p>%s</p> </br></br>
<p>或使用Google Authenticator扫描以下二维码自动添加</p></br>
<img src="data:image/jpeg;base64,%s" /></br></br>
`
	to := []string{u.Email}
	str := fmt.Sprintf("otpauth://totp/%s?secret=%s", u.Username, authCode)
	bsData, err := tools.GetQrCodeBase64(str)
	if err != nil {
		tools.Logger.Err(err).Msg("创建二维码base64失败")
		bsData = ""
	}
	msg :=fmt.Sprintf(template,u.Email,authCode,bsData)
	isSend, err := send(to, msg)
	if !isSend {
		tools.Logger.Err(err)
	}
}
