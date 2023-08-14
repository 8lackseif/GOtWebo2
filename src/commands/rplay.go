package commands

import (
	"botwebo2/lib/music"

	"github.com/bwmarrin/discordgo"
)

func init() {
	//declare command
	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        "rplay",
		Description: "/rplay [anilist_user|stop|song|skip]",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "option",
				Description: "[anilist_user|stop|song|skip]",
				Required:    true,
			},
		},
	})
	//declare command function
	CommandHandlers["rplay"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		option := optionMap["option"]

		embed, err := music.Rplay(option.StringValue(), i.GuildID)

		if err != nil {
			println(err)
		}

		s.ChannelMessageSendEmbed(i.ChannelID, embed.MessageEmbed)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "",
			},
		})
	}

}
