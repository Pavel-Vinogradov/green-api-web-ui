package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string, timeout time.Duration) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetSettings(idInstance, apiTokenInstance string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/waInstance%s/getSettings/%s", c.baseURL, idInstance, apiTokenInstance)
	return c.makeRequest("GET", url, nil)
}

func (c *Client) GetStateInstance(idInstance, apiTokenInstance string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", c.baseURL, idInstance, apiTokenInstance)
	return c.makeRequest("GET", url, nil)
}

func (c *Client) SendMessage(idInstance, apiTokenInstance, chatID, message string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", c.baseURL, idInstance, apiTokenInstance)
	body := map[string]string{
		"chatId":  chatID,
		"message": message,
	}
	return c.makeRequest("POST", url, body)
}

func (c *Client) SendFileByUrl(idInstance, apiTokenInstance, chatID, urlFile, fileName, caption string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", c.baseURL, idInstance, apiTokenInstance)
	body := map[string]interface{}{
		"chatId":   chatID,
		"urlFile":  urlFile,
		"fileName": fileName,
	}
	if caption != "" {
		body["caption"] = caption
	}
	return c.makeRequest("POST", url, body)
}

func (c *Client) makeRequest(method, url string, body interface{}) (map[string]interface{}, error) {
	log.Printf("Making %s request to: %s", method, url)

	var reqBody io.Reader

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return result, nil
}
