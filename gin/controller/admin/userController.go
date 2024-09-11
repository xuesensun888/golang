package admin

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	c.String(200, "用户列表1")
}

//	func (con UserController) Add(c *gin.Context) {
//		c.String(200, "用户添加2")
//	}
func (con UserController) Edit(c *gin.Context) {
	c.String(200, "用户修改3")
}
func (con UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/add.html", gin.H{
		"title": "首页",
	})
}
func (con UserController) DoUpload(c *gin.Context) {
	//获取表单中的username
	username := c.PostForm("username")
	//获取文件
	file, err := c.FormFile("file")
	//判断文件是否存在
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"sussess": "fail",
			"message": err.Error(),
		})
		return
	}
	//设置需要上传的目录
	dst := path.Join("/opt/cargo", file.Filename)
	//上传文件到指定目录
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"success":  "true",
		"username": username,
		"dst":      dst,
	})
}
