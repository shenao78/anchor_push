package anchor_push

import (
	"fmt"
	"time"
)

type Response struct {
	Status  uint64       `json:"status"`
	Code    string       `json:"code"`
	Success bool         `json:"success"`
	Data    []*Anchorage `json:"data"`
}

type Anchorage struct {
	ElementId             string `json:"elementId"`
	ShipName              string `json:"shipName"`
	CbStatus              string `json:"cbStatus"`
	PredictAnchorGround   string `json:"predictAnchorGround"`
	PredictAnchorPosition string `json:"predictAnchorPosition"`
	ArrangeAnchorTime     string `json:"arrangeAnchorTime"`
	ArrangeMoveAnchorTime string `json:"arrangeMoveAnchorTime"`
	PublishTimeStr        string `json:"publishTime"`

	PushTime time.Time `json:"-"`
}

func (a *Anchorage) FormatMsg() string {
	template := `[锚地预约状态变更通知]
船舶名称：%s
锚地名称：%s
锚位名称：%s
抛锚时间：%s
离锚时间：%s
状    态：%s
发布时间：%s
`
	return fmt.Sprintf(
		template,
		a.ShipName,
		a.PredictAnchorGround,
		a.PredictAnchorPosition,
		a.ArrangeAnchorTime,
		a.ArrangeMoveAnchorTime,
		a.CbStatus,
		a.PublishTimeStr,
	)
}

func (a *Anchorage) fillPushTime() {
	pushTime, _ := parseTime(a.PublishTimeStr)
	a.PushTime = pushTime
}

func parseTime(t string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	year := time.Now().Year()
	return time.ParseInLocation("2006-01-02 15:04", fmt.Sprintf("%d-%s", year, t), loc)
}
