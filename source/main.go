package main

import (
	"gweb"
	"net/http"
)

func main() {
	r := gweb.New()
	r.GET("/", func(c *gweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})
	r.GET("/hello", func(c *gweb.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gweb.Context) {
		c.JSON(http.StatusOK, gweb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
