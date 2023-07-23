package commands

import (
	"botwebo2/lib/danbooru"
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	//declare command
	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        "danbooru",
		Description: "/danbooru | /danboru [something]",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "string-option",
				Description: "String option",
				Required:    false,
			},
		},
	})

	//declare command functionality
	CommandHandlers["danbooru"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		option := optionMap["string-option"]

		//check if option exists
		var tag string

		if option == nil {
			tag = "*"
		} else {
			tag = option.StringValue()
		}

		//get random image
		randomImage, err := danbooru.SendDanbooruImage(tag)

		if err != nil {
			log.Fatalln(err)
		}

		//get similar tags if cannot find tag matches
		if randomImage == "" {
			randomImage, err = danbooru.GetSimilarTags(option.StringValue())

			if err != nil {
				log.Fatalln(err)
			}
		}

		//return message to user
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: randomImage,
			},
		})
	}
}
