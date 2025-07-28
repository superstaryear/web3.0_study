package main

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Name  string
	Posts []Post
}

type Post struct {
	ID            uint
	Name          string
	Words         int
	CommentStatus string
	User          *User
	UserID        uint
	Comments      []Comment
}

/*
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
*/
func (post *Post) BeforeCreate(db *gorm.DB) error {
	length := len(post.Name)
	post.Words = length
	return nil
}

type Comment struct {
	ID     uint
	Remark string
	PostId uint
	Post   *Post
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Print(&comment)
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", comment.PostId).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", comment.PostId).Update("comment_status", "无评论").Error
	}
	return nil
}

type MaxResult struct {
	Num uint
	Pid uint
}

/*
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。,
编写Go代码，使用Gorm创建这些模型对应的数据库表
*/
func main() {
	DB.AutoMigrate(&User{}, &Post{}, &Comment{})
	DB.Create(User{
		ID:   1,
		Name: "张三",
		Posts: []Post{
			{ID: 1, Name: "python", Comments: []Comment{
				Comment{ID: 1, Remark: "python是世界上最好的语言"},
			}},
			{ID: 2, Name: "java", Comments: []Comment{
				{ID: 2, Remark: "java是世界上最好的语言"},
				{ID: 3, Remark: "不接受反驳，哈哈"},
			}},
			{ID: 3, Name: "php", Comments: []Comment{
				{ID: 4, Remark: "php是世界上最好的语言"},
			}},
		},
	})
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var user User
	DB.Preload("Posts").Preload("Posts.Comments").Take(&user, 1)
	fmt.Println(user)
	////编写Go代码，使用Gorm查询评论数量最多的文章信息
	var result MaxResult
	maxQuerySql := "select  count(*) as num,p.id as pid from posts p \nleft join comments c on p.id  = c.post_id \ngroup by c.post_id  order by num  desc limit 1"
	err := DB.Raw(maxQuerySql).Scan(&result).Error
	if err != nil {
		panic(err)
	}
	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
	var comment Comment
	DB.Take(&comment, 1)
	DB.Delete(&comment)
}
