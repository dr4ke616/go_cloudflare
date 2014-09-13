package cloudflare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	Token       string
	Email       string
	URL         string
	HttpHandler *http.Client
}

func NewClient(email string, token string) (*Client, error) {

	client := Client{
		Token:       token,
		Email:       email,
		URL:         "https://www.cloudflare.com/api_json.html",
		HttpHandler: http.DefaultClient,
	}

	return &client, nil
}

func (c *Client) NewRequest(params map[string]string, method string, action string) (*http.Request, error) {
	data := url.Values{}
	u, err := url.Parse(c.URL)

	if err != nil {
		return nil, fmt.Errorf("Error parsing base URL: %s", err)
	}

	data.Add("email", c.Email)
	data.Add("tkn", c.Token)
	data.Add("a", action)

	for k, v := range params {
		data.Add(k, v)
	}

	u.RawQuery = data.Encode()

	req, err := http.NewRequest(method, u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("Error creating request: %s", err)
	}

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &out); err != nil {
		return err
	}

	return nil
}

func checkResponse(resp *http.Response, err error) (*http.Response, error) {

	if err != nil {
		return resp, err
	}

	switch i := resp.StatusCode; {
	case i == 200:
		return resp, nil
	default:
		return nil, fmt.Errorf("API Error: %s", resp.Status)
	}
}
