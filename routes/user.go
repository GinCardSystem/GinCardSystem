package routes

import (
	"GinCardSystem/common/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.Engine) {
	v1 := c.Group("/api/v1/user")

	// User login routes.
	v1.GET("/login", user.Login)
	v1.POST("/login", user.Login)
	v1.GET("/logout", ToDoFunc)

	// User register routes.
	v1.GET("/register", ToDoFunc)
	v1.POST("/register", ToDoFunc)

	// User security text routes.
	v1.GET("/captcha", ToDoFunc)
	v1.POST("/captcha", ToDoFunc)
}

func ToDoFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ToDo",
	})
}
