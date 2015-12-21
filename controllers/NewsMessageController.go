package controllers

import (
	"github.com/astaxie/beego"
)

type NewsMessageController struct {
	beego.Controller
}

func (this *NewsMessageController) Get() {
	this.TplNames = "newsMessage.html"
}
