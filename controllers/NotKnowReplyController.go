package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"wechat/models/octopus"
)

type NotKnowReplyController struct {
	beego.Controller
}

func (this *NotKnowReplyController) Get() {
	textMessages, err := octopus.FindAllTextMessage()
	if err != nil {
		this.TplNames = "notKnowReply.html"
	}
	this.Data["TextMessage"] = textMessages
	newsMessagesLists, err, _ := octopus.FindAllNewsMessageList()
	if err != nil {
		this.TplNames = "notKnowReply.html"
	}
	this.Data["NewsMessagesLists"] = newsMessagesLists
	notKnowReplys, err := octopus.FindNotKnowReplys()
	if err != nil {
		this.TplNames = "notKnowReply.html"
	}
	this.Data["NotKnowReplys"] = notKnowReplys
	this.TplNames = "notKnowReply.html"
}
func (this *NotKnowReplyController) Post() {
	modelMessage := this.GetString("modelMessage")
	messageType := this.GetString("type")
	_, err := octopus.InsertOneNotKnowReply(modelMessage, messageType)
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
func (this *NotKnowReplyController) Delete() {
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
	err = octopus.DeleteNotKnowReply(messageTitles)
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
