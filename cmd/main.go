package main

import (
	"anchor_push"
	"github.com/scalalang2/golang-fifo/sieve"
)

func main() {
	pusher := &anchor_push.Pusher{
		UrlList: []string{
			"https://zkpt.zj.msa.gov.cn/trafficflow-api/api/v1/out/apply-anchorages/index?portArea=NORTH_PORT&anchorGround=eb956b95dbfc4a10bfb778f2d69d991d",
			"https://zkpt.zj.msa.gov.cn/trafficflow-api/api/v1/out/apply-anchorages/index?portArea=CORE_PORT&anchorGround=daf9884c56a141d891062ba3b1909876",
		},
		DingURL: "https://oapi.dingtalk.com/robot/send?access_token=b2d810ec119282b67b1d98f6fca1a03c96c58ed67d3256c5b5877d149d3aa893",
		Pushed:  sieve.New[string, bool](10000, 0),
	}
	pusher.Start()
}
