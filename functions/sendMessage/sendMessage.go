package sendMessage

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SendMessage(content string, s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		fmt.Println(err)
	}
}
