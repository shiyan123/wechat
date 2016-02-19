package octopus

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
	"wechat/models/octopus/wxBaseParameter"
)

type NewsMessagesList struct {
	Id                int64
	NewsMessagesTitle string
	NewsMessagesId    string
	LastUpdateTime    time.Time
}
type SendNewsMessageData struct {
	Type   string `json:"type"`
	Offset int64  `json:"offset"`
	Count  int    `json:"count"`
}
type AcceptDataNewsMessageList struct {
	TotalCount int    `json:"total_count"`
	ItemCount  int    `json:"item_count"`
	Items      []Item `json:"item"`
}
type Item struct {
	MediaId    string      `json:"media_id"`
	Content    NewsContent `json:"content"`
	UpdateTime int64       `json:"update_time"`
}
type NewsContent struct {
	NerwsItems []NerwsItem `json:"news_item"`
}
type NerwsItem struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Contet           string `json:"contet"`
	Url              string `json:"url"`
	ContentSourceUrl string `json:"content_source_url"`
}

func InsertNewsMessageList(newsMessagesLists []NewsMessagesList) error {
	o := orm.NewOrm()
	qs := o.QueryTable("news_messages_list")
	i, _ := qs.PrepareInsert()
	for _, newsMessagesList := range newsMessagesLists {
		id, err := i.Insert(&newsMessagesList)
		if err != nil {
			return err
		}
		fmt.Println(id)
	}
	i.Close()
	return nil
}
func FindAllNewsMessageList() ([]NewsMessagesList, error, int64) {
	o := orm.NewOrm()
	var newsMessagesLists []NewsMessagesList
	num, err := o.Raw("select * from news_messages_list").QueryRows(&newsMessagesLists)
	if err == orm.ErrNoRows && num == 0 {
		err = errors.New("查询不到")
		return nil, err, 0
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
		return nil, err, 0
	} else {
		return newsMessagesLists, nil, num
	}
}
func GetNewsMessageList(offset int64) error {
	url := wxBaseParameter.GetWxChatHttpUrl().MaterialListUrl
	accessToken, _ := GetOrUpdateAccessToken()
	newUrl := strings.Replace(url, "ACCESS_TOKEN", accessToken.AccessToken, -1)
	req := httplib.Post(newUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	data := &SendNewsMessageData{
		Type:   "news",
		Offset: offset,
		Count:  20,
	}
	a, _ := json.Marshal(data)
	fmt.Println(string(a))
	b, _ := req.JsonBody(data)
	fmt.Println(b)
	str, err := req.String()
	if err != nil {
		return err
	}
	var acceptDataNewsMessageList AcceptDataNewsMessageList
	err = json.Unmarshal([]byte(str), &acceptDataNewsMessageList)
	if err != nil {
		return err
	}
	var newsMessagesLists []NewsMessagesList
	for i := 0; i < acceptDataNewsMessageList.ItemCount; {
		t := time.Unix(acceptDataNewsMessageList.Items[i].UpdateTime, 0)
		newsMessagesList := NewsMessagesList{
			NewsMessagesTitle: acceptDataNewsMessageList.Items[i].Content.NerwsItems[0].Title,
			NewsMessagesId:    acceptDataNewsMessageList.Items[i].MediaId,
			LastUpdateTime:    t,
		}
		newsMessagesLists = append(newsMessagesLists, newsMessagesList)
		i++
	}
	err = InsertNewsMessageList(newsMessagesLists)
	if err != nil {
		return err
	}
	err = GetNewsMessageInfo(newsMessagesLists)
	if err != nil {
		return err
	}
	return nil
}
func GetNewsMessageInfo(newsMessagesLists []NewsMessagesList) error {
	for _, newsMessagesList := range newsMessagesLists {
		fmt.Println(newsMessagesList.NewsMessagesId)
	}
	return nil
}
