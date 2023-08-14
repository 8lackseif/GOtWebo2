package commands

import (
	"botwebo2/lib/danbooru"

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
				Name:        "danbooru_tag",
				Description: "danbooru_tag",
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

		option := optionMap["danbooru_tag"]

		//check if option exists
		var tag string

		if option == nil {
			tag = "*"
		} else {
			tag = option.StringValue()
		}

		//get random image
		embed, err := danbooru.SendDanbooruImage(tag)

		//variables

		if err != nil {
			println(err)
		}

		//get similar tags if cannot find tag matches
		if embed.Image.URL == "" {

			embed, err = danbooru.GetSimilarTags(option.StringValue())

			if err != nil {
				println(err)
			}
		}

		//return embed to user
		s.ChannelMessageSendEmbed(i.ChannelID, embed.MessageEmbed)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "",
			},
		})
	}
}
