package main

import (
	"fmt"
	"crypto/md5"
 	"encoding/hex"
	"io"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Email string "`json:email`"
	Password string "`json:password`"
}

func Login(c *gin.Context){
	var body Credentials
	err := c.ShouldBindJSON(&body)
	if err != nil {
		panic(err)
	}
	fmt.Println(body.Email)
	rows, err1 := db.Query("SELECT email, password FROM users WHERE email = $1 LIMIT 1", body.Email)
	if !rows.Next() {
		c.JSON(401, gin.H{
			"code": 401,
			"message": "Bad Credentials",
		})
		return
	}
	if err1 != nil {
		panic(err1)
	}

	var passwordFromDB, email string
	rows.Scan(&email, &passwordFromDB)
	if hash(body.Password) == passwordFromDB {
		c.JSON(200, gin.H{
			"code": 200,
			"message": "OK",
		})
		return
	}
	c.JSON(401, gin.H{
		"code": 401,
		"message": "Bad Credentials",
	})
	
}

func Register(c *gin.Context){
	var body Credentials
	err := c.ShouldBindJSON(&body)
	check(err)
	_, err1 := db.Query("INSERT INTO users (email, password) VALUES ($1, $2)", body.Email, hash(body.Password))
	if err1 != nil {
		panic(err1)
	}
	c.JSON(201, gin.H{
		"code": 201,
		"message": "Account created",
	})
}

func hash(text string) string {
	hasher := md5.New()
	_, err := io.WriteString(hasher, text)
	if err != nil {
	panic(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}