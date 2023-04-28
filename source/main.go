package main

import (
	"fmt"
	"gweb"
	"html/template"
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
	r := gweb.New()
	r.Use(gweb.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Brown", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gweb.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gweb.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gweb.H{
			"title":  "gweb",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gweb.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gweb.H{
			"title": "gweb",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
