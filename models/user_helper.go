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

// ListRucciUser returns an array of view: default.
func (m *UserDB) ListRucciUser(ctx context.Context) []*app.RucciUser {
	defer goa.MeasureSince([]string{"goa", "db", "rucciUser", "listrucciUser"}, time.Now())

	var native []*User
	var objs []*app.RucciUser
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToRucciUser())
	}

	return objs
}

// UserToRucciUser returns the RucciUser representation of User.
func (m *User) UserToRucciUser() *app.RucciUser {
	user := &app.RucciUser{}
	user.Disabled = m.Disabled
	user.Email = m.Email
	user.ID = m.ID
	user.Surname = &m.Surname

	return user
}

// OneRucciUser returns an array of view: default.
func (m *UserDB) OneRucciUser(ctx context.Context, id int) (*app.RucciUser, error) {
	defer goa.MeasureSince([]string{"goa", "db", "rucciUser", "onerucciUser"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Preload("Posts").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToRucciUser()
	return &view, err
}

// MediaType Retrieval Functions

// ListRucciUserLink returns an array of view: link.
func (m *UserDB) ListRucciUserLink(ctx context.Context) []*app.RucciUserLink {
	defer goa.MeasureSince([]string{"goa", "db", "rucciUser", "listrucciUserlink"}, time.Now())

	var native []*User
	var objs []*app.RucciUserLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToRucciUserLink())
	}

	return objs
}

// UserToRucciUserLink returns the RucciUser representation of User.
func (m *User) UserToRucciUserLink() *app.RucciUserLink {
	user := &app.RucciUserLink{}
	user.ID = m.ID

	return user
}

// OneRucciUserLink returns an array of view: link.
func (m *UserDB) OneRucciUserLink(ctx context.Context, id int) (*app.RucciUserLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "rucciUser", "onerucciUserlink"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Preload("Posts").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToRucciUserLink()
	return &view, err
}

// MediaType Retrieval Functions

// ListRucciUserTiny returns an array of view: tiny.
func (m *UserDB) ListRucciUserTiny(ctx context.Context) []*app.RucciUserTiny {
	defer goa.MeasureSince([]string{"goa", "db", "rucciUser", "listrucciUsertiny"}, time.Now())

	var native []*User
	var objs []*app.RucciUserTiny
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToRucciUserTiny())
	}

	return objs
}

// UserToRucciUserTiny returns the RucciUser representation of User.
func (m *User) UserToRucciUserTiny() *app.RucciUserTiny {
	user := &app.RucciUserTiny{}
	user.Email = m.Email
	user.ID = m.ID

	return user
}

// OneRucciUserTiny returns an array of view: tiny.
func (m *UserDB) OneRucciUserTiny(ctx context.Context, id int) (*app.RucciUserTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "rucciUser", "onerucciUsertiny"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Preload("Posts").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToRucciUserTiny()
	return &view, err
}
