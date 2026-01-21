package main

import (
	"html/template"
	"io"
	"reformata-cms-core/configs"
	"reformata-cms-core/routes"
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(c *echo.Context, w io.Writer, name string, data any) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	site_config := configs.Get("config.yaml")

	println("Starting Site For: ", site_config.Name)

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	routes.BaseRoutes(e)

	if err := e.Start(fmt.Sprintf(":%v",site_config.Server.Port)); err != nil {
		e.Logger.Error("Failed to start server", "error", err)
	}
}
