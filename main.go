package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type server struct {
	Engine *gin.Engine
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	engine := gin.Default()
	serv := server{Engine: engine}.Engine

	// APIS routes

	serv.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	serv.POST("/login", func(c *gin.Context) {
		type User struct {
			Name  string
			Email string
		}
		var user User

		err := c.BindJSON(&user)

		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid request",
			})
		} else {
			c.JSON(200, user)
		}

	})

	// start server
	fmt.Printf("Server run on port : %s", port)
	serv.Run(":" + port)

}
