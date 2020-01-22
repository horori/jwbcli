package jwapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// LatestVideos struct {
type LatestVideos struct {
	Category struct {
		Key         string   `json:"key"`
		Type        string   `json:"type"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
		Images      struct {
		} `json:"images"`
		ParentCategory interface{} `json:"parentCategory"`
		Media          []struct {
			GUID                       string    `json:"guid"`
			LanguageAgnosticNaturalKey string    `json:"languageAgnosticNaturalKey"`
			NaturalKey                 string    `json:"naturalKey"`
			Type                       string    `json:"type"`
			PrimaryCategory            string    `json:"primaryCategory"`
			Title                      string    `json:"title"`
			Description                string    `json:"description"`
			FirstPublished             time.Time `json:"firstPublished"`
			Duration                   float64   `json:"duration"`
			DurationFormattedHHMM      string    `json:"durationFormattedHHMM"`
			DurationFormattedMinSec    string    `json:"durationFormattedMinSec"`
			Tags                       []string  `json:"tags"`
			Files                      []struct {
				ProgressiveDownloadURL string    `json:"progressiveDownloadURL"`
				FlashStreamingURL      string    `json:"flashStreamingURL"`
				Checksum               string    `json:"checksum"`
				Filesize               int       `json:"filesize"`
				ModifiedDatetime       time.Time `json:"modifiedDatetime"`
				BitRate                float64   `json:"bitRate"`
				Duration               float64   `json:"duration"`
				FrameHeight            int       `json:"frameHeight"`
				FrameWidth             int       `json:"frameWidth"`
				Label                  string    `json:"label"`
				FrameRate              float64   `json:"frameRate"`
				Mimetype               string    `json:"mimetype"`
				Subtitled              bool      `json:"subtitled"`
				Subtitles              struct {
					URL              string    `json:"url"`
					ModifiedDatetime time.Time `json:"modifiedDatetime"`
					Checksum         string    `json:"checksum"`
				} `json:"subtitles"`
			} `json:"files"`
			Images struct {
				Pss struct {
					Sm string `json:"sm"`
					Lg string `json:"lg"`
					Xs string `json:"xs"`
					Md string `json:"md"`
				} `json:"pss"`
				Sqs struct {
					Sm string `json:"sm"`
					Lg string `json:"lg"`
					Xs string `json:"xs"`
					Md string `json:"md"`
				} `json:"sqs"`
				Pns struct {
					Md string `json:"md"`
					Lg string `json:"lg"`
					Xs string `json:"xs"`
				} `json:"pns"`
				Rps struct {
					Xl string `json:"xl"`
				} `json:"rps"`
				Sqr struct {
					Sm string `json:"sm"`
					Lg string `json:"lg"`
					Xs string `json:"xs"`
					Xl string `json:"xl"`
					Md string `json:"md"`
				} `json:"sqr"`
				Pnr struct {
					Sm string `json:"sm"`
					Lg string `json:"lg"`
					Xs string `json:"xs"`
					Md string `json:"md"`
				} `json:"pnr"`
				Wss struct {
					Sm string `json:"sm"`
					Lg string `json:"lg"`
					Xs string `json:"xs"`
				} `json:"wss"`
				Wsr struct {
					Sm string `json:"sm"`
					Xs string `json:"xs"`
					Lg string `json:"lg"`
				} `json:"wsr"`
				Lsr struct {
					Sm string `json:"sm"`
					Lg string `json:"lg"`
					Xs string `json:"xs"`
					Xl string `json:"xl"`
					Md string `json:"md"`
				} `json:"lsr"`
				Rph struct {
					Xl string `json:"xl"`
				} `json:"rph"`
				Lss struct {
					Sm string `json:"sm"`
					Lg string `json:"lg"`
					Xs string `json:"xs"`
					Xl string `json:"xl"`
					Md string `json:"md"`
				} `json:"lss"`
				Psr struct {
					Sm string `json:"sm"`
					Xs string `json:"xs"`
					Lg string `json:"lg"`
					Md string `json:"md"`
				} `json:"psr"`
			} `json:"images"`
			AvailableLanguages []string `json:"availableLanguages"`
		} `json:"media"`
	} `json:"category"`
	Pagination struct {
		TotalCount int `json:"totalCount"`
		Offset     int `json:"offset"`
		Limit      int `json:"limit"`
	} `json:"pagination"`
}

// ParseLatestVideo Return LatestVideo object
func ParseLatestVideo(language string) (data *LatestVideos, err error) {

	httpResponse, err := http.Get("https://data.jw-api.org/mediator/v1/categories/" + language + "/LatestVideos?detailed=1&clientType=tvjworg")
	if err != nil {
		return
	}

	jsonString, err := ioutil.ReadAll(httpResponse.Body)
	if httpResponse != nil {
		defer httpResponse.Body.Close()
	}

	data = new(LatestVideos)
	if err := json.Unmarshal(jsonString, data); err != nil {
		log.Fatal(err)
	}
	return
}
