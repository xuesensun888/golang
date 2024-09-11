package routers

import (
	"operation/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutesInit(r *gin.Engine) {
	userRouters := r.Group("/")
	{
		userRouters.POST("/add", controller.UserController{}.Add)
		userRouters.GET("/index", controller.UserController{}.Index)
		userRouters.GET("/edit", controller.UserController{}.Edit)
		userRouters.GET("/delete", controller.UserController{}.Delete)
	}
}
