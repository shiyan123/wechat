package wxBaseParameter

type WeChatHttpUrl struct {
	AccessTokenUrl string
	CustomMenuUrl  string
}

func GetWxChatHttpUrl() *WeChatHttpUrl {
	url := &WeChatHttpUrl{
		AccessTokenUrl: "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET",
		CustomMenuUrl:  "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN",
	}
	return url
}
