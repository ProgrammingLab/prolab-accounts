package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
)

// UserBlogStore accesses users data
type UserBlogStore interface {
	ListUserBlogs() ([]*record.Blog, error)
	GetUserBlog(blogID int64) (*record.Blog, error)
	CreateUserBlog(blog *record.Blog) error
	UpdateUserBlog(blog *record.Blog) error
	DeleteUserBlog(blogID int64) error
}
