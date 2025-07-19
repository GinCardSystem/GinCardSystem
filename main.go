package main

import (
	"GinCardSystem/internal/db"
	"GinCardSystem/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Init() // 初始化数据库连接
	if err := db.Init(); err != nil {
		log.Println("Failed to initialize database:", err)
		panic(err)
	}

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
