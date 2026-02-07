package routes

import (
	"net/http"
	"reformata-cms-core/configs"

	"github.com/labstack/echo/v5"
)

var Config *configs.SiteConfig = configs.Get("config.yaml")

type BaseData struct {
	Page  string
	Title string
	Body  string
}

func BaseRoutes(e *echo.Echo) {

	e.GET("/", home)

}

func home(c *echo.Context) error {
	d := BaseData{Page: "home", Title: Config.Name, Body: "This is purely an experiment"}
	return c.Render(http.StatusOK, "base", d)
}

// func beliefs_routes(e *echo.Echo) {
// 	e.GET("/beliefs/",)
// }
