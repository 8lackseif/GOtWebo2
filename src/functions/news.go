package functions

import (
	"botwebo2/lib/animeNews"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	ANIME_CHANNEL = "anime-webonews"
	MANGA_CHANNEL = "manga-webonews"
	INTERVAL      = 20
)

func AnimeNews(dg *discordgo.Session) {
	for true {

		//get channels which have to send news
		var mangaChannels, animeChannels []string

		for _, guild := range dg.State.Guilds {
			for _, channel := range guild.Channels {

				if channel.Name == ANIME_CHANNEL {
					animeChannels = append(animeChannels, channel.ID)
				}

				if channel.Name == MANGA_CHANNEL {
					mangaChannels = append(mangaChannels, channel.ID)
				}
			}

		}

		//get the news
		animeNews, mangaNews, err := animeNews.CheckNews()

		if err != nil {
			println(err.Error())
		}

		if animeNews != nil && mangaNews != nil {
			//send news
			for _, new := range animeNews {
				for _, channelID := range animeChannels {
					dg.ChannelMessageSend(channelID, new)
				}
			}

			for _, new := range mangaNews {
				for _, channelID := range mangaChannels {
					dg.ChannelMessageSend(channelID, new)
				}
			}
		}

		time.Sleep(INTERVAL * time.Minute)
	}
}
