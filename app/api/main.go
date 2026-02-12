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

	router.Run()
	defer CloseConnection()
}