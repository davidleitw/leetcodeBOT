package bot

import "github.com/bwmarrin/discordgo"

func HelpMsg() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:  color,
		Author: author,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "指令使用方法",
				Value: "https://hackmd.io/5nBbJrxrTrGTO1WOraXtaQ",
			},
		},
	}
}
