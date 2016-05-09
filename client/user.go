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

// CreateUserPayload is the user create action payload.
type CreateUserPayload struct {
	// Flag for if the user is disabled or not
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// Email of user
	Email string `json:"email" xml:"email"`
	// Given name of user
	GivenName *string `json:"givenName,omitempty" xml:"givenName,omitempty"`
	Password  string  `json:"password" xml:"password"`
	// Surname of user
	Surname *string `json:"surname,omitempty" xml:"surname,omitempty"`
}

// CreateUserPath computes a request path to the create action of user.
func CreateUserPath() string {
	return fmt.Sprintf("/users")
}

// Create new user
func (c *Client) CreateUser(ctx context.Context, path string, payload *CreateUserPayload) (*http.Response, error) {
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

// DeleteUserPath computes a request path to the delete action of user.
func DeleteUserPath(userID int) string {
	return fmt.Sprintf("/users/%v", userID)
}

// DeleteUser makes a request to the delete action endpoint of the user resource
func (c *Client) DeleteUser(ctx context.Context, path string) (*http.Response, error) {
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

// ShowUserPath computes a request path to the show action of user.
func ShowUserPath(userID int) string {
	return fmt.Sprintf("/users/%v", userID)
}

// Retrieve user with given id. IDs 1 and 2 pre-exist in the system.
func (c *Client) ShowUser(ctx context.Context, path string) (*http.Response, error) {
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

// UpdateUserPayload is the user update action payload.
type UpdateUserPayload struct {
	// Flag for if the user is disabled or not
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// Email of user
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Given name of user
	GivenName   *string `json:"givenName,omitempty" xml:"givenName,omitempty"`
	NewPassword *string `json:"newPassword,omitempty" xml:"newPassword,omitempty"`
	OldPassword *string `json:"oldPassword,omitempty" xml:"oldPassword,omitempty"`
	// Surname of user
	Surname *string `json:"surname,omitempty" xml:"surname,omitempty"`
}

// UpdateUserPath computes a request path to the update action of user.
func UpdateUserPath(userID int) string {
	return fmt.Sprintf("/users/%v", userID)
}

// Change user attributes
func (c *Client) UpdateUser(ctx context.Context, path string, payload *UpdateUserPayload) (*http.Response, error) {
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
