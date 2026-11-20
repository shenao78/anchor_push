package anchor_push

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const maxAnchoragePages = 100

func Get(rawURL string, result interface{}) error {
	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Get(rawURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {

	}
	return json.Unmarshal(body, result)
}

// GetAllAnchorages 按 page 顺序拉取全部公示数据，直到 hasMore 为 false。
func GetAllAnchorages(rawURL string) ([]*Anchorage, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	q := u.Query()

	var all []*Anchorage
	for page := 1; page <= maxAnchoragePages; page++ {
		q.Set("page", strconv.Itoa(page))
		u.RawQuery = q.Encode()

		resp := &Response{}
		if err := Get(u.String(), resp); err != nil {
			return nil, err
		}
		if resp.Status != 200 || resp.Code != "10000" {
			return nil, fmt.Errorf("获取锚地预约动态失败: status=%d code=%s", resp.Status, resp.Code)
		}

		all = append(all, resp.Data.List...)
		if !resp.Data.HasMore {
			return all, nil
		}
	}
	return nil, fmt.Errorf("分页超过上限 %d 仍 hasMore=true", maxAnchoragePages)
}

func Post(url string, payload []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 20 * time.Second}
	_, err = client.Do(req)
	return err
}
