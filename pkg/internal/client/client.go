package client

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/ocodo/opengist-cli/internal/config"
)

type FileContent struct {
	Content string `json:"content"`
}

type AddPayload struct {
	Files       map[string]FileContent `json:"files"`
	Title       string                 `json:"title,omitempty"`
	Description string                 `json:"description,omitempty"`
	Visibility  string                 `json:"visibility,omitempty"`
}

type EditPayload struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Visibility  string `json:"visibility,omitempty"`
}

type Gist struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

type Client struct {
	http *resty.Client
}

func NewClient() (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	baseURL := strings.TrimRight(cfg.URL, "/")

	r := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Authorization", "Bearer "+cfg.Token).
		SetHeader("Accept", "application/json").
		SetTimeout(30 * time.Second)

	return &Client{http: r}, nil
}

func (c *Client) Add(payload *AddPayload) (*resty.Response, error) {
	return c.http.R().
		SetBody(payload).
		Post("/api/gists")
}

func (c *Client) Edit(gistID string, payload *EditPayload) (*resty.Response, error) {
	return c.http.R().
		SetBody(payload).
		Patch("/api/gists/" + gistID)
}

func (c *Client) Delete(gistID string) (*resty.Response, error) {
	return c.http.R().
		Delete("/api/gists/" + gistID)
}

func (c *Client) List(username string, page, perPage int) ([]Gist, *resty.Response, error) {
	endpoint := "/api/gists"
	if username != "" {
		endpoint = fmt.Sprintf("/api/users/%s/gists", username)
	}

	var gists []Gist
	resp, err := c.http.R().
		SetQueryParam("page", fmt.Sprintf("%d", page)).
		SetQueryParam("per_page", fmt.Sprintf("%d", perPage)).
		SetResult(&gists).
		Get(endpoint)

	if err != nil {
		return nil, nil, err
	}

	return gists, resp, nil
}

func HandleResponse(resp *resty.Response) {
	if resp.IsSuccess() {
		if len(resp.Body()) == 0 {
			return
		}

		var data map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &data); err == nil {
			if htmlURL, ok := data["html_url"].(string); ok && htmlURL != "" {
				fmt.Println(htmlURL)
				return
			}
			if url, ok := data["url"].(string); ok && url != "" {
				fmt.Println(url)
				return
			}
		}

		fmt.Println(resp.String())
		return
	}

	fmt.Fprintf(os.Stderr, "opengist-cli: OpenGist returned HTTP %d\n", resp.StatusCode())
	if len(resp.Body()) > 0 {
		fmt.Fprintln(os.Stderr, resp.String())
	}
}
