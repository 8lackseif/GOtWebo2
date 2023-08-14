package commands

// import (
// 	"github.com/bwmarrin/discordgo"
// )

// func init() {
// 	//declare command
// 	Commands = append(Commands, &discordgo.ApplicationCommand{
// 		Name:        "pekofy",
// 		Description: "/pekofy text",
// 	})
// 	//declare command function
// 	CommandHandlers["pekofy"] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 		// Access options in the order provided by the user.

// 		s.ChannelMessageSendReply(i.ChannelID, option.StringValue()+" peko", &discordgo.MessageReference{
// 			MessageID: i.Message.ID,
// 			ChannelID: i.Message.ChannelID,
// 			GuildID:   i.Message.GuildID,
// 		})
// 	}
// }
