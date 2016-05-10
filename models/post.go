//************************************************************************//
// API "rucci.io": Models
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

// Blog Post
type Post struct {
	ID        int `gorm:"primary_key"` // primary key
	Content   string
	CreatedAt time.Time
	DeletedAt *time.Time
	Title     string
	UpdatedAt time.Time
	UserID    int // has many Post
	User      User
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Post) TableName() string {
	return "posts"

}

// PostDB is the implementation of the storage interface for
// Post.
type PostDB struct {
	Db gorm.DB
}

// NewPostDB creates a new storage type.
func NewPostDB(db gorm.DB) *PostDB {
	return &PostDB{Db: db}
}

// DB returns the underlying database.
func (m *PostDB) DB() interface{} {
	return &m.Db
}

// PostStorage represents the storage interface.
type PostStorage interface {
	DB() interface{}
	List(ctx context.Context) []Post
	Get(ctx context.Context, id int) (Post, error)
	Add(ctx context.Context, post *Post) (*Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int) error

	ListPost(ctx context.Context, userID int) []*app.Post
	OnePost(ctx context.Context, id int, userID int) (*app.Post, error)

	ListPostLink(ctx context.Context, userID int) []*app.PostLink
	OnePostLink(ctx context.Context, id int, userID int) (*app.PostLink, error)

	UpdateFromCreatePostPayload(ctx context.Context, payload *app.CreatePostPayload, id int) error

	UpdateFromUpdatePostPayload(ctx context.Context, payload *app.UpdatePostPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *PostDB) TableName() string {
	return "posts"

}

// Belongs To Relationships

// PostFilterByUser is a gorm filter for a Belongs To relationship.
func PostFilterByUser(userID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {
	if userID > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", userID)
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Post as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *PostDB) Get(ctx context.Context, id int) (Post, error) {
	defer goa.MeasureSince([]string{"goa", "db", "post", "get"}, time.Now())

	var native Post
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return Post{}, nil
	}

	return native, err
}

// List returns an array of Post
func (m *PostDB) List(ctx context.Context) []Post {
	defer goa.MeasureSince([]string{"goa", "db", "post", "list"}, time.Now())

	var objs []Post
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error listing Post", "error", err.Error())
		return objs
	}

	return objs
}

// Add creates a new record.  /// Maybe shouldn't return the model, it's a pointer.
func (m *PostDB) Add(ctx context.Context, model *Post) (*Post, error) {
	defer goa.MeasureSince([]string{"goa", "db", "post", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error updating Post", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *PostDB) Update(ctx context.Context, model *Post) error {
	defer goa.MeasureSince([]string{"goa", "db", "post", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		return err
	}
	err = m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *PostDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "post", "delete"}, time.Now())

	var obj Post

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error retrieving Post", "error", err.Error())
		return err
	}

	return nil
}

// PostFromCreatePostPayload Converts source CreatePostPayload to target Post model
// only copying the non-nil fields from the source.
func PostFromCreatePostPayload(payload *app.CreatePostPayload) *Post {
	post := &Post{}
	post.Content = payload.Content
	post.Title = payload.Title

	return post
}

// UpdateFromCreatePostPayload applies non-nil changes from CreatePostPayload to the model and saves it
func (m *PostDB) UpdateFromCreatePostPayload(ctx context.Context, payload *app.CreatePostPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "post", "updatefromcreatePostPayload"}, time.Now())

	var obj Post
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Post", "error", err.Error())
		return err
	}
	obj.Content = payload.Content
	obj.Title = payload.Title

	err = m.Db.Save(&obj).Error
	return err
}

// PostFromUpdatePostPayload Converts source UpdatePostPayload to target Post model
// only copying the non-nil fields from the source.
func PostFromUpdatePostPayload(payload *app.UpdatePostPayload) *Post {
	post := &Post{}
	if payload.Content != nil {
		post.Content = *payload.Content
	}
	if payload.Title != nil {
		post.Title = *payload.Title
	}

	return post
}

// UpdateFromUpdatePostPayload applies non-nil changes from UpdatePostPayload to the model and saves it
func (m *PostDB) UpdateFromUpdatePostPayload(ctx context.Context, payload *app.UpdatePostPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "post", "updatefromupdatePostPayload"}, time.Now())

	var obj Post
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Post", "error", err.Error())
		return err
	}
	if payload.Content != nil {
		obj.Content = *payload.Content
	}
	if payload.Title != nil {
		obj.Title = *payload.Title
	}

	err = m.Db.Save(&obj).Error
	return err
}
