package wxBaseParameter

type WeChatHttpUrl struct {
	AccessTokenUrl        string
	CustomMenuUrl         string
	MaterialListUrl       string
	SubscribeUsersListUrl string
	SubscribeUsersUrl     string
}

func GetWxChatHttpUrl() *WeChatHttpUrl {
	url := &WeChatHttpUrl{
		AccessTokenUrl:        "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET",
		CustomMenuUrl:         "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN",
		MaterialListUrl:       "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=ACCESS_TOKEN",
		SubscribeUsersListUrl: "https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID",
		SubscribeUsersUrl:     "https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN",
	}
	return url
}
