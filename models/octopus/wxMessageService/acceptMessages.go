package wxMessageService

import (
	"encoding/xml"
	"fmt"
)

type BaseMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	Ticket       string `xml:"Ticket"`
}
type TextMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        int64  `xml:"MsgId"`
}
type PicturesMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	PicUrl       string `xml:"PicUrl"`
	MediaId      string `xml:"MediaId"`
	MsgId        int64  `xml:"MsgId"`
}
type VoiceMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	MediaId      string `xml:"MediaId"`
	Format       string `xml:"Format"`
	MsgId        int64  `xml:"MsgId"`
}
type VideoMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	MediaId      string `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
	MsgId        int64  `xml:"MsgId"`
}
type SmallVideoMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	MediaId      string `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
	MsgId        int64  `xml:"MsgId"`
}
type LocationMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Location_X   string `xml:"Location_X"`
	Location_Y   string `xml:"Location_Y"`
	Scale        string `xml:"Scale"`
	Label        string `xml:"Label"`
	MsgId        int64  `xml:"MsgId"`
}
type LinksMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Title        string `xml:"Title"`
	Description  string `xml:"Description"`
	Url          int64  `xml:"Url"`
	MsgId        string `xml:"MsgId"`
}
type BaseEventsMessages struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
}
type ScannedQRCodeEventMessages struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	Ticket       string `xml:"Ticket"`
}
type SendLocationEventMessages struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	Latitude     string `xml:"Latitude"`
	Longitude    string `xml:"Longitude"`
	Precision    string `xml:"Precision"`
}
type CustomMenuEventMessages struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
}

