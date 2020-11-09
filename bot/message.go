package bot

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
	uuid "github.com/satori/go.uuid"
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

func ListReportsMessage(reports []model.ReportsResult, nick string) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	msg.Description = fmt.Sprintf("%s預計要報告的題目如下:\n", nick)
	fields := []*discordgo.MessageEmbedField{}
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Problem",
		Value:  "----------",
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Title: ",
		Value:  "----------------------------",
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Difficulty: ",
		Value:  "----------",
		Inline: true,
	})
	// "\u200b"
	for _, report := range reports {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("No.%d\n", report.ProblemID),
			Value:  "\u200b",
			Inline: true,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   report.Title,
			Value:  "\u200b",
			Inline: true,
		})
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   difficulty(report.Difficulty),
			Value:  "\u200b",
			Inline: true,
		})
	}
	msg.Fields = fields
	return msg
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

func SetStudyGroupStartTimeMessage(SID uuid.UUID, exist bool) *discordgo.MessageEmbed {
	sg := model.GetStudyGroup(SID)
	msg := baseEmbedMessage()
	fields := []*discordgo.MessageEmbedField{}

	if exist {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "成功修改了下次讀書會的開始時間:",
			Value:  "--------------------",
			Inline: false,
		})
	} else {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "成功創建了一次新的讀書會，基本資訊如下:",
			Value:  "--------------------",
			Inline: false,
		})
	}

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "預計開始時間",
		Value:  time.Time(sg.StartTime).Format("2006/01/02 15:04"),
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "目前參與人數",
		Value:  fmt.Sprintf("%d人", sg.Attendance),
		Inline: true,
	})

	msg.Fields = fields
	return msg
}

func StudyGroupInfoMessage(sg *model.StudyGroup, cnt int64) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	fields := []*discordgo.MessageEmbedField{}

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "下一次leetcode讀書會的資訊如下:",
		Value:  "---------------------------",
		Inline: false,
	})

	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "預計開始時間",
		Value:  time.Time(sg.StartTime).Format("2006/01/02 15:04"),
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "目前參與人數",
		Value:  fmt.Sprintf("%d人", sg.Attendance),
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "報告題目數",
		Value:  fmt.Sprintf("%d題", cnt),
		Inline: true,
	})

	msg.Fields = fields
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
