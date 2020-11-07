package bot

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
)

// return message template function in this file.

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

func baseEmbedMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:  color,
		Author: author,
		Footer: footer,
	}
}

func AddReportMessage(problems []*model.Problem, nike string, month, day int) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	fields := []*discordgo.MessageEmbedField{}
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   fmt.Sprintf("%s成功添加了以下題目:", nike),
		Value:  "--------------------",
		Inline: false,
	})

	for _, problem := range problems {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Problem",
			Value:  fmt.Sprintf("No.%d\n", problem.ProblemID),
			Inline: true,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Title: ",
			Value:  problem.ProblemTitle,
			Inline: true,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Difficulty: ",
			Value:  difficulty(problem.Difficulty),
			Inline: true,
		})
	}
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "--------------------",
		Value:  fmt.Sprintf("讀書會舉辦日期: %d月%d日", month, day),
		Inline: false,
	})
	msg.Fields = fields
	return msg
}

func SearchProblemsMsg(problems []*model.Problem) *discordgo.MessageEmbed {
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

func SearchProblemsDetailMsg(problems []*model.Problem) *discordgo.MessageEmbed {
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
