package controllers

import (
	"github.com/astaxie/beego"
	"wechat/models/octopus"
)

type SubscribeUsersController struct {
	beego.Controller
}

func (this *SubscribeUsersController) Get() {
	subscribeUsers, err := octopus.FindAllSubscribeUsers()
	if err != nil {
		this.TplNames = "subscribeUsers.html"
	}
	this.Data["SubscribeUsers"] = subscribeUsers
	this.TplNames = "subscribeUsers.html"
}
func (this *SubscribeUsersController) Put() {
	a := ""
	err := octopus.GetSubscribeUserList(a)
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
