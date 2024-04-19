package v1

import (
	"github.com/gin-gonic/gin"
	"myginblog/model"
	"myginblog/utils/errMsg"
	"net/http"
	"strconv"
)

// GetProfile 获取个人信息设置
func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetProfileInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errMsg.GetCodeMsg(code),
	})
}

// UpdateProfile 更新个人设置
func UpdateProfile(c *gin.Context) {
	var profile model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&profile)

	code := model.UpdateProfileInfo(id, &profile)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetCodeMsg(code),
	})
}
