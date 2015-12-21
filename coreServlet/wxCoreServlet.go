package CoreServlet

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
	"wechat/models/octopus"
	. "wechat/models/octopus/wxMessageService"
)

//检查token
func CheckSignature(signature, timestamp, nonce string) error {
	accounts, err := octopus.FindAllAccount()
	if err != nil {
		return err
	}
	token := accounts[0].Token
	s := []string{token, timestamp, nonce}
	sort.Sort(sort.StringSlice(s))
	s0 := strings.Join(s, "")
	t := sha1.New()
	io.WriteString(t, s0)
	s1 := fmt.Sprintf("%x", t.Sum(nil))
	if signature == s1 {
		_, err := octopus.GetOrUpdateAccessToken()
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

//接收message
func AcceptMessage(data []byte) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	textMessage, picturesMessage, voiceMessage, videoMessage, smallVideoMessage, locationMessage, linksMessage, baseEventsMessages, scannedQRCodeEventMessages, sendLocationEventMessages, customMenuEventMessages, err := ParsingMessageXml(data)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := DealWithMessage(textMessage, picturesMessage, voiceMessage, videoMessage, smallVideoMessage, locationMessage, linksMessage, baseEventsMessages, scannedQRCodeEventMessages, sendLocationEventMessages, customMenuEventMessages)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
}

//处理message
func DealWithMessage(textMessage *TextMessage, picturesMessage *PicturesMessage, voiceMessage *VoiceMessage, videoMessage *VideoMessage, smallVideoMessage *SmallVideoMessage, locationMessage *LocationMessage, linksMessage *LinksMessage, baseEventsMessages *BaseEventsMessages, scannedQRCodeEventMessages *ScannedQRCodeEventMessages, sendLocationEventMessages *SendLocationEventMessages, customMenuEventMessages *CustomMenuEventMessages) (*ReplyTextMessage, *ReplyPictureMessage, *ReplyVoiceMessage, *ReplyVideoMessage, *ReplyMusicMessage, *ReplyNewsTextMessage, error) {
	var err error
	switch {
	case textMessage != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageByText(textMessage)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	case picturesMessage != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageByPicture(picturesMessage)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	case voiceMessage != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageByVoice(voiceMessage)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	case videoMessage != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageByVideo(videoMessage)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	case smallVideoMessage != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageBySmallVideo(smallVideoMessage)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	case locationMessage != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageByLocation(locationMessage)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	case linksMessage != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageByLinks(linksMessage)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	case baseEventsMessages != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageByBaseEvents(baseEventsMessages)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	case scannedQRCodeEventMessages != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageByScannedQRCodeEvent(scannedQRCodeEventMessages)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	case sendLocationEventMessages != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageBySendLocationEvent(sendLocationEventMessages)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	case customMenuEventMessages != nil:
		replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, err := SendXMLMessageByCustomMenuEvent(customMenuEventMessages)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, err
		}
		return replyTextMessage, replyPictureMessage, replyVoiceMessage, replyVideoMessage, replyMusicMessage, replyNewsTextMessage, nil
	}
	return nil, nil, nil, nil, nil, nil, err
}
