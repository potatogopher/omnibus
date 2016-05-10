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

// CreatePostPayload is the post create action payload.
type CreatePostPayload struct {
	Content string `json:"content" xml:"content"`
	Title   string `json:"title" xml:"title"`
}

// CreatePostPath computes a request path to the create action of post.
func CreatePostPath() string {
	return fmt.Sprintf("/posts")
}

// Create new post
func (c *Client) CreatePost(ctx context.Context, path string, payload *CreatePostPayload) (*http.Response, error) {
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

// DeletePostPath computes a request path to the delete action of post.
func DeletePostPath(postID int) string {
	return fmt.Sprintf("/posts/%v", postID)
}

// DeletePost makes a request to the delete action endpoint of the post resource
func (c *Client) DeletePost(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}

// ShowPostPath computes a request path to the show action of post.
func ShowPostPath(postID int) string {
	return fmt.Sprintf("/posts/%v", postID)
}

// Retrieve post with given id.
func (c *Client) ShowPost(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}

// UpdatePostPayload is the post update action payload.
type UpdatePostPayload struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

// UpdatePostPath computes a request path to the update action of post.
func UpdatePostPath(postID int) string {
	return fmt.Sprintf("/posts/%v", postID)
}

// Change post attributes
func (c *Client) UpdatePost(ctx context.Context, path string, payload *UpdatePostPayload) (*http.Response, error) {
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
	req, err := http.NewRequest("PUT", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}
