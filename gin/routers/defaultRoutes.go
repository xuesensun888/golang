package routers

import (
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRoutes := r.Group("/default")
	{
		defaultRoutes.GET("/", func(c *gin.Context) {
			c.String(200, "前台首页")
		})
		defaultRoutes.GET("/user", func(c *gin.Context) {
			c.String(200, "前台用户列表列表")
		})
		defaultRoutes.GET("news", func(c *gin.Context) {
			c.String(200, "新闻")
		})
	}
}
