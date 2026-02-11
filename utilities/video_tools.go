package utilities

import (
	"reformata-cms-core/configs"
	"regexp"
	"time"

	"github.com/imroc/req/v3"
)

type VideoSearchId struct {
	Kind    string `json:"kind"`
	VideoId string `json:"videoId"`
}

type VideoSearchSnippet struct {
	PublishTime time.Time `json:"publishTime"`
	Title       string    `json:"title"`
}

type VideoSearchResult struct {
	Kind string             `json:"kind"`
	Id   VideoSearchId      `json:"id"`
	Info VideoSearchSnippet `json:"snippet"`
}

type VideoSearchSnippetResponse struct {
	Items []VideoSearchResult `json:"items"`
}

func FetchRecentVideo() VideoSearchResult {
	client := req.C()

	api_url := configs.Config.ConfigFile.YouTubeApiUrl
	channel_id := configs.Config.ConfigFile.YouTubeChannelId
	api_key := configs.Config.Secrets.YouTubeAPIKey

	var response VideoSearchSnippetResponse

	resp, err := client.NewRequest().
		SetQueryParam("channelId", channel_id).
		SetQueryParam("key", api_key).
		SetQueryParam("order", "date").
		SetQueryParam("part", "snippet").
		SetSuccessResult(&response).
		Get(api_url)

	if err != nil {
		println("Unable to query the YouTube API", err)
	}

	if resp.Err != nil {
		println("Unable to query the YouTube API", resp.Err)
	}

	for index, value := range response.Items {
		match, _ := regexp.MatchString("Sunday Service", value.Info.Title)

		if match {
			return response.Items[index]
		}
	}
	return response.Items[0]
}
