package wechatService

import (
	"fmt"
	"github.com/astaxie/beego"
	coreServlet "wechat/coreServlet"
)

type CoreServletController struct {
	beego.Controller
}

func (this *CoreServletController) Get() {
	signature := this.GetString("signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	echostr := this.GetString("echostr")
	err := coreServlet.CheckSignature(signature, timestamp, nonce)
	if err != nil {
		fmt.Println("connect is fail !!")

	} else {
		fmt.Println("connect is success !!")
		this.Data["echostr"] = echostr
		this.Ctx.WriteString(echostr)
	}
}
func (this *CoreServletController) Post() {
	replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := coreServlet.AcceptMessage(this.Ctx.Input.RequestBody)
	var message string
	if err != nil {
		this.Data["xml"] = message
		this.ServeXml()
	}
	switch {
	case replyTextMessage != nil:
		this.Data["xml"] = replyTextMessage
		this.ServeXml()
	case replyPictureMessage != nil:
		this.Data["xml"] = replyPictureMessage
		this.ServeXml()
	case replyVoiceMessage != nil:
		this.Data["xml"] = replyVoiceMessage
		this.ServeXml()
	case replyVideoMessage != nil:
		this.Data["xml"] = replyVideoMessage
		this.ServeXml()
	case replyMusicMessage != nil:
		this.Data["xml"] = replyMusicMessage
		this.ServeXml()
	case replyNewsTextMessage != nil:
		this.Data["xml"] = replyNewsTextMessage
		this.ServeXml()
	default:
		this.Data["xml"] = message
		this.ServeXml()
	}
}
