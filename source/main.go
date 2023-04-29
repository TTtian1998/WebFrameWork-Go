package main

import (
	"fmt"
	"gweb"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	//use Default to demonstrate customize middleware and print trace message
	//net/http also contains their panic handler, so when there is a panic when using gweb.New
	//the app won't terminate
	r := gweb.Default()
	r.GET("/", func(c *gweb.Context) {
		c.String(http.StatusOK, "Hello GWeb\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gweb.Context) {
		names := []string{"gweb"}
		c.String(http.StatusOK, names[100])
	})
	r.GET("/panic", func(c *gweb.Context) {
		names := []string{"gweb"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
