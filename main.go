package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-dbot/commands"
	"go-dbot/config"

	"github.com/bwmarrin/discordgo"
)

var prefix string

func main() {
	var token string

	config.Apply(&token, &prefix)

	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Bot successfully started!")

	discord.AddHandler(messageCreate)
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()

	if err != nil {
		log.Fatal(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	commands.HandleCommands(prefix, s, m)
}
