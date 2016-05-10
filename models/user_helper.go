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

// ListUser returns an array of view: default.
func (m *UserDB) ListUser(ctx context.Context) []*app.User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuser"}, time.Now())

	var native []*User
	var objs []*app.User
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUser())
	}

	return objs
}

// UserToUser returns the User representation of User.
func (m *User) UserToUser() *app.User {
	user := &app.User{}
	user.Disabled = m.Disabled
	user.Email = m.Email
	user.ID = m.ID
	user.Surname = &m.Surname

	return user
}

// OneUser returns an array of view: default.
func (m *UserDB) OneUser(ctx context.Context, id int) (*app.User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuser"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Preload("Posts").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUser()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserLink returns an array of view: link.
func (m *UserDB) ListUserLink(ctx context.Context) []*app.UserLink {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listuserlink"}, time.Now())

	var native []*User
	var objs []*app.UserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUserLink())
	}

	return objs
}

// UserToUserLink returns the User representation of User.
func (m *User) UserToUserLink() *app.UserLink {
	user := &app.UserLink{}
	user.ID = m.ID

	return user
}

// OneUserLink returns an array of view: link.
func (m *UserDB) OneUserLink(ctx context.Context, id int) (*app.UserLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneuserlink"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Preload("Posts").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUserLink()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserTiny returns an array of view: tiny.
func (m *UserDB) ListUserTiny(ctx context.Context) []*app.UserTiny {
	defer goa.MeasureSince([]string{"goa", "db", "user", "listusertiny"}, time.Now())

	var native []*User
	var objs []*app.UserTiny
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUserTiny())
	}

	return objs
}

// UserToUserTiny returns the User representation of User.
func (m *User) UserToUserTiny() *app.UserTiny {
	user := &app.UserTiny{}
	user.Email = m.Email
	user.ID = m.ID

	return user
}

// OneUserTiny returns an array of view: tiny.
func (m *UserDB) OneUserTiny(ctx context.Context, id int) (*app.UserTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "oneusertiny"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Preload("Posts").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUserTiny()
	return &view, err
}
