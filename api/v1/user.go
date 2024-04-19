package v1

import (
	"github.com/gin-gonic/gin"
	"myginblog/model"
	"myginblog/utils/errMsg"
	"net/http"
	"strconv"
)

var code int

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)

	code = model.CheckUser(user.Username)
	if code == errMsg.ERROR_USERNAME_USED {
		code = errMsg.ERROR_USERNAME_USED
	}
	if code == errMsg.SUCCESS || code == errMsg.ERROR_USER_NOT_EXIST {
		code = errMsg.SUCCESS
		user.Password = model.EncryptPwd(user.Password)
		model.CreateUser(&user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		//"data":    user,
		"message": errMsg.GetCodeMsg(code),
	})
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": errMsg.GetCodeMsg(code),
		},
	)

}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	username := c.Query("username")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	users, total, code := model.GetUsers(username, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"total":   total,
		"message": errMsg.GetCodeMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&user)

	code = model.CheckUpUser(id, user.Username)
	if code == errMsg.SUCCESS {
		model.EditUser(id, &user)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetCodeMsg(code),
	})
}

// ChangeUserPwd 修改密码
func ChangeUserPwd(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&user)

	user.Password = model.EncryptPwd(user.Password)
	code = model.ChangeUserPass(id, &user)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetCodeMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetCodeMsg(code),
	})
}
