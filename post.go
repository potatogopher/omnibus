package main

import (
	"github.com/goadesign/goa"
	"goa-blog/app"
)

// PostController implements the post resource.
type PostController struct {
	*goa.Controller
}

// NewPostController creates a post controller.
func NewPostController(service *goa.Service) *PostController {
	return &PostController{Controller: service.NewController("PostController")}
}

// Create runs the create action.
func (c *PostController) Create(ctx *app.CreatePostContext) error {
	// TBD: implement
	return nil
}

// Delete runs the delete action.
func (c *PostController) Delete(ctx *app.DeletePostContext) error {
	// TBD: implement
	return nil
}

// Show runs the show action.
func (c *PostController) Show(ctx *app.ShowPostContext) error {
	// TBD: implement
	res := &app.User{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *PostController) Update(ctx *app.UpdatePostContext) error {
	// TBD: implement
	return nil
}
