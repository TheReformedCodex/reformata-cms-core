package main

import (
	"fmt"
	"html/template"
	"io"
	"reformata-cms-core/configs"
	"reformata-cms-core/routes"

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
	site_config := configs.GetConfig("config.yaml")

	println("Starting Site For: ", site_config.ConfigFile.Name)

	e := echo.New()
	e.Static("/static", "static")
	e.Use(middleware.RequestLogger())

	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/**/*.html")),
	}

	routes.BaseRoutes(e)

	if err := e.Start(fmt.Sprintf(":%v", site_config.ConfigFile.Server.Port)); err != nil {
		e.Logger.Error("Failed to start server", "error", err)
	}
}
