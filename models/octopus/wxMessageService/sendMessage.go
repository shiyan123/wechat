package wxMessageService

import (
	"encoding/xml"
	"fmt"
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
	Content      string   `xml:"Content"`
}

func SendXMLMessageByText(textMessage *TextMessage) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	fmt.Println("00000000000000")
	text, _ := octopus.FindOneTextMessage("联系我们")
	fmt.Println(text.MessageContent)
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   textMessage.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      text.MessageContent,
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
	replyTextMessage := &ReplyTextMessage{
		ToUserName:   baseEventsMessages.FromUserName,
		FromUserName: "wecareer",
		CreateTime:   123,
		MsgType:      "text",
		Content:      "baseEventsMessages 1",
	}
	return replyTextMessage, nil, nil, nil, nil, nil, nil
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
