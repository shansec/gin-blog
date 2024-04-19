package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myginblog/api/v1"
	"myginblog/middleware"
	"myginblog/utils"
)

func InitRoute() {
	gin.SetMode(utils.Model)
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static", "static/admin/static")
	r.StaticFile("admin/favicon.ico", "static/admin/favicon.ico")

	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 修改密码
		auth.PUT("user/changePw/:id", v1.ChangeUserPwd)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		// 文件上传
		auth.POST("user/upload", v1.Upload)
		// 更新用户设置
		auth.PUT("profile/:id", v1.UpdateProfile)
	}

	public := r.Group("api/v1")
	{
		// User模块的路由接口
		public.POST("user/add", v1.AddUser)
		public.GET("user/get", v1.GetUsers)
		public.GET("user/getInfo/:id", v1.GetUserInfo)
		public.POST("user/login", v1.Login)
		// 分类模块的路由接口
		public.GET("category/get", v1.GetCategory)
		public.GET("category/getInfo/:id", v1.GetCategoryInfo)
		// 文章模块的路由接口
		public.GET("article/get", v1.GetArticle)
		public.GET("article/info/:id", v1.GetArticleInfo)
		public.GET("article/cateArt/:id", v1.GetCateArticleInfo)
		// 获取用户设置
		public.GET("profile/get/:id", v1.GetProfile)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
