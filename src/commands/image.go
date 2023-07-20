package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {

	Commands = append(Commands, &discordgo.ApplicationCommand{
		Name:        "image",
		Description: "send image",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "string-option",
				Description: "String option",
				Required:    true,
			},
		},
	})

	CommandHandlers["image"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Access options in the order provided by the user.
		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		images := map[string]string{
			"no":        "https://cdn.discordapp.com/attachments/734750766895595581/843267512022859816/no.gif",
			"yes":       "https://cdn.discordapp.com/attachments/734750766895595581/843267625975021568/yes.gif",
			"haachama":  "https://cdn.discordapp.com/attachments/734750766895595581/843267894686908442/haachama.jpg",
			"pekora":    "https://cdn.discordapp.com/attachments/734750766895595581/843268060445016105/pekora.jpg",
			"smug":      "https://cdn.discordapp.com/attachments/734750766895595581/843268167089258517/smug.jpg",
			"pray":      "https://cdn.discordapp.com/attachments/649025469219340288/853772957028319292/unknown.png",
			"please":    "https://i1.sndcdn.com/avatars-Izsdy6YmsiXZk1Sr-8AXfwA-t500x500.jpg",
			"trembling": "https://cdn.discordapp.com/attachments/709788450408366162/972992631456030830/the-quintessential-quintuplets-itsuki.gif",
		}

		option := optionMap["string-option"]
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: images[option.StringValue()],
			},
		})
	}
}
