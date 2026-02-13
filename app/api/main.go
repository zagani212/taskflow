package main

import (
	"github.com/gin-gonic/gin"
)

func check(err error){
	if err != nil {
		panic(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main () {
	router := gin.Default()
	CreateConnection()
	router.Use(CORSMiddleware())
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