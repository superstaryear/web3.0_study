package model

import (
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

func (UserEntity) TableName() string {
	return "users" // 新的表名
}

type PostEntity struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    UserEntity
}

func (PostEntity) TableName() string {
	return "posts" // 新的表名
}

type CommentEntity struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    UserEntity
	PostID  uint
	Post    PostEntity
}

func (CommentEntity) TableName() string {
	return "comments" // 新的表名
}
