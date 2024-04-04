package openai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/namth2302/go-rest"
)

type Client struct {
	config ClientConfig
}

func NewClient(openAiKey string) *Client {
	config := DefaultConfig(openAiKey)
	return NewClientWithConfig(config)
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config: config,
	}
}

func NewOrgClient(openAiKey string, orgId string) *Client {
	config := DefaultConfig(openAiKey)
	config.OrgID = orgId
	return NewClientWithConfig(config)
}

type requestOptions struct {
	Body        any
	QueryParams any
}

type requestOption func(**requestOptions)

type Response interface {
}

func (c *Client) setHeaders() map[string]string {
	headers := make(map[string]string)

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	if c.config.OpenAiKey != "" {
		// OpenAI or Azure AD authentication
		headers["Authorization"] = fmt.Sprintf("Bearer %s", c.config.OpenAiKey)
	}
	if c.config.OrgID != "" {
		headers["OpenAI-Organization"] = c.config.OrgID
	}
	return headers
}

func isFailureStatusCode(resp *rest.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

// fullURL returns full URL for request.
// args[0] is model name, if API type is Azure, model name is required to get deployment name.
func (c *Client) fullURL(suffix string) string {
	// /openai/deployments/{model}/chat/completions?api-version={api_version}
	return fmt.Sprintf("%s%s", c.config.BaseURL, suffix)
}

func (c *Client) newRequest(method rest.Method, url string, setters ...requestOption) (rest.Request, error) {
	req := rest.Request{
		Method:  method,
		BaseURL: url,
		Headers: c.setHeaders(),
	}
	for _, setter := range setters {
		fmt.Println(setter)
	}
	return req, nil
}

func (c *Client) sendRequest(req rest.Request, v Response) error {
	resp, err := rest.Send(req)
	if err != nil {
		err = fmt.Errorf("failed to send request: %v", err)
		return err
	}

	if isFailureStatusCode(resp) {
		err = fmt.Errorf("failed to send request openai, status code: %d", resp.StatusCode)
		return err
	}

	return json.NewDecoder(strings.NewReader(resp.Body)).Decode(v)
}
