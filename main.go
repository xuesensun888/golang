package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	//"os/exec"
	"path/filepath"

	//"os/exec"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 设置静态文件目录，用于存放上传的文件
	r.Static("/uploads", "./uploads")

	// 上传文件的路由处理
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("上传文件失败: %s", err.Error()))
			return
		}

		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, "/opt/cargo/"+filename); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("保存文件失败: %s", err.Error()))
			return
		} else {
			out, err := exec.Command("date").Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("this date is %s", out)
		}
		c.String(http.StatusOK, "文件上传成功")
	})

	fmt.Println("Server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
