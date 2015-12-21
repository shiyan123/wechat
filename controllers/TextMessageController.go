package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"wechat/models/octopus"
)

type TextMessageController struct {
	beego.Controller
}

func (this *TextMessageController) Get() {
	textMessages, err := octopus.FindAllTextMessage()
	if err != nil {
		this.TplNames = "textMessage.html"
	}
	this.Data["TextMessage"] = textMessages
	this.TplNames = "textMessage.html"
}
func (this *TextMessageController) Post() {
	messageTitle := this.GetString("title")
	messageContent := this.GetString("content")
	_, err := octopus.InsertOneTextMessage(messageTitle, messageContent)
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
func (this *TextMessageController) Delete() {
	data := this.Ctx.Input.RequestBody
	messageTitles := make(map[string]string)
	err := json.Unmarshal(data, &messageTitles)
	if err != nil {
		message := &Message{
			Message: "操作失败",
		}
		this.Data["json"] = message
		this.ServeJson()
	}
	err = octopus.DeleteTextMessage(messageTitles)
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
