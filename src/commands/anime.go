package commands

import (
	"botwebo2/lib/animeStuff"

	"github.com/bwmarrin/discordgo"
)

const (
	ANIME_CHANNEL = "anime-webonews"
	MANGA_CHANNEL = "manga-webonews"
)

func init() {
	//declare command
	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        "anime",
		Description: "/anime [anime_name]",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "anime_name",
				Description: "anime_name",
				Required:    true,
			},
		},
	})
	//declare command function
	CommandHandlers["anime"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		option := optionMap["anime_name"]

		message, err := animeStuff.TimeUntilAiring(option.StringValue())
		if err != nil {
			println(err)
		}

		s.ChannelMessageSendEmbed(i.ChannelID, message.MessageEmbed)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "",
			},
		})
	}

}
