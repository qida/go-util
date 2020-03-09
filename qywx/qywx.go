package qywx

import (
	"time"

	"github.com/astaxie/beego/httplib"
)

type MsgRestult struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}
type MsgMarkdown struct {
	Touser   string `json:"touser"`
	Toparty  string `json:"toparty"`
	Totag    string `json:"totag"`
	Msgtype  string `json:"msgtype"`
	Agentid  int    `json:"agentid"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
	EnableDuplicateCheck int `json:"enable_duplicate_check"`
}

type MsgAccessToken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
type QYWX struct {
	AccessToken string
	TimeExpires time.Time

	CorpId     string
	CorpSecret string
}

func NewQywx(corpid, corpsecret string) (qy *QYWX, err error) {
	req := httplib.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken")
	req.Param("corpid", corpid)
	req.Param("corpsecret", corpsecret)
	var msgAccessToken MsgAccessToken
	err = req.ToJSON(&msgAccessToken)
	if err != nil {
		return
	}
	qy = &QYWX{
		AccessToken: msgAccessToken.AccessToken,
		TimeExpires: time.Now().Add(7200),
		CorpId:      corpid,
		CorpSecret:  corpsecret,
	}
	return
}
func (q *QYWX) GetAccessToken() string {
	if time.Now().Sub(q.TimeExpires).Seconds() > 0 {
		q, _ = NewQywx(q.CorpId, q.CorpSecret)
	}
	return q.AccessToken
}
