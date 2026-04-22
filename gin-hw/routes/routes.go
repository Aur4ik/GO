package main

import (
	"taskmanager/handlers"
	"taskmanager/middleware"

	"github.com/gin-gonic/gin"
)
func main(){
r := gin.Default()


r.POST("/register", handlers.Register)
r.POST("/login", handlers.Login)


auth := r.Group("/")
auth.Use(middleware.AuthMiddleware())

auth.POST("/books", handlers.CreateBook)
auth.DELETE("/books/:id", handlers.DeleteBook)


r.GET("/books", handlers.GetAllBooks)
r.GET("/books/:id", handlers.GetBookByID)

r.Run(":8080")

}