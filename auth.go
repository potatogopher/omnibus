package main

import (
	"github.com/goadesign/goa"
	"goa-blog/app"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
}

// Token runs the token action.
func (c *AuthController) Token(ctx *app.TokenAuthContext) error {
	// TBD: implement
	return nil
}
