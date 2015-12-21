package controllers

import (
	"github.com/astaxie/beego"
)

type SubscribeUsersController struct {
	beego.Controller
}

func (this *SubscribeUsersController) Get() {

	this.TplNames = "subscribeUsers.html"
}
