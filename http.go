package fuiousdk

import (
	"net/http"
	"bytes"
	"encoding/json"
	"time"
)

func HttpSend(url string, body interface{}, v interface{}) error {
	bodyData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := http.Client{
		Timeout: time.Second * 3, // set 3 second timeout
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(v)
}
