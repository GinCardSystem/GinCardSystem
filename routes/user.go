package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"GinCardSystem/common/services"
)

func UserRoutes(user *gin.Engine) {
	v1 := user.Group("/api/v1/user")

	// User login routes.
	v1.GET("/login", services.Login)
	v1.POST("/login", ToDoFunc)
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
