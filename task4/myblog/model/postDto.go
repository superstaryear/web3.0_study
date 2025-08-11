package model

/*
创建请求体
*/
type PostCreateRequest struct {
	Title   string `json:"title" binding:"required" msg:"文章标题不能为空"`
	Content string `json:"content" binding:"required" msg:"文章内容不能为空"`
}

/*
创建返回体
*/
type PostCreateResponse struct {
	PostId uint   `json:"postId"`
	Title  string `json:"title"`
}

/*
* 单一查询请求体
 */
type PostQueryRequest struct {
	PostId uint `uri:"postId"`
}

/*
单一查询返回体
*/
type PostQueryResponse struct {
	Id      uint
	Title   string
	Content string
	UserID  uint
	User    UserEntity
}

/*
分页查询请求体
*/
type PostPageRequest struct {
	//页码
	Page int `json:"page" binding:"required,min=1" msg:"页码不能为空"`
	//每页数量
	Limit int `json:"limit" binding:"required,min=1,max=100" msg:"每页数量不能为空"`
	//文章标题
	Title string `json:"title"`
}

/*
分页查询返回体
*/
type PostPageResponse struct {
	Data []PostQueryResponse `json:"data"`
	Meta struct {
		Page      int `json:"page"`
		PerPage   int `json:"per_page"`
		Total     int `json:"total"`
		TotalPage int `json:"total_page"`
	} `json:"meta"`
}

/*
更新请求体
*/
type PostUpdateRequest struct {
	PostId  uint   `json:"postId" binding:"required" msg:"更新主键id不能为空"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

/*
更新返回体
*/
type PostUpdateResponse struct {
	Success bool
}

/*
删除请求体
*/
type PostDeleteRequest struct {
	PostId []uint `json:"postId" binding:"required" msg:"删除主键id不能为空"`
}

/*
删除返回体
*/
type PostDeleteResponse struct {
	Success bool
}
