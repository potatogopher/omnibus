package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"

	"goa-blog/app"
	"goa-blog/models"
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
	p := models.Post{}

	p.Title = ctx.Payload.Title
	p.Content = ctx.Payload.Content

	post, err := pdb.Add(ctx, &p)
	if err != nil {
		return ErrDatabaseError(err)
	}

	ctx.ResponseData.Header().Set("Location", app.PostHref(post.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *PostController) Delete(ctx *app.DeletePostContext) error {
	pdb.Delete(ctx.Context, ctx.PostID)
	return ctx.NoContent()
}

// Show runs the show action.
func (c *PostController) Show(ctx *app.ShowPostContext) error {
	post, err := pdb.OnePost(ctx.Context, ctx.PostID, 0)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}

	post.Href = app.PostHref(post.ID)
	return ctx.OK(post)
}

// Update runs the update action.
func (c *PostController) Update(ctx *app.UpdatePostContext) error {
	post, err := pdb.Get(ctx.Context, ctx.PostID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}

	post.Title = *ctx.Payload.Title
	post.Content = *ctx.Payload.Content

	err = pdb.Update(ctx, &post)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return nil
}
