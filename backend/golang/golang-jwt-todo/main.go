package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Text string
	Done bool
}

var todos []Todo
var loggedInUser string

func main() {
	// Initialize the server
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Todos":        todos,
			"LoggedInUser": loggedInUser != "",
			"Username":     loggedInUser,
		})
	})

	router.POST("/add", func(c *gin.Context) {
		text := c.PostForm("todo")
		todo := Todo{Text: text, Done: false}
		todos = append(todos, todo)
		c.Redirect(http.StatusSeeOther, "/")
	})
}
