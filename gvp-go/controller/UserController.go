package controller

import (
	"gvp/common"
	"gvp/model"
	"gvp/util"
	"log"
	"net/http"

	"gvp/dto"
	"gvp/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {

	db := common.GetDB()

	//get data
	var requestUser = model.User{}
	c.Bind(&requestUser)
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//check data
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "phone must be longer than 11")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "password can not be less than 6")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if IsTelephoneExist(db, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "telephone already exist")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "bcrypt wrong")
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	db.Create(&newUser)

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "system worng")
		log.Printf("token generate error: %v", err)
		return
	}

	response.Success(c, gin.H{"token": token}, "user register successful")
}

func Login(c *gin.Context) {
	db := common.GetDB()
	var requestUser = model.User{}
	c.Bind(&requestUser)

	//get data
	telephone := requestUser.Telephone
	password := requestUser.Password

	//check data
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "phone must be 11")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "password can not be less than 6")
		return
	}

	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "user not exist")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "wrong password or telephone")
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "system worng")
		log.Printf("token generate error: %v", err)
		return
	}

	response.Success(c, gin.H{"token": token}, "successful login")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).Find(&user)
	return user.ID != 0
}
