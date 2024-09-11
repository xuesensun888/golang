package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"remotehttp/conf"
	"remotehttp/model"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 获取版本信息
func Version(c *gin.Context) {
	app := c.Param("app")
	v, err := model.Conn.GetVersion(app)
	var errCode int
	if err != nil {
		errCode = 9999999
	} else {
		errCode = 0
	}
	c.JSON(http.StatusOK, gin.H{
		"errcode": errCode,
		"data":    v,
	})
}
func Header(c *gin.Context) {
	name := c.Param("name")
	filePath := conf.DefaultConfig.DownloadDir + "/" + name
	fmt.Println(filePath)
	stat, err := os.Stat(filePath)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.Header("Content-Length", fmt.Sprintf("%d", stat.Size()))
	c.Header("Content-Type", "application/octet-stream")
	//c.Header("Accept-Ranges", "bytes")
	c.Writer.Header().Set("Conten-Disposition", "attachment; filename="+name)
	c.Writer.Flush()
}

// 多线程断点续传
func Download(c *gin.Context) {
	s := c.Request.Header.Get("Range")
	// 如果没有使用分段下载，则跳转到单线程下载
	if s == "" {
		SingleDownload(c)
		return
	}

	name := c.Param("name")
	filePath := conf.DefaultConfig.DownloadDir + "/" + name
	f, err := os.Open(filePath)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	defer f.Close()

	ranges := strings.Split(s, "=")[1]
	start, err := strconv.ParseInt(strings.Split(ranges, "-")[0], 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	end, err := strconv.ParseInt(strings.Split(ranges, "-")[1], 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	length := end - start + 1

	c.Header("Content-Length", fmt.Sprintf("%d", length))
	c.Header("Content-Range", fmt.Sprintf("%d-%d/%d", start, end, length)) // "start-end/length"
	c.Header("Last-Modified", "Fri, 19 May 2023 14:38:31 GMT")

	c.Writer.WriteHeader(http.StatusPartialContent)

	_, err = f.Seek(start, io.SeekStart)
	if err != nil {
		c.AbortWithError(http.StatusServiceUnavailable, err)
		return
	}

	sizeof := int64(1024 * 1024)
	b := make([]byte, sizeof)

	for l := length; l != 0; {
		n, err := f.Read(b)
		if err != nil {
			c.AbortWithError(http.StatusServiceUnavailable, err)
			return
		}
		if l < sizeof {
			n = int(l)
		}
		c.Writer.Write(b[:n])
		l -= int64(n)
	}
}

// 单线程下载
func SingleDownload(c *gin.Context) {
	name := c.Param("name")
	filePath := conf.DefaultConfig.DownloadDir + "/" + name
	c.File(filePath)
}
