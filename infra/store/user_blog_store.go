package store

import (
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// UserBlogStore accesses users data
type UserBlogStore interface {
	CreateUserBlog(userID model.UserID, blog *record.Blog) error
	UpdateUserBlog(userID model.UserID, blog *record.Blog) error
	DeleteUserBlog(blogID int64) error
}
