package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"operation/models"
	"time"

	"github.com/gin-gonic/gin"
)

type WebhookMessage struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

// type faultInfo struct {
// 	ID     string `json:"id"`
// 	Title  string `json:"title"`
// 	Reason string `json:"reason"`
// }

func Handler_fault(c *gin.Context) {
	webhook := "https://open.feishu.cn/open-apis/bot/v2/hook/94c3a628-9242-4442-9009-10aa323137bd"
	var fault models.FaultReport

	err := c.ShouldBind(&fault)
	if err != nil {
		fmt.Println("err:", err)
	}

	if fault.Timestamp.IsZero() {
		fault.Timestamp = time.Now()
		fmt.Println(fault.Timestamp)
	}

	err = models.Conn.DB.Create(&fault).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	message := WebhookMessage{
		MsgType: "text",
	}
	message.Content.Text = fmt.Sprintf("故障通知\n故障ID: %s\n故障标签: %s\n故障原因:%s\n", fault.ID, fault.Title, fault.Description)
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("jsondata error", err)
	}
	http.Post(webhook, "application/json", bytes.NewBuffer(jsonData))
	c.JSON(http.StatusOK, gin.H{
		"data": fault,
	})
}
