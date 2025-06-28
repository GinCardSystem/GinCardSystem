package routes

import (
	user2 "GinCardSystem/common/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRoutes(user *gin.Engine) {
	v1 := user.Group("/api/v1/user")

	// User login routes.
	v1.GET("/login", user2.Login)
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
