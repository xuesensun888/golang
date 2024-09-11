package routers

import (
	"operation/controller"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/upload")
	{
		adminRouters.POST("/", controller.IndexController{}.Upload)
		adminRouters.POST("/fault", controller.FaultController{}.Fault)
		adminRouters.GET("/getfault/:id", controller.FaultController{}.GetFault)

	}
}
