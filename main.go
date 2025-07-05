package main

import (
	"fmt"
	"go-htmx-tutor/stringutils"
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

	engine.GET("/foo", func(c *gin.Context) {
		greet_result := stringutils.Greet("emir")

		c.JSON(http.StatusOK, gin.H{
			"message": greet_result,
		})
	})

	server_host := "127.0.0.1:8889"
	fmt.Printf("Server host: http://%s\n", server_host)

	engine.Run(server_host)
}
