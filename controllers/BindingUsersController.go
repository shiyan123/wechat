package controllers

import (
	"github.com/astaxie/beego"
)

type BindingUsersController struct {
	beego.Controller
}

func (this *BindingUsersController) Get() {

	this.TplNames = "bindingUsers.html"
}