func ParsingTextXml(data []byte) (*TextMessage, error) {
	var textMessage TextMessage
	err := xml.Unmarshal(data, &textMessage)
	if err != nil {
		return nil, err
	}
	fmt.Println(textMessage.MsgType)
	return &textMessage, nil
}
func ParsingPicturesXml(data []byte) (*PicturesMessage, error) {
	var picturesMessage PicturesMessage
	err := xml.Unmarshal(data, &picturesMessage)
	if err != nil {
		return nil, err
	}
	fmt.Println(picturesMessage.MsgType)
	return &picturesMessage, nil
}
func ParsingVoiceXml(data []byte) (*VoiceMessage, error) {
	var voiceMessage VoiceMessage
	err := xml.Unmarshal(data, &voiceMessage)
	if err != nil {
		return nil, err
	}
	fmt.Println(voiceMessage.MsgType)
	return &voiceMessage, nil
}
func ParsingVideoXml(data []byte) (*VideoMessage, error) {
	var videoMessage VideoMessage
	err := xml.Unmarshal(data, &videoMessage)
	if err != nil {
		return nil, err
	}
	fmt.Println(videoMessage.MsgType)
	return &videoMessage, nil
}
func ParsingSmallVideoXml(data []byte) (*SmallVideoMessage, error) {
	var smallVideoMessage SmallVideoMessage
	err := xml.Unmarshal(data, &smallVideoMessage)
	if err != nil {
		return nil, err
	}
	fmt.Println(smallVideoMessage.MsgType)
	return &smallVideoMessage, nil
}
func ParsingLocationXml(data []byte) (*LocationMessage, error) {
	var locationMessage LocationMessage
	err := xml.Unmarshal(data, &locationMessage)
	if err != nil {
		return nil, err
	}
	fmt.Println(locationMessage.MsgType)
	return &locationMessage, nil
}
func ParsingLinksXml(data []byte) (*LinksMessage, error) {
	var linksMessage LinksMessage
	err := xml.Unmarshal(data, &linksMessage)
	if err != nil {
		return nil, err
	}
	fmt.Println(linksMessage.MsgType)
	return &linksMessage, nil
}
func ParsingBaseEventsXml(data []byte) (*BaseEventsMessages, error) {
	var baseEventsMessages BaseEventsMessages
	err := xml.Unmarshal(data, &baseEventsMessages)
	if err != nil {
		return nil, err
	}
	fmt.Println(baseEventsMessages.MsgType)
	return &baseEventsMessages, nil
}
func ParsingScannedQRCodeEventXml(data []byte) (*ScannedQRCodeEventMessages, error) {
	var scannedQRCodeEventMessages ScannedQRCodeEventMessages
	err := xml.Unmarshal(data, &scannedQRCodeEventMessages)
	if err != nil {
		return nil, err
	}
	fmt.Println(scannedQRCodeEventMessages.MsgType)
	return &scannedQRCodeEventMessages, nil
}
func ParsingSendLocationEventXml(data []byte) (*SendLocationEventMessages, error) {
	var sendLocationEventMessages SendLocationEventMessages
	err := xml.Unmarshal(data, &sendLocationEventMessages)
	if err != nil {
		return nil, err
	}
	fmt.Println(sendLocationEventMessages.MsgType)
	return &sendLocationEventMessages, nil
}
func ParsingCustomMenuEvent(data []byte) (*CustomMenuEventMessages, error) {
	var customMenuEventMessages CustomMenuEventMessages
	err := xml.Unmarshal(data, &customMenuEventMessages)
	if err != nil {
		return nil, err
	}
	fmt.Println(customMenuEventMessages.MsgType)
	return &customMenuEventMessages, nil
}
func ParsingMessageXml(data []byte) (
	*TextMessage,
	*PicturesMessage,
	*VoiceMessage,
	*VideoMessage,
	*SmallVideoMessage,
	*LocationMessage,
	*LinksMessage,
	*BaseEventsMessages,
	*ScannedQRCodeEventMessages,
	*SendLocationEventMessages,
	*CustomMenuEventMessages,
	error) {
	var baseMessage BaseMessage
	err := xml.Unmarshal(data, &baseMessage)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	switch baseMessage.MsgType {
	case "text":
		textMessage, _ := ParsingTextXml(data)
		return textMessage, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil
	case "image":
		picturesMessage, _ := ParsingPicturesXml(data)
		return nil, picturesMessage, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil
	case "voice":
		voiceMessage, _ := ParsingVoiceXml(data)
		return nil, nil, voiceMessage, nil, nil, nil, nil, nil, nil, nil, nil, nil
	case "video":
		videoMessage, _ := ParsingVideoXml(data)
		return nil, nil, nil, videoMessage, nil, nil, nil, nil, nil, nil, nil, nil
	case "shortvideo":
		smallVideoMessage, _ := ParsingSmallVideoXml(data)
		return nil, nil, nil, nil, smallVideoMessage, nil, nil, nil, nil, nil, nil, nil
	case "location":
		locationMessage, _ := ParsingLocationXml(data)
		return nil, nil, nil, nil, nil, locationMessage, nil, nil, nil, nil, nil, nil
	case "link":
		linksMessage, _ := ParsingLinksXml(data)
		return nil, nil, nil, nil, nil, nil, linksMessage, nil, nil, nil, nil, nil
	case "event":
		if baseMessage.Event == "subscribe" || baseMessage.Event == "unsubscribe" && baseMessage.Ticket == "" {
			baseEventsMessages, _ := ParsingBaseEventsXml(data)
			return nil, nil, nil, nil, nil, nil, nil, baseEventsMessages, nil, nil, nil, nil
		}
		if baseMessage.Ticket != "" {
			scannedQRCodeEventMessages, _ := ParsingScannedQRCodeEventXml(data)
			return nil, nil, nil, nil, nil, nil, nil, nil, scannedQRCodeEventMessages, nil, nil, nil
		}
		if baseMessage.Event == "LOCATION" {
			sendLocationEventMessages, _ := ParsingSendLocationEventXml(data)
			return nil, nil, nil, nil, nil, nil, nil, nil, nil, sendLocationEventMessages, nil, nil
		}
		if baseMessage.Event == "CLICK" {
			customMenuEventMessages, _ := ParsingCustomMenuEvent(data)
			return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, customMenuEventMessages, nil
		}
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
}
