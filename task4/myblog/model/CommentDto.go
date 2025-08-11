package model

/*
创建请求体
*/
type CommentCreateRequest struct {
	PostId  uint   `json:"postId" binding:"required" msg:"文章Id不能为空"`
	Content string `json:"content" binding:"required" msg:"评论内容不能为空"`
}

/*
创建返回体
*/
type CommentCreateResponse struct {
	CommentId uint `json:"commentId"`
}

/*
* 评论列表请求体
 */
type CommentQueryRequest struct {
	PostId uint `json:"postId" binding:"required" msg:"文章Id不能为空"`
}

/*
评论列表返回体
*/
type CommentQueryResponse struct {
	Id      uint
	Content string
	UserID  uint
	User    UserEntity
	PostID  uint
	Post    PostEntity
}
