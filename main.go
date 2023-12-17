package main

import (
	"github.com/TlexCypher/ginAuthenticationCatchUp/controllers"
	"github.com/TlexCypher/ginAuthenticationCatchUp/initializes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializes.LoadVariables()
	initializes.ConnectToDb()
	initializes.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.Run()
}
