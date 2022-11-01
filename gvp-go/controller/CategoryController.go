package controller

import (
	"gvp/model"
	"gvp/repository"
	"gvp/response"
	"gvp/vo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: repository}
}

func (cc CategoryController) Create(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "分类名称错误")
		return
	}
	category, err := cc.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(c, gin.H{"category": category}, "")
}

func (cc CategoryController) Update(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "分类名称错误")
		return
	}

	categoryID, _ := strconv.Atoi(c.Params.ByName("id"))

	updateCategory, err := cc.Repository.SelectByID(categoryID)
	if err != nil {
		response.Fail(c, nil, "分类不存在")
	}
	category, err := cc.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(c, gin.H{"category": category}, "修改成功")
}
func (cc CategoryController) Show(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Params.ByName("id"))

	category, err := cc.Repository.SelectByID(categoryID)
	if err != nil {
		response.Fail(c, nil, "分类不存在")
	}
	response.Success(c, gin.H{"category": category}, "")
}
func (cc CategoryController) Delete(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Params.ByName("id"))
	if err := cc.Repository.DeleteByID(categoryID); err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "")
}
