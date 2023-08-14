package main

import (
	"botwebo2/commands"
	"botwebo2/functions"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	//bot and commands

	//Load tokens
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
		return
	}

	//ADDING HANDLERS
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			go h(s, i)
		}
	})

	//OPEN BOT
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// REGISTER COMMANDS
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.Commands))
	//create chan for goroutines
	ch := make(chan *discordgo.ApplicationCommand)
	//create commands with goroutines
	for _, v := range commands.Commands {
		go createCommand(ch, dg, v)
	}
	//register commands received from chan
	go func() {
		i := 0
		for cmd := range ch {
			registeredCommands[i] = cmd
			i++
		}
	}()

	//functions
	go functions.AnimeNews(dg)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func createCommand(ch chan *discordgo.ApplicationCommand, dg *discordgo.Session, cmd *discordgo.ApplicationCommand) {
	cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", cmd)
	if err != nil {
		log.Fatalln("Cannot create '%s' command: %s\n" + cmd.Name + err.Error())
	}
	ch <- cmd
}
