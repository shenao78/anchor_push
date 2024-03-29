package main

import (
	"anchor_push"
	"github.com/scalalang2/golang-fifo/sieve"
)

func main() {
	pusher := &anchor_push.Pusher{
		URL:     "https://zkpt.zj.msa.gov.cn/trafficflow-api/api/v1/out/apply-anchorages/index?portArea=NORTH_PORT&anchorGround=eb956b95dbfc4a10bfb778f2d69d991d",
		DingURL: "https://oapi.dingtalk.com/robot/send?access_token=b2d810ec119282b67b1d98f6fca1a03c96c58ed67d3256c5b5877d149d3aa893",
		Pushed:  sieve.New[string, bool](10000, 0),
	}
	pusher.Start()
}
