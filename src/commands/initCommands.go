package commands

import "github.com/bwmarrin/discordgo"

//initialize
var Commands []*discordgo.ApplicationCommand

var CommandHandlers = make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate))
