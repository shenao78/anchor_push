package anchor_push

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

type Pusher struct {
	URL            string
	DingURL        string
	LastPushedTime time.Time
}

func (p *Pusher) Start() {
	ticker := time.NewTicker(time.Second)
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
		for _, anchor := range anchors {
			anchor.fillPushTime()
		}
		sort.Slice(anchors, func(i, j int) bool {
			return anchors[i].PushTime.After(anchors[j].PushTime)
		})
		var pushAnchors []*Anchorage
		for _, anchor := range anchors {
			if !anchor.PushTime.After(p.LastPushedTime) {
				break
			}
			if anchor.CbStatus == "提前离锚" || anchor.CbStatus == "用户取消" {
				pushAnchors = append(pushAnchors, anchor)
			}
		}
		for _, anchor := range pushAnchors {
			if err := p.SendDingMsg(anchor.FormatMsg()); err != nil {
				fmt.Println("钉钉发送失败", err)
			}
		}
		if len(pushAnchors) > 0 {
			p.LastPushedTime = pushAnchors[0].PushTime
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
