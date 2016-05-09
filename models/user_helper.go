//************************************************************************//
// API "Atlas": Model Helpers
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

// ListAtlasUser returns an array of view: default.
func (m *UserDB) ListAtlasUser(ctx context.Context) []*app.AtlasUser {
	defer goa.MeasureSince([]string{"goa", "db", "atlasUser", "listatlasUser"}, time.Now())

	var native []*User
	var objs []*app.AtlasUser
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToAtlasUser())
	}

	return objs
}

// UserToAtlasUser returns the AtlasUser representation of User.
func (m *User) UserToAtlasUser() *app.AtlasUser {
	user := &app.AtlasUser{}
	user.Disabled = m.Disabled
	user.Email = m.Email
	user.GivenName = &m.GivenName
	user.ID = m.ID
	user.Surname = &m.Surname

	return user
}

// OneAtlasUser returns an array of view: default.
func (m *UserDB) OneAtlasUser(ctx context.Context, id int) (*app.AtlasUser, error) {
	defer goa.MeasureSince([]string{"goa", "db", "atlasUser", "oneatlasUser"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToAtlasUser()
	return &view, err
}

// MediaType Retrieval Functions

// ListAtlasUserLink returns an array of view: link.
func (m *UserDB) ListAtlasUserLink(ctx context.Context) []*app.AtlasUserLink {
	defer goa.MeasureSince([]string{"goa", "db", "atlasUser", "listatlasUserlink"}, time.Now())

	var native []*User
	var objs []*app.AtlasUserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToAtlasUserLink())
	}

	return objs
}

// UserToAtlasUserLink returns the AtlasUser representation of User.
func (m *User) UserToAtlasUserLink() *app.AtlasUserLink {
	user := &app.AtlasUserLink{}
	user.ID = m.ID

	return user
}

// OneAtlasUserLink returns an array of view: link.
func (m *UserDB) OneAtlasUserLink(ctx context.Context, id int) (*app.AtlasUserLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "atlasUser", "oneatlasUserlink"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToAtlasUserLink()
	return &view, err
}

// MediaType Retrieval Functions

// ListAtlasUserTiny returns an array of view: tiny.
func (m *UserDB) ListAtlasUserTiny(ctx context.Context) []*app.AtlasUserTiny {
	defer goa.MeasureSince([]string{"goa", "db", "atlasUser", "listatlasUsertiny"}, time.Now())

	var native []*User
	var objs []*app.AtlasUserTiny
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToAtlasUserTiny())
	}

	return objs
}

// UserToAtlasUserTiny returns the AtlasUser representation of User.
func (m *User) UserToAtlasUserTiny() *app.AtlasUserTiny {
	user := &app.AtlasUserTiny{}
	user.Email = m.Email
	user.ID = m.ID

	return user
}

// OneAtlasUserTiny returns an array of view: tiny.
func (m *UserDB) OneAtlasUserTiny(ctx context.Context, id int) (*app.AtlasUserTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "atlasUser", "oneatlasUsertiny"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToAtlasUserTiny()
	return &view, err
}
