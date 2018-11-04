package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Template struct {
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if t, ok := templates[name]; ok {
		return t.ExecuteTemplate(w, "layout.html", data)
	}
	c.Echo().Logger.Debugf("Template[%s] Not Found.", name)
	return templates["error"].ExecuteTemplate(w, "layout.html", "Internal Server Error")
}

func loadTemplates()  {
	var baseTemplate = "templates/layout.html"
	templates = make(map[string]*template.Template)

	templates["login"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/login.html"))
}