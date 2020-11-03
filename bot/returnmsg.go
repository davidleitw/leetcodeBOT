package bot

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
)

const color int = 0xFFD700

var author *discordgo.MessageEmbedAuthor = &discordgo.MessageEmbedAuthor{
	URL:     "https://github.com/davidleitw/leetcodeBOT",
	Name:    "薯條醬",
	IconURL: "https://img.icons8.com/doodle/2x/github.png",
}

var footer *discordgo.MessageEmbedFooter = &discordgo.MessageEmbedFooter{
	Text:    "  一起刷leetcode 一起進步 (･ω´･ )",
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
		return "Difficulty error"
	}
}

func userEmbedMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:  color,
		Footer: footer,
	}
}

func UserProblemsEmbedMessage(UserID string, problems []*model.Problem) *discordgo.MessageEmbed {
	msg := userEmbedMessage()

	return msg
}

func baseEmbedMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:  color,
		Author: author,
		Footer: footer,
	}
}

func ProblemsMsg(problems []*model.Problem) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	fields := []*discordgo.MessageEmbedField{}

	for _, problem := range problems {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Problem",
			Value:  fmt.Sprintf("No.%d\n", problem.ProblemID),
			Inline: false,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Title: ",
			Value:  problem.ProblemTitle,
			Inline: true,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "URL: ",
			Value:  problem.ProblemURL,
			Inline: true,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Difficulty: ",
			Value:  difficulty(problem.Difficulty),
			Inline: true,
		})
	}

	msg.Fields = fields
	msg.Timestamp = time.Now().Format("2006-0102 15:04")
	return msg
}

func ProblemsDetailMsg(problems []*model.Problem) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	fields := []*discordgo.MessageEmbedField{}

	for _, problem := range problems {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Problem",
			Value:  fmt.Sprintf("No.%d\n", problem.ProblemID),
			Inline: false,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Title: ",
			Value:  problem.ProblemTitle,
			Inline: true,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "URL: ",
			Value:  problem.ProblemURL,
			Inline: true,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Difficulty: ",
			Value:  difficulty(problem.Difficulty),
			Inline: true,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Ac",
			Value:  fmt.Sprintf("%d\n", problem.Accept),
			Inline: true,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Submit",
			Value:  fmt.Sprintf("%d\n", problem.Submit),
			Inline: true,
		})
	}

	msg.Fields = fields
	msg.Timestamp = time.Now().Format("2006-0102 15:04")
	return msg
}
