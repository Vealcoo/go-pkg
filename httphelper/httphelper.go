package httphelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

func SendPostWithJSON(apiHost string, res interface{}, data interface{}, timeOut time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: timeOut}
	req, err := http.NewRequest("POST", apiHost, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	apiRes, err := client.Do(req)
	if err != nil {
		return err
	}
	defer apiRes.Body.Close()

	resBody, err := io.ReadAll(apiRes.Body)
	if err != nil {
		return err
	}

	if apiRes.StatusCode != http.StatusOK {
		return fmt.Errorf("req: %s, res: %s", string(jsonData), string(resBody))
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return err
	}

	return nil
}

func SendPostWithJSONAndToken(apiHost string, res interface{}, data interface{}, token string, timeOut time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: timeOut}
	req, err := http.NewRequest("POST", apiHost, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer "+token)

	apiRes, err := client.Do(req)
	if err != nil {
		return err
	}
	defer apiRes.Body.Close()

	resBody, err := io.ReadAll(apiRes.Body)
	if err != nil {
		return err
	}

	if apiRes.StatusCode != http.StatusOK {
		return fmt.Errorf("req: %s, res: %s", string(jsonData), string(resBody))
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return err
	}

	return nil
}

func SendPostWithUrl(apiHost string, res interface{}, data url.Values, timeOut time.Duration) error {
	// New multiport writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range data {
		writer.WriteField(k, v[0])
	}

	err := writer.Close()
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: timeOut}
	req, err := http.NewRequest("POST", apiHost, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	apiRes, err := client.Do(req)
	if err != nil {
		return err
	}
	defer apiRes.Body.Close()

	resBody, err := io.ReadAll(apiRes.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return err
	}

	return nil
}
