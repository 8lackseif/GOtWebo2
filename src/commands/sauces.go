package commands

import (
	"botwebo2/lib/sauces"
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	//declare command
	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        "sauce",
		Description: "/sauce [image_url]",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "string-option",
				Description: "image_url",
				Required:    true,
			},
		},
	})
	//declare command function
	CommandHandlers["sauce"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}
		option := optionMap["string-option"]

		message, err := sauces.GetSauce(option.StringValue())

		if err != nil {
			log.Fatalln(err)
		}

		s.ChannelMessageSendEmbed(i.ChannelID, message.MessageEmbed)
	}
}
