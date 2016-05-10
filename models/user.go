//************************************************************************//
// API "rucci.io": Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/goa-blog
// --design=goa-blog/design
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

// Atlas User
type User struct {
	ID           int `gorm:"primary_key"` // primary key
	CreatedAt    time.Time
	DeletedAt    *time.Time
	Disabled     bool
	Email        string
	GivenName    string
	PasswordHash string
	PasswordSalt string
	Posts        []Post // has many Posts
	Surname      string
	UpdatedAt    time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m User) TableName() string {
	return "users"

}

// UserDB is the implementation of the storage interface for
// User.
type UserDB struct {
	Db gorm.DB
}

// NewUserDB creates a new storage type.
func NewUserDB(db gorm.DB) *UserDB {
	return &UserDB{Db: db}
}

// DB returns the underlying database.
func (m *UserDB) DB() interface{} {
	return &m.Db
}

// UserStorage represents the storage interface.
type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) []User
	Get(ctx context.Context, id int) (User, error)
	Add(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error

	ListUser(ctx context.Context) []*app.User
	OneUser(ctx context.Context, id int) (*app.User, error)

	ListUserLink(ctx context.Context) []*app.UserLink
	OneUserLink(ctx context.Context, id int) (*app.UserLink, error)

	ListUserTiny(ctx context.Context) []*app.UserTiny
	OneUserTiny(ctx context.Context, id int) (*app.UserTiny, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserDB) TableName() string {
	return "users"

}

// CRUD Functions

// Get returns a single User as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserDB) Get(ctx context.Context, id int) (User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "get"}, time.Now())

	var native User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return User{}, nil
	}

	return native, err
}

// List returns an array of User
func (m *UserDB) List(ctx context.Context) []User {
	defer goa.MeasureSince([]string{"goa", "db", "user", "list"}, time.Now())

	var objs []User
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *UserDB) Add(ctx context.Context, model *User) (*User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error updating User", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *UserDB) Update(ctx context.Context, model *User) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "delete"}, time.Now())

	var obj User

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error retrieving User", "error", err.Error())
		return err
	}

	return nil
}
