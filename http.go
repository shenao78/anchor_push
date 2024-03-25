package anchor_push

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func Get(url string, result interface{}) error {
	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {

	}
	return json.Unmarshal(body, result)
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
