package controller

import (
	"gvp/common"
	"gvp/model"
	"gvp/response"
	"gvp/vo"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IPostController interface {
	RestController
	PageList(c *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}

func (pc PostController) Create(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	if err := c.ShouldBind(&requestPost); err != nil {
		response.Fail(c, nil, "绑定失败")
		return
	}
	user, _ := c.Get("user")
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	if err := pc.DB.Create(&post).Error; err != nil {
		panic(err)
	}
	response.Success(c, nil, "创建成功")
}

func (pc PostController) Update(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	if err := c.ShouldBind(&requestPost); err != nil {
		response.Fail(c, nil, "绑定失败")
		return
	}
	postId := c.Params.ByName("id")

	var post model.Post
	if err := pc.DB.Where("id=?", postId).First(&post).Error; err != nil {
		response.Fail(c, nil, "未找到文章")
		return
	}

	user, _ := c.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(c, nil, "没有权限")
		return
	}

	if err := pc.DB.Model(&post).Updates(&requestPost).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}

	response.Success(c, gin.H{"post": post}, "更新成功")

}

func (pc PostController) Show(c *gin.Context) {
	postId := c.Params.ByName("id")

	var post model.Post
	if err := pc.DB.Preload("Category").Where("id=?", postId).First(&post).Error; err != nil {
		response.Fail(c, nil, "未找到文章")
		return
	}

	response.Success(c, gin.H{"post": post}, "")
}

func (pc PostController) Delete(c *gin.Context) {
	postId := c.Params.ByName("id")

	var post model.Post
	if err := pc.DB.Where("id=?", postId).First(&post).Error; err != nil {
		response.Fail(c, nil, "未找到文章")
		return
	}

	user, _ := c.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(c, nil, "没有权限")
		return
	}

	pc.DB.Delete(&post)
	response.Success(c, nil, "删除成功")
}

func (pc PostController) PageList(c *gin.Context){
	pageNum,_:=strconv.Atoi(c.DefaultQuery("pageNum","1"))
	pageSize,_:=strconv.Atoi(c.DefaultQuery("pageSize","1"))

	var posts []model.Post
	pc.DB.Order("created_at desc").Offset((pageNum-1)*pageSize).Limit(pageSize).Find(&posts)

	var total int64
	pc.DB.Model(model.Post{}).Count(&total)

	response.Success(c,gin.H{"posts":posts,"total":total},"success")
}