package routes

import (
	"net/http"
	"reformata-cms-core/configs"
	"reformata-cms-core/utilities"
	"strconv"

	"github.com/labstack/echo/v5"
)

// var Config *configs.SiteConfig = configs.Get("config.yaml")

type InfoData struct {
	Page  string
	Title string
}

type HomeData struct {
	Page    string
	Title   string
	VideoId string
}

func BaseRoutes(e *echo.Echo) {

	e.GET("/", home)
	e.GET("/missions", missions)
	e.GET("/about", about)

}



func home(c *echo.Context) error {
	// println(configs.Config.Secrets.YouTubeAPIKey)
	// println(configs.Config.ConfigFile.YouTubeApiUrl)
	// println(configs.Config.ConfigFile.Name)

	// routeCaches := cache.New[string, string]()

	recent_video := utilities.FetchRecentVideo()

	d := HomeData{Page: "home", Title: configs.Config.ConfigFile.Name, VideoId: recent_video.Id.VideoId}

	hx, err := strconv.ParseBool(c.Request().Header.Get("Hx-Request"))

	if err == nil && hx {
		return c.Render(http.StatusOK, "home", d)
	}

	return c.Render(http.StatusOK, "base", d)
}

func missions(c *echo.Context) error {
	return nil
}

func about(c *echo.Context) error {
	data := InfoData{Page: "about", Title: "About Us"}
	hx, err := strconv.ParseBool(c.Request().Header.Get("Hx-Request"))

	if err == nil && hx {
		return c.Render(http.StatusOK, "about", data)
	}

	return c.Render(http.StatusOK, "base", data)
}

// func beliefs_routes(e *echo.Echo) {
// 	e.GET("/beliefs/",)
// }
