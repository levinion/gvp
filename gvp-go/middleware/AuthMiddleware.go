package middleware

import (
	"gvp/common"
	"gvp/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString:=c.GetHeader("Authorization")
		if tokenString=="" || !strings.HasPrefix(tokenString,"Bearer"){
			c.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"no auth"})
			c.Abort()
			return
		}

		tokenString=tokenString[7:]
		token,claims,err:=common.ParseToken(tokenString)
		if err!=nil || !token.Valid{
			c.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"no auth"})
			c.Abort()
			return
		}

		userId:=claims.UserId
		db:=common.GetDB()
		var user model.User
		db.First(&user,userId)
		if user.ID==0{
			c.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"no auth"})
			c.Abort()
			return
		}

		c.Set("user",user)
		c.Next()
	}
}