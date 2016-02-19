package wxMessageService

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"wechat/models/octopus"
)

type ReplyTextMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}
type ReplyPictureMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}
type ReplyVoiceMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}
type ReplyVideoMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}
type ReplyMusicMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}
type ReplyNewsTextMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int      `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	ArticleCount int      `xml:"ArticleCount"`
	articles     Articles `xml:"Articles"`
}
type Articles struct {
	item []Item `xml:"item"`
}
type Item struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
	Url         string `xml:"Url"`
}

func SendXMLMessageByText(textMessage *TextMessage) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	keyWord, err := octopus.FindOneKeyWord(textMessage)
	if err == nil {
		switch keyWord.MessageType {
		case "文本消息":
			text, _ := octopus.FindOneTextMessage(keyWord.MessageTitle)
			replyTextMessage := &ReplyTextMessage{
				ToUserName:   textMessage.FromUserName,
				FromUserName: "wecareer",
				CreateTime:   12345678,
				MsgType:      "text",
				Content:      text.MessageContent,
			}
			return replyTextMessage, nil, nil, nil, nil, nil, nil
		case "图文消息":

		}
	} else if textMessage == "人工服务" {

	} else if textMessage == "智能助手" {

	} else {

	}
	if err != nil {
		notKnowReplys := octopus.FindNotKnowReplys()
	}
	fmt.Println("00000000000000")
	text, _ := octopus.FindOneTextMessage("联系我们")
	fmt.Println(text.MessageContent)
	fmt.Println("*********************")
	fmt.Println(textMessage.Content)
	tulingData, _ := TuLingTobot(textMessage.Content, textMessage.FromUserName)
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   textMessage.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      tulingData.Text,
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
}
func SendXMLMessageByPicture(picturesMessage *PicturesMessage) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   picturesMessage.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      "picturesMessage 1",
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
}
func SendXMLMessageByVoice(voiceMessage *VoiceMessage) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   voiceMessage.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      "voiceMessage 1",
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
}
func SendXMLMessageByVideo(videoMessage *VideoMessage) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   videoMessage.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      "videoMessage 1",
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
}
func SendXMLMessageBySmallVideo(smallVideoMessage *SmallVideoMessage) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   smallVideoMessage.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      "smallVideoMessage 1",
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
}
func SendXMLMessageByLocation(locationMessage *LocationMessage) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   locationMessage.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      "locationMessage 1",
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
}
func SendXMLMessageByLinks(linksMessage *LinksMessage) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   linksMessage.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      "linksMessage 1",
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
}
func SendXMLMessageByBaseEvents(baseEventsMessages *BaseEventsMessages) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	//推送事件分析
	switch baseEventsMessages.Event {
	case "subscribe":
		welcomeText, _ := octopus.FindWelcomeMessage()
		textMessage, _ := octopus.FindOneTextMessage(welcomeText[0].MessageTitle)
		replyTextMessage := &ReplyTextMessage{
			ToUserName:   baseEventsMessages.FromUserName,
			FromUserName: "wecareer",
			CreateTime:   123,
			MsgType:      "text",
			Content:      textMessage.MessageContent,
		}
		return replyTextMessage, nil, nil, nil, nil, nil, nil
	case "unsubscribe":

	}
	return nil, nil, nil, nil, nil, nil, nil
}
func SendXMLMessageByScannedQRCodeEvent(scannedQRCodeEventMessages *ScannedQRCodeEventMessages) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   scannedQRCodeEventMessages.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      "scannedQRCodeEventMessages 1",
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
}
func SendXMLMessageBySendLocationEvent(SendLocationEventMessages *SendLocationEventMessages) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   SendLocationEventMessages.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      "SendLocationEventMessages 1",
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
}
func SendXMLMessageByCustomMenuEvent(customMenuEventMessages *CustomMenuEventMessages) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   customMenuEventMessages.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      "customMenuEventMessages 1",
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
}

type SendTuLingData struct {
	Key    string `json:"key"`
	Info   string `json:"info"`
	UserId string `json:"userid"`
}
type AcceptData struct {
	Code int64  `json:"code"`
	Text string `json:"text"`
}

func TuLingTobot(text, openId string) (*AcceptData, error) {
	req := httplib.Get("http://www.tuling123.com/openapi/api?key=4acfe853b09fd0342b28344615ba5f07&info=" + text)
	str, err := req.String()
	if err != nil {
		return nil, err
	}
	var acceptData AcceptData
	fmt.Println(str)
	err = json.Unmarshal([]byte(str), &acceptData)
	if err != nil {
		return nil, err
	}
	return &acceptData, nil
}
