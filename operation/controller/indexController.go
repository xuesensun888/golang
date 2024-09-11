package controller

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (con IndexController) Upload(c *gin.Context) {
	file, err := c.FormFile("face")
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
}
