//************************************************************************//
// API "rucci.io": Model Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/goa-atlas
// --design=goa-atlas/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"goa-blog/app"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListPost returns an array of view: default.
func (m *PostDB) ListPost(ctx context.Context, userID int) []*app.Post {
	defer goa.MeasureSince([]string{"goa", "db", "post", "listpost"}, time.Now())

	var native []*Post
	var objs []*app.Post
	err := m.Db.Scopes(PostFilterByUser(userID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Post", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.PostToPost())
	}

	return objs
}

// PostToPost returns the Post representation of Post.
func (m *Post) PostToPost() *app.Post {
	post := &app.Post{}
	post.Content = m.Content
	post.ID = m.ID
	post.Title = m.Title

	return post
}

// OnePost returns an array of view: default.
func (m *PostDB) OnePost(ctx context.Context, id int, userID int) (*app.Post, error) {
	defer goa.MeasureSince([]string{"goa", "db", "post", "onepost"}, time.Now())

	var native Post
	err := m.Db.Scopes(PostFilterByUser(userID, &m.Db)).Table(m.TableName()).Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Post", "error", err.Error())
		return nil, err
	}

	view := *native.PostToPost()
	return &view, err
}

// MediaType Retrieval Functions

// ListPostLink returns an array of view: link.
func (m *PostDB) ListPostLink(ctx context.Context, userID int) []*app.PostLink {
	defer goa.MeasureSince([]string{"goa", "db", "post", "listpostlink"}, time.Now())

	var native []*Post
	var objs []*app.PostLink
	err := m.Db.Scopes(PostFilterByUser(userID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Post", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.PostToPostLink())
	}

	return objs
}

// PostToPostLink returns the Post representation of Post.
func (m *Post) PostToPostLink() *app.PostLink {
	post := &app.PostLink{}
	post.ID = m.ID

	return post
}

// OnePostLink returns an array of view: link.
func (m *PostDB) OnePostLink(ctx context.Context, id int, userID int) (*app.PostLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "post", "onepostlink"}, time.Now())

	var native Post
	err := m.Db.Scopes(PostFilterByUser(userID, &m.Db)).Table(m.TableName()).Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Post", "error", err.Error())
		return nil, err
	}

	view := *native.PostToPostLink()
	return &view, err
}
