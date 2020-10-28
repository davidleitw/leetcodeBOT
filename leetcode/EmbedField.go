package leetcode

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
)

const color int = 0xFFD700

var author *discordgo.MessageEmbedAuthor = &discordgo.MessageEmbedAuthor{
	URL:     "https://github.com/davidleitw/leetcodeBOT",
	Name:    "薯條醬",
	IconURL: "https://www.flaticon.com/svg/static/icons/svg/25/25231.svg",
}

var footer *discordgo.MessageEmbedFooter = &discordgo.MessageEmbedFooter{
	Text:    "一起刷leetcode 一起進步 (･ω´･ )",
	IconURL: "https://cdn.icon-icons.com/icons2/2389/PNG/512/leetcode_logo_icon_145113.png",
}

func difficulty(d int) string {
	switch d {
	case 1:
		return "eazy"
	case 2:
		return "medium"
	case 3:
		return "hard"
	default:
		return "nil"
	}
}

func baseEmbedMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:  color,
		Author: author,
	}
}

func ProblemEmbedMessage(problems []*model.Problem) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	fields := []*discordgo.MessageEmbedField{}

	for _, problem := range problems {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "----------",
			Value:  fmt.Sprintf("No.%d\n", problem.ProblemID),
			Inline: false,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Title",
			Value:  problem.ProblemTitle,
			Inline: false,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "URL",
			Value:  problem.ProblemURL,
			Inline: false,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Difficulty",
			Value:  difficulty(problem.Difficulty),
			Inline: false,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "----------",
			Value:  "          ",
			Inline: false,
		})
	}

	msg.Fields = fields
	msg.Footer = footer
	return msg
}
