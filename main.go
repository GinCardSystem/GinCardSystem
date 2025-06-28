package main

import (
	"GinCardSystem/config"
	"GinCardSystem/internal/db"
	"GinCardSystem/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// 加载配置
	config.LoadConfigFile()

	db.TestQuickStart()

	e := gin.Default()
	e.GET("/", ping)

	routes.UserRoutes(e)

	err := e.Run(":8080")

	if err != nil {
		panic(err)
	}

}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
