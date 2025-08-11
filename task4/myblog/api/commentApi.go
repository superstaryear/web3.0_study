package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"myblog/common"
	"myblog/model"
)

/*
实现评论的创建功能，已认证的用户可以对文章发表评论
use  jwtHandler 中间件
*/
func CommentCreate(c *gin.Context) {
	var commentCreateRequest model.CommentCreateRequest
	if err := c.ShouldBindJSON(&commentCreateRequest); err != nil {
		errMsg := common.GetValidMsg(err, &commentCreateRequest)
		common.Error(c, 500, "评论创建请求参数错误"+errMsg)
		return
	}
	var commentEntity model.CommentEntity
	copier.Copy(&commentEntity, &commentCreateRequest)
	userId, _ := c.Get("userId")
	//userName, _ := c.Get("userName")
	//fmt.Printf("type:%T userId:%v \n", userId, userId)
	//fmt.Printf("type:%T userName:%v\n", userName, userName)
	commentEntity.UserID = userId.(uint)
	if err := common.DB.Create(&commentEntity).Error; err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	var commentCreateResp model.CommentCreateResponse = model.CommentCreateResponse{
		CommentId: commentEntity.ID,
	}
	common.Success(c, commentCreateResp, "创建评论成功")
}

/*
读取所有评论列表
*/
func CommentQuery(c *gin.Context) {
	var commentQueryReq model.CommentQueryRequest
	err := c.ShouldBindJSON(&commentQueryReq)
	if err != nil {
		errMsg := common.GetValidMsg(err, &commentQueryReq)
		common.Error(c, 500, "读取所有评论列表请求参数错误"+errMsg)
	}
	//通过主键id查询文章
	var commentEntity []model.CommentEntity
	if err := common.DB.Preload("User").Preload("Post").Where("post_id=?", commentQueryReq.PostId).Find(&commentEntity).Error; err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	var commentQueryResp []model.CommentQueryResponse
	for _, comment := range commentEntity {
		var data model.CommentQueryResponse = model.CommentQueryResponse{}
		copier.Copy(&data, &comment)
		commentQueryResp = append(commentQueryResp, data)
	}
	common.Success(c, commentQueryResp, "查询评论列表成功")
}
