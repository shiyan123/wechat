package routers

import (
	"github.com/astaxie/beego"
	"wechat/controllers"
	"wechat/controllers/wechatService"
)

func init() {
	beego.Router("login", &controllers.LoginController{})
	beego.Router("index", &controllers.IndexController{})
	beego.Router("accountCenter", &controllers.AccountCenterController{})
	beego.Router("customMenu", &controllers.CustomMenuController{})
	beego.Router("subscribeWelcome", &controllers.SubscribeWelcomeController{})
	beego.Router("keyWord", &controllers.KeyWordController{})
	beego.Router("notKnowReply", &controllers.NotKnowReplyController{})
	beego.Router("subscribeUsers", &controllers.SubscribeUsersController{})
	beego.Router("bindingUsers", &controllers.BindingUsersController{})
	beego.Router("wxCore", &wechatService.CoreServletController{})
	beego.Router("textMessage", &controllers.TextMessageController{})
	beego.Router("newsMessage", &controllers.NewsMessageController{})
	beego.Router("newsMessageList", &controllers.NewsMessageListController{})
}
