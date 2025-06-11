package salesforceclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) InsertDataRows(rows []interface{}, brand, extensionKey string) (*InsertDataRowsResponse, error) {
	url := fmt.Sprintf("%s/v1/InsertRows/%s/%s", c.Config.Url, brand, extensionKey)
	payload := map[string]interface{}{"items": rows}

	resp, err := c.post(url, payload, &InsertDataRowsResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*InsertDataRowsResponse), nil
}

func (c *Client) InsertDataRow(row interface{}, brand, extensionKey string) (*InsertDataRowsResponse, error) {
	url := fmt.Sprintf("%s/v1/InsertRows/%s/%s", c.Config.Url, brand, extensionKey)
	resp, err := c.post(url, row, &InsertDataRowsResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*InsertDataRowsResponse), nil
}

func (c *Client) CheckInsert(requestId, brand string) (*CheckInsertRowsStatus, error) {
	url := fmt.Sprintf("%s/v1/InsertRows/%s/check/%s", c.Config.Url, brand, requestId)
	resp, err := c.get(url, &CheckInsertRowsStatus{})
	if err != nil {
		return nil, err
	}
	return resp.(*CheckInsertRowsStatus), nil
}

func (c *Client) SendEmail(payload SendEmailRequest, brand string) (*SendEmailResponse, error) {
	url := fmt.Sprintf("%s/v1/%s/fireEvent", c.Config.Url, brand)
	resp, err := c.post(url, payload, &SendEmailResponse{})
	if err != nil {
		return nil, err
	}
	return resp.(*SendEmailResponse), nil
}

// Helper methods

func (c *Client) post(url string, payload interface{}, out interface{}) (interface{}, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := c.buildRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	return c.doRequest(req, out)
}

func (c *Client) get(url string, out interface{}) (interface{}, error) {
	req, err := c.buildRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.doRequest(req, out)
}

func (c *Client) buildRequest(method, url string, body []byte) (*http.Request, error) {
	var buf io.Reader
	if body != nil {
		buf = bytes.NewBuffer(body)
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apiKey", c.Config.ApiKey)
	return req, nil
}

func (c *Client) doRequest(req *http.Request, out interface{}) (interface{}, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error: %s", string(respBody))
	}

	if err := json.Unmarshal(respBody, out); err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %w", err)
	}
	return out, nil
}
