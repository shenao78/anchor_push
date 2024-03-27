package anchor_push

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Pusher struct {
	URL        string
	DingURL    string
	lastPushed string
}

func (p *Pusher) Start() {
	ticker := time.NewTicker(time.Second * 5)
	for ; true; <-ticker.C {
		resp := &Response{}
		if err := Get(p.URL, resp); err != nil {
			fmt.Println(err)
			continue
		}
		if resp.Status != 200 || resp.Code != "10000" {
			fmt.Println("获取锚地预约动态失败！")
			continue
		}
		anchors := resp.Data

		var pushAnchors []*Anchorage
		for _, anchor := range anchors {
			msg := anchor.FormatMsg()
			if p.lastPushed == msgHash(msg) {
				break
			}

			if anchor.CbStatus == "提前离锚" || anchor.CbStatus == "用户取消" {
				pushAnchors = append(pushAnchors, anchor)
			}
		}
		if p.lastPushed != "" {
			for _, anchor := range pushAnchors {
				if err := p.SendDingMsg(anchor.FormatMsg()); err != nil {
					fmt.Println("钉钉发送失败", err)
				}
			}
		}
		if len(pushAnchors) > 0 {
			p.lastPushed = msgHash(pushAnchors[0].FormatMsg())
		}
	}
}

type DingMsg struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

func (p *Pusher) SendDingMsg(content string) error {
	msg := &DingMsg{
		MsgType: "text",
		Text:    Text{Content: content},
	}
	jsonMsg, _ := json.Marshal(msg)
	return Post(p.DingURL, jsonMsg)
}

func msgHash(msg string) string {
	sha := sha256.New()
	sha.Write([]byte(msg))
	return hex.EncodeToString(sha.Sum(nil))
}
