package user

import (
	"GinCardSystem/common/response"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		response.StatusSuccess(c, gin.H{
			"jwt": "text",
		})
		return // 添加 return 避免继续执行
	}

	if c.Request.Method == "GET" {
		response.StatusSuccess(c, gin.H{
			"jwt": "text",
		})
		return // 添加 return 避免继续执行
	}
}
