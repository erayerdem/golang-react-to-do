package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"todo"
)

var router = gin.Default()

func main() {
	router.Use(cors.Default())
	defer router.Run()
	router.GET("/todo", todo.GetTodos)
	router.POST("/todo", todo.AddTodo)
	router.POST("/deletetodo", todo.DeleteTodo)
	router.PUT("/todo", todo.UpdateTodo)

}
