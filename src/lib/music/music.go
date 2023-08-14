package music

import (
	"botwebo2/lib/embed"

	"github.com/bwmarrin/discordgo"
)

var (
	PlayerList map[string]PlayerFunctions
)

func init() {
	PlayerList = map[string]PlayerFunctions{}
}

func Play(option string, guildID string, channel string, dg *discordgo.Session) (embed.Embed, error) {
	//get the player of guild
	guild, ok := PlayerList[guildID]
	//switch the option
	switch option {
	case "stop":
		if ok {
			return guild.stop()
		}
	case "song":
		if ok {
			return guild.song()
		}
	case "skip":
		if ok {
			return guild.skip()
		}
	case "list":
		if ok {
			return guild.list()
		}
	case "loop":
		if ok {
			return guild.loop()
		}
	case "empty":
		if ok {
			return guild.emptyList()
		}
	case "shuffle":
		if ok {
			return guild.shuffle()
		}
	default:
		if ok {
			return guild.addSongs(option, channel, dg)
		} else {
			//si no existe crear
			PlayerList[guildID] = &Player{guildID: guildID}
			return PlayerList[guildID].addSongs(option, channel, dg)
		}
	}

	return *embed.NewEmbed().SetColor(embed.ErrorColor).SetDescription("you cannot do that cause i am not in a voiceChannel"), nil
}

func Rplay(option string, guildID string) (embed.Embed, error) {
	guild, ok := PlayerList[guildID]

	switch option {
	case "stop":
		if ok {
			return guild.stop()
		}
	case "song":
		if ok {
			return guild.song()
		}
	case "skip":
		if ok {
			return guild.skip()
		}
	default:
		if ok {
			return guild.random(option)
		} else {
			PlayerList[guildID] = &Player{guildID: guildID}
			return PlayerList[guildID].random(option)
		}
	}

	return *embed.NewEmbed().SetColor(embed.ErrorColor).SetDescription("you cannot do that cause i am not in a voiceChannel"), nil
}
