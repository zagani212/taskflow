package main

import (
	"github.com/gin-gonic/gin"
)

type Task struct{
	Id int "`json:id`"
	Name string "`json:name`"
	Status string "`json:status`"
}

func createTask(c *gin.Context) {
	var body Task
	err := c.ShouldBindJSON(&body)
	check(err)
	_, err1 := db.Query("INSERT INTO TASKS (name, status) VALUES ($1, $2)", body.Name, body.Status)
	check(err1)
	c.JSON(201, gin.H{
		"code": 201,
		"message": "task created",
	})
}

func updateTask(c *gin.Context) {
	var body Task
	err := c.ShouldBindJSON(&body)
	id := c.Param("id")
	check(err)
	_, err1 := db.Query("UPDATE TASKS SET name = $1, status = $2 WHERE id = $3", body.Name, body.Status, id)
	check(err1)
	c.JSON(201, gin.H{
		"code": 200,
		"message": "task modified",
	})
}

func removeTask(c *gin.Context) {
	var body Task
	err := c.ShouldBindJSON(&body)
	id := c.Param("id")
	check(err)
	_, err1 := db.Query("DELETE FROM TASKS WHERE id = $1", id)
	check(err1)
	c.JSON(201, gin.H{
		"code": 204,
		"message": "task removed",
	})
}

func getTasks(c *gin.Context) {
	rows, err1 := db.Query("SELECT * FROM TASKS")
	check(err1)
	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.Name, &task.Status)
		tasks = append(tasks, task)
	}
	c.JSON(201, gin.H{
		"code": 200,
		"tasks": tasks,
	})
}