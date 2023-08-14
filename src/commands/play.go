package commands

import (
	"botwebo2/lib/embed"
	"botwebo2/lib/music"

	"github.com/bwmarrin/discordgo"
)

func init() {
	//declare command
	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        "play",
		Description: "/play [youtube_link|stop|song|skip|list|loop|empty]",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "option",
				Description: "[youtube_link|stop|song|skip|list|loop|empty]",
				Required:    true,
			},
		},
	})
	//declare command function
	CommandHandlers["play"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		channel := ""
		for _, g := range s.State.Guilds {
			for _, v := range g.VoiceStates {
				if v.UserID == i.Member.User.ID {
					channel = v.ChannelID
				}
			}
		}

		if channel == "" {
			s.ChannelMessageSendEmbed(i.ChannelID, *&embed.NewEmbed().SetColor(embed.ErrorColor).SetDescription("you have to join a channel.").MessageEmbed)
		}

		option := optionMap["option"]

		embed, err := music.Play(option.StringValue(), i.GuildID, channel, s)

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
