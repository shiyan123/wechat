package controllers

import (
	"github.com/astaxie/beego"
	"wechat/models/octopus"
)

type AccountCenterController struct {
	beego.Controller
}
type Message struct {
	Message string
}

func (this *AccountCenterController) Get() {
	accounts, err := octopus.FindAllAccount()
	if err != nil {
		this.TplNames = "accountCenter.html"
	}
	this.Data["Accounts"] = accounts
	this.TplNames = "accountCenter.html"

}

func (this *AccountCenterController) Post() {
	accountName := this.GetString("accountName")
	token := this.GetString("token")
	wechatNumber := this.GetString("wechatNumber")
	oldId := this.GetString("oldId")
	accountType := this.GetString("accountType")
	email := this.GetString("email")
	description := this.GetString("description")
	appid := this.GetString("appid")
	appsecret := this.GetString("appsecret")
	err := octopus.InsertOneWechatAccount(accountName, token, wechatNumber, oldId, accountType, email, description, appid, appsecret)
	if err != nil {
		message := &Message{
			Message: "操作失败",
		}
		this.Data["json"] = message
		this.ServeJson()
	}
	message := &Message{
		Message: "操作成功",
	}
	this.Data["json"] = message
	this.ServeJson()
}
