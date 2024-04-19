package v1

import (
	"github.com/gin-gonic/gin"
	"myginblog/model"
	"myginblog/utils/errMsg"
	"net/http"
	"strconv"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBind(&article)

	code = model.CreateArticle(&article)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errMsg.GetCodeMsg(code),
	})
}

// GetCateArticleInfo 查询分类下的所有文章
func GetCateArticleInfo(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, total := model.GetCateArticle(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errMsg.GetCodeMsg(code),
	})
}

// GetArticle 查询文章列表
func GetArticle(c *gin.Context) {
	title := c.Query("title")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	articles, code, total := model.GetArticle(title, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"total":   total,
		"message": errMsg.GetCodeMsg(code),
	})
}

// GetArticleInfo 查询单个文章信息
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetSignalArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errMsg.GetCodeMsg(code),
	})
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&article)

	code = model.EditArticle(id, &article)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetCodeMsg(code),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetCodeMsg(code),
	})
}
