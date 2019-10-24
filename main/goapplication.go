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
	router.GET("/", todo.GetTodos)
	router.POST("/todo", todo.AddTodo)
	router.DELETE("/todo", todo.DeleteTodo)
	router.PUT("/todo", todo.UpdateTodo)

}
