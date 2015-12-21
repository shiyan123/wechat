package controllers

import (
	"github.com/astaxie/beego"
)

type NotKnowReplyController struct {
	beego.Controller
}

func (this *NotKnowReplyController) Get() {

	this.TplNames = "notKnowReply.html"
}
