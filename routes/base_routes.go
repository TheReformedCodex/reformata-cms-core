package routes

import (
	"fmt"
	"net/http"
	"reformata-cms-core/configs"
	"reformata-cms-core/utilities"

	"github.com/labstack/echo/v5"
)

// var Config *configs.SiteConfig = configs.Get("config.yaml")

type BaseData struct {
	Page    string
	Title   string
	VideoId string
}

func BaseRoutes(e *echo.Echo) {

	e.GET("/", home)
	e.GET("/missions", missions)

}

func home(c *echo.Context) error {
	// println(configs.Config.Secrets.YouTubeAPIKey)
	// println(configs.Config.ConfigFile.YouTubeApiUrl)
	// println(configs.Config.ConfigFile.Name)
	recent_video := utilities.FetchRecentVideo()
	fmt.Println(recent_video.Id.VideoId)
	d := BaseData{Page: "home", Title: configs.Config.ConfigFile.Name, VideoId: recent_video.Id.VideoId}
	return c.Render(http.StatusOK, "base", d)
}

func missions(c *echo.Context) error {
	return nil
}

// func beliefs_routes(e *echo.Echo) {
// 	e.GET("/beliefs/",)
// }
