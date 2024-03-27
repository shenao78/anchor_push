package anchor_push

import (
	"fmt"
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
