package routers

import (
	"gotest/gin/controller/admin"
	"gotest/gin/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRoutes := r.Group("/admin", middlewares.InitMiddleware)
	{
		adminRoutes.GET("/", admin.IndexController{}.Index)
		adminRoutes.GET("/user", admin.UserController{}.Index)
		adminRoutes.GET("/user/add", admin.UserController{}.Add)
		adminRoutes.GET("/user/Edit", admin.UserController{}.Edit)

		adminRoutes.GET("article", admin.ArticleController{}.Index)
		adminRoutes.GET("article/add", admin.ArticleController{}.Add)
		adminRoutes.GET("article/edit", admin.ArticleController{}.Edit)
	}
}
