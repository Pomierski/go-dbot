package commands

import (
	"go-dbot/functions/helpEmbed"
	"go-dbot/functions/randomCat"
	"go-dbot/functions/randomDog"
	. "go-dbot/functions/sendMessage"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleCommands(prefix string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content[0:1] != prefix {
		return
	}

	messageSlice := strings.Split(m.Content[1:], " ")
	command := messageSlice[0]
	//params := messageSlice[1:]

	switch command {
	case "avatar":
		{
			if len(m.Mentions) != 0 {
				avatarUrl := m.Mentions[0].AvatarURL("")
				SendMessage(avatarUrl, s, m)
			} else {
				SendMessage(m.Author.AvatarURL(""), s, m)
			}
		}
	case "cat":
		{
			catImageUrl := randomCat.Get()
			if len(catImageUrl) != 0 {
				SendMessage(catImageUrl, s, m)
			} else {
				SendMessage("Error occured while fetching your cat image, try again later 😿", s, m)
			}
		}
	case "dog":
		{
			dogImageUrl := randomDog.Get()
			if len(dogImageUrl) != 0 {
				SendMessage(dogImageUrl, s, m)
			} else {
				SendMessage("Error occured while fetching your dog image, try again later 🐶", s, m)
			}
		}
	case "help":
		{
			s.ChannelMessageSendEmbed(m.ChannelID, helpEmbed.GetEmbed())
		}
	}
}
