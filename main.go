package main

import (
	"net/http"

	"gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// r.POST("/users/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	// plain text response
	// 	// c.String(http.StatusOK, name)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"name": name,
	// 	})
	// })

	r.Run()
	// you can custumize the server port like this, default is 8080
	// r.Run(":8080")
}
