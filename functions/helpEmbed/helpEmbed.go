package helpEmbed

import "github.com/bwmarrin/discordgo"

func GetEmbed() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x00ff00, // Green
		Title:       "Go-DBOT commands help ☑️",
		Description: "Check all those glorious commands below",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "$avatar @optionalMention",
				Value:  "Retrive user avatar, leave @optionalMention empty to get your avatar",
				Inline: false,
			},
			{
				Name:   "$cat",
				Value:  "Random cat image 🐱",
				Inline: false,
			},
			{
				Name:   "$dog",
				Value:  "Random dog image 🐕",
				Inline: false,
			},
		},
	}

	return embed
}
