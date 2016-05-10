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
	"goa-atlas/app"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListRucciPost returns an array of view: default.
func (m *PostDB) ListRucciPost(ctx context.Context, userID int) []*app.RucciPost {
	defer goa.MeasureSince([]string{"goa", "db", "rucciPost", "listrucciPost"}, time.Now())

	var native []*Post
	var objs []*app.RucciPost
	err := m.Db.Scopes(PostFilterByUser(userID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Post", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.PostToRucciPost())
	}

	return objs
}

// PostToRucciPost returns the RucciPost representation of Post.
func (m *Post) PostToRucciPost() *app.RucciPost {
	post := &app.RucciPost{}
	post.Content = m.Content
	post.ID = m.ID
	post.Title = m.Title

	return post
}

// OneRucciPost returns an array of view: default.
func (m *PostDB) OneRucciPost(ctx context.Context, id int, userID int) (*app.RucciPost, error) {
	defer goa.MeasureSince([]string{"goa", "db", "rucciPost", "onerucciPost"}, time.Now())

	var native Post
	err := m.Db.Scopes(PostFilterByUser(userID, &m.Db)).Table(m.TableName()).Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Post", "error", err.Error())
		return nil, err
	}

	view := *native.PostToRucciPost()
	return &view, err
}

// MediaType Retrieval Functions

// ListRucciPostLink returns an array of view: link.
func (m *PostDB) ListRucciPostLink(ctx context.Context, userID int) []*app.RucciPostLink {
	defer goa.MeasureSince([]string{"goa", "db", "rucciPost", "listrucciPostlink"}, time.Now())

	var native []*Post
	var objs []*app.RucciPostLink
	err := m.Db.Scopes(PostFilterByUser(userID, &m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Post", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.PostToRucciPostLink())
	}

	return objs
}

// PostToRucciPostLink returns the RucciPost representation of Post.
func (m *Post) PostToRucciPostLink() *app.RucciPostLink {
	post := &app.RucciPostLink{}
	post.ID = m.ID

	return post
}

// OneRucciPostLink returns an array of view: link.
func (m *PostDB) OneRucciPostLink(ctx context.Context, id int, userID int) (*app.RucciPostLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "rucciPost", "onerucciPostlink"}, time.Now())

	var native Post
	err := m.Db.Scopes(PostFilterByUser(userID, &m.Db)).Table(m.TableName()).Preload("User").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Post", "error", err.Error())
		return nil, err
	}

	view := *native.PostToRucciPostLink()
	return &view, err
}
