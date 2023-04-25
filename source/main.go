package main

import (
	"gweb"
	"net/http"
)

func main() {
	r := gweb.New()
	r.GET("/", func(c *gweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gweb.Context) {
		// expect /hello?name=tzx
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gweb.Context) {
		// expect /hello/tzx
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gweb.Context) {
		c.JSON(http.StatusOK, gweb.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
