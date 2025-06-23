package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()
	defer DB.Close()

	name := "World"
	fmt.Printf("Hello, %s\n", name)

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")

	// pointer using for better performance
	engine.GET("/", func(c *gin.Context) {
		todos := ReadToDoList()

		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	engine.POST("/todos", func(c *gin.Context) {
		title := c.PostForm("title")
		status := c.PostForm("status")
		id, _ := CreateToDo(title, status)

		c.HTML(http.StatusOK, "task.html", gin.H{
			"Title":  title,
			"Status": status,
			"Id":     id,
		})
	})

	engine.DELETE("/todos/:id", func(c *gin.Context) {
		param := c.Param("id")
		id, _ := strconv.ParseInt(param, 10, 64)
		DeleteToDo(id)
	})

	engine.Run("127.0.0.1:8889")
}
