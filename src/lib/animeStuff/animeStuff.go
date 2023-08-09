package animeStuff

import (
	"botwebo2/lib/embed"
	"botwebo2/lib/myRequests"
	"encoding/json"
	"fmt"
	"time"
)

const (
	graphql = "https://graphql.anilist.co"

	query = `query($name: String) {
        Media (search: $name, type: ANIME) {
            episodes
            status(version: 2)
            season
            seasonYear
            nextAiringEpisode{
                timeUntilAiring
                episode
            }
            title {
                romaji
            }
        }
    }`
)

// struct of http request json
type Info struct {
	Data Data `json:"data"`
}

type Data struct {
	Media Media `json:"media"`
}

type Media struct {
	Episodes          int         `json:"episodes"`
	Status            string      `json:"status"`
	Season            string      `json:"season"`
	SeasonYear        int         `json:"seasonYear"`
	NextAiringEpisode NextEpisode `json:"nextAiringEpisode"`
	Title             Title       `json:"title"`
}

type NextEpisode struct {
	TimeUntilAiring int `json:"timeUntilAiring"`
	Episode         int `json:"episode"`
}

type Title struct {
	Romaji string `json:"romaji"`
}

func TimeUntilAiring(anime string) (embed.Embed, error) {
	//creates embed
	message := embed.NewEmbed()

	//creates variables
	temp := map[string]interface{}{
		"name": anime,
	}

	variables, err := json.Marshal(temp)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error converting map to json"), err
	}

	//request to anilist
	resp, err := myRequests.PostQuery(graphql, query, string(variables))
	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error http request"), err
	}

	//convert string to json struct
	info := Info{}

	err = json.Unmarshal([]byte(resp), &info)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription(resp), err
	}

	//set title to embed
	message.SetTitle(info.Data.Media.Title.Romaji)

	//change description depends on anime status
	switch info.Data.Media.Status {
	case "FINISHED":
		message.SetDescription(fmt.Sprintf("Show has **already ended.** (%d episodes)", info.Data.Media.Episodes))

	case "RELEASING":
		date := time.Now().Local().Add(time.Second * time.Duration(info.Data.Media.NextAiringEpisode.TimeUntilAiring))
		message.SetDescription(fmt.Sprintf("Episode **%d** airs in **%d days (%s**)",
			info.Data.Media.NextAiringEpisode.Episode,
			info.Data.Media.NextAiringEpisode.TimeUntilAiring/86400,
			fmt.Sprintf("%s, %04d-%02d-%02d, %02d:%02d", date.Weekday(), date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute())))

	case "NOT_YET_RELEASE":
		description := ""
		if info.Data.Media.SeasonYear != 0 {
			description += "Airing in **" + fmt.Sprintf("%d", info.Data.Media.SeasonYear)
			if info.Data.Media.Season != "" {
				description += ", " + info.Data.Media.Season
			}
			description += "**"
		} else {
			description = "Unknown release date"
		}
		message.SetDescription(description)

	}

	//set color to embed
	message.SetColor(embed.SuccessColor)
	return *message, err
}
