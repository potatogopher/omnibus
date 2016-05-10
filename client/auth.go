package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
)

// TokenAuthPayload is the auth token action payload.
type TokenAuthPayload struct {
	Email    string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
}

// TokenAuthPath computes a request path to the token action of auth.
func TokenAuthPath() string {
	return fmt.Sprintf("/auth/token")
}

// Obtain an access token
func (c *Client) TokenAuth(ctx context.Context, path string, payload *TokenAuthPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}
