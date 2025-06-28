package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// StatusSuccess http 200 response
func StatusSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data, // 正确使用数据
	})
}

// StatusNotFound http 404 response
func StatusNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "not found",
		"data":    "null",
	})
}

// StatusRequestNotAllowed http 403 response
func StatusRequestNotAllowed(c *gin.Context, data interface{}) {
	c.JSON(http.StatusForbidden, gin.H{
		"message": "request not allowed",
		"data":    "null",
	})
}
