package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"math"
	"myblog/common"
	"myblog/model"
)

/*
实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容
use  jwtHandler 中间件
*/
func PostCreate(c *gin.Context) {
	var postCreateReq model.PostCreateRequest
	if err := c.ShouldBindJSON(&postCreateReq); err != nil {
		errMsg := common.GetValidMsg(err, &postCreateReq)
		common.Error(c, 500, "文章创建请求参数错误"+errMsg)
		return
	}
	var postEntity model.PostEntity
	copier.Copy(&postEntity, &postCreateReq)
	userId, _ := c.Get("userId")
	//userName, _ := c.Get("userName")
	//fmt.Printf("type:%T userId:%v \n", userId, userId)
	//fmt.Printf("type:%T userName:%v\n", userName, userName)
	postEntity.UserID = userId.(uint)
	if err := common.DB.Create(&postEntity).Error; err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	var postCreateResp model.PostCreateResponse = model.PostCreateResponse{
		PostId: postEntity.ID,
		Title:  postEntity.Title,
	}
	common.Success(c, postCreateResp, "创建文章成功")
}

/*
单个文章的详细信息
*/
func PostQuery(c *gin.Context) {
	var postQueryReq model.PostQueryRequest
	err := c.ShouldBindUri(&postQueryReq)
	if err != nil {
		errMsg := common.GetValidMsg(err, &postQueryReq)
		common.Error(c, 500, "文章查询单个文章请求参数错误"+errMsg)
	}
	//通过主键id查询文章
	var postEntity model.PostEntity
	if err := common.DB.Preload("User").Take(&postEntity, postQueryReq.PostId).Error; err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	var postQueryResp model.PostQueryResponse
	copier.Copy(&postQueryResp, &postEntity)
	common.Success(c, postQueryResp, "查询文章成功")
}

/*
获取所有文章列表
*/
func PostPage(c *gin.Context) {
	var postPageReq model.PostPageRequest
	if err := c.ShouldBindJSON(&postPageReq); err != nil {
		errMsg := common.GetValidMsg(err, &postPageReq)
		common.Error(c, 500, "获取所有文章列表请求参数错误"+errMsg)
	}
	var posts []model.PostEntity
	var total int64
	// 构建查询
	query := common.DB.Model(&model.PostEntity{}).Preload("User")
	// 添加标题条件
	if postPageReq.Title != "" {
		query = query.Where("title LIKE ?", "%"+postPageReq.Title+"%")
	}
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		common.Error(c, 500, "获取总数失败: "+err.Error())
		return
	}
	// 分页查询
	if err := query.Offset((postPageReq.Page - 1) * postPageReq.Limit).
		Limit(postPageReq.Limit).
		Find(&posts).Error; err != nil {
		common.Error(c, 500, "查询失败: "+err.Error())
		return
	}
	// 计算总页数
	totalPage := int(math.Ceil(float64(total) / float64(postPageReq.Limit)))
	// 转换为响应DTO
	var response model.PostPageResponse
	for _, post := range posts {
		response.Data = append(response.Data, model.PostQueryResponse{
			Id:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			UserID:  post.UserID,
			User:    post.User,
		})
	}

	response.Meta.Page = postPageReq.Page
	response.Meta.PerPage = postPageReq.Limit
	response.Meta.Total = int(total)
	response.Meta.TotalPage = totalPage
	common.Success(c, response, "获取所有文章列表成功")
}

/*
文章更新 只有文章的作者才能更新自己的文章
*/
func PostUpdate(c *gin.Context) {
	var postUpdateReq model.PostUpdateRequest
	if err := c.ShouldBindJSON(&postUpdateReq); err != nil {
		errMsg := common.GetValidMsg(err, &postUpdateReq)
		common.Error(c, 500, "更新文章请求参数错误"+errMsg)
	}
	var postEntity model.PostEntity
	if err := common.DB.Take(&postEntity, postUpdateReq.PostId).Error; err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	userId, _ := c.Get("userId")
	//fmt.Printf("type:%T userId:%v \n", userId, userId)
	//只有文章的作者才能更新自己的文章
	if postEntity.UserID != userId.(uint) {
		common.Error(c, 500, "修改非法")
		return
	}
	if postUpdateReq.Content != "" {
		postEntity.Content = postUpdateReq.Content
	}

	if postUpdateReq.Title != "" {
		postEntity.Title = postUpdateReq.Title
	}
	//更新文章
	err := common.DB.Save(&postEntity).Error
	isSuccess := err == nil
	var postUpdateResp model.PostUpdateResponse = model.PostUpdateResponse{
		Success: isSuccess,
	}
	common.Success(c, postUpdateResp, "更新文章成功")
}

/*
文章删除
*/
func PostDelete(c *gin.Context) {
	var postDeleteReq model.PostDeleteRequest
	if err := c.ShouldBindJSON(&postDeleteReq); err != nil {
		errMsg := common.GetValidMsg(err, &postDeleteReq)
		common.Error(c, 500, errMsg)
		return
	}
	userId, _ := c.Get("userId")
	//fmt.Printf("type:%T userId:%v \n", userId, userId)
	var postEntity []model.PostEntity
	if err := common.DB.Where("user_id = ?", userId.(uint)).Take(&postEntity, postDeleteReq.PostId).Error; err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	err := common.DB.Delete(&postEntity).Error
	isSuccess := err == nil
	var postDeleteResp model.PostDeleteResponse = model.PostDeleteResponse{
		Success: isSuccess,
	}
	common.Success(c, postDeleteResp, "删除文章成功")
}
