package commands

import (
	"botwebo2/lib/animeStuff"

	"github.com/bwmarrin/discordgo"
)

func init() {
	//declare command
	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        "anime",
		Description: "/anime [anime_name]",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "string-option",
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
		option := optionMap["string-option"]

		message, err := animeStuff.TimeUntilAiring(option.StringValue())
		if err != nil {
			println(err)
		}

		s.ChannelMessageSendEmbed(i.ChannelID, message.MessageEmbed)
	}
}
