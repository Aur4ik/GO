package main

import (
	"net/http"
	"taskmanager/db"
	"taskmanager/taskmodel"

	"github.com/gin-gonic/gin"
)

func main() {
	initDB()

	r := gin.Default()

	r.POST("/tasks", createTask)
	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTaskByID)
	r.DELETE("/tasks/:id", deleteTask)

	r.Run(":8080")
}

func createTask(c *gin.Context) {
	var task taskmodel.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный JSON",
		})
		return
	}

	err := db.QueryRow(
		"INSERT INTO tasks(title, description, done) VALUES($1,$2,$3) RETURNING id",
		task.Title, task.Description, task.Done,
	).Scan(&task.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func getTasks(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, description, done FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
		return
	}
	defer rows.Close()
   
	var tasks []taskmodel.Task

	for rows.Next() {
		var t taskmodel.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done)
		tasks = append(tasks, t)
	}

	c.JSON(http.StatusOK, tasks)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	var task taskmodel.Task

	err := db.QueryRow(
		"SELECT id, title, description, done FROM tasks WHERE id=$1",
		id,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Done)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Удалено"})
}
