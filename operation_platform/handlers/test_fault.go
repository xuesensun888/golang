package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Hander_test(c *gin.Context) {
	value := c.Query("user_id")
	fmt.Println(value)
}
