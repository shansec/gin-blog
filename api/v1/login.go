package v1

import (
	"github.com/gin-gonic/gin"
	"myginblog/middleware"
	"myginblog/model"
	"myginblog/utils/errMsg"
	"net/http"
)

// Login 登录
func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	var token string
	var code int

	code = model.CheckLogin(user.Username, user.Password)
	if code == errMsg.SUCCESS {
		token, _ = middleware.SetToken(user.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetCodeMsg(code),
		"token":   token,
	})
}
