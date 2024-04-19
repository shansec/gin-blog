package v1

import (
	"github.com/gin-gonic/gin"
	"myginblog/model"
	"myginblog/utils/errMsg"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBind(&category)
	code = model.CheckCategory(category.Name)
	if code == errMsg.ERROR_CATENAME_USED {
		code = errMsg.ERROR_CATENAME_USED
	}
	if code == errMsg.SUCCESS {
		model.CreateCategory(&category)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		//"data":    category,
		"message": errMsg.GetCodeMsg(code),
	})
}

// GetCategoryInfo 查询单个分类信息
func GetCategoryInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetCate(id)
	maps["name"] = data.Name
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    maps,
		"total":   1,
		"message": errMsg.GetCodeMsg(code),
	})
}

// GetCategory 查询分类列表
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	categoryS, total := model.GetCategory(pageSize, pageNum)
	code = errMsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    categoryS,
		"total":   total,
		"message": errMsg.GetCodeMsg(code),
	})
}

// EditCategory 编辑分类
func EditCategory(c *gin.Context) {
	var category model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&category)

	code = model.CheckCategory(category.Name)
	if code == errMsg.SUCCESS {
		model.EditCategory(id, &category)
	}
	if code == errMsg.ERROR_CATENAME_USED {
		code = errMsg.ERROR_CATENAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetCodeMsg(code),
	})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errMsg.GetCodeMsg(code),
	})
}
