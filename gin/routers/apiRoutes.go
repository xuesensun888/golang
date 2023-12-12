package routers

import (
	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/", func(c *gin.Context) {
			c.String(200, "api首页")
		})
		apiRoutes.GET("/userlist", func(c *gin.Context) {
			c.String(200, "一个userlist接口")
		})
		apiRoutes.GET("/sxs", func(c *gin.Context) {
			c.String(200, "sxs的接口")
		})
	}
}
