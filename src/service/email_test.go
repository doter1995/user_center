package service

import (
	"github.com/doter1995/user_center/src/config"
	"github.com/doter1995/user_center/src/model"
	"github.com/doter1995/user_center/src/tools"
	"testing"
)

func Test_sendRegisterEmail(t *testing.T) {
	config.InitConfig("../../config.yaml")
	tools.InitLog()
	u := model.User{
		Username: "doter1995",
		Email:    "wdzhang@thoughtworks.com",
		AuthCode: "MPNP3K3LBG4EUGDZMZUOIODSKR2PXBM6",
	}
	sendRegisterEmail(u, "MPNP3K3LBG4EUGDZMZUOIODSKR2PXBM6")

}
