package v1

import (
	"github.com/gin-gonic/gin"
	"myginblog/model"
	"myginblog/utils/errMsg"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UpLoadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetCodeMsg(code),
		"url":     url,
	})
}
