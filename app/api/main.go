package main

import (
	"github.com/gin-gonic/gin"
)

func check(err error){
	if err != nil {
		panic(err)
	}
}

func main () {
	router := gin.Default()
	CreateConnection()
	router.POST("/auth/login", Login)
	router.POST("/auth/register", Register)
	router.POST("/tasks", createTask)
	router.GET("/tasks", getTasks)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", removeTask)
	router.GET("/health", func(c *gin.Context){
		c.JSON(200, "OK")
	})

	router.Run()
	defer CloseConnection()
}