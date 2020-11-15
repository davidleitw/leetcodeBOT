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
		return "Eazy"
	case 2:
		return "Medium"
	case 3:
		return "Hard"
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

type fields []*discordgo.MessageEmbedField

func baseField() fields {
	return []*discordgo.MessageEmbedField{}
}

func (fs *fields) appendEmbedfield(field *discordgo.MessageEmbedField) {
	*fs = append(*fs, field)
}

func (fs *fields) setProblemHeader() {
	*fs = append(*fs, &discordgo.MessageEmbedField{
		Name:   "Problem",
		Value:  "----------",
		Inline: true,
	})
	*fs = append(*fs, &discordgo.MessageEmbedField{
		Name:   "Title: ",
		Value:  "----------------------------",
		Inline: true,
	})
	*fs = append(*fs, &discordgo.MessageEmbedField{
		Name:   "Difficulty: ",
		Value:  "----------",
		Inline: true,
	})
}

func (fs *fields) appendReport(report model.ReportsResult) {
	*fs = append(*fs, &discordgo.MessageEmbedField{
		Name:   fmt.Sprintf("No.%d\n", report.ProblemID),
		Value:  "\u200b",
		Inline: true,
	})
	*fs = append(*fs, &discordgo.MessageEmbedField{
		Name:   report.Title,
		Value:  "\u200b",
		Inline: true,
	})
	*fs = append(*fs, &discordgo.MessageEmbedField{
		Name:   difficulty(report.Difficulty),
		Value:  "\u200b",
		Inline: true,
	})
}

func (fs *fields) appendProblem(problem *model.Problem) {
	*fs = append(*fs, &discordgo.MessageEmbedField{
		Name:   fmt.Sprintf("No.%d\n", problem.ProblemID),
		Value:  "\u200b",
		Inline: true,
	})
	*fs = append(*fs, &discordgo.MessageEmbedField{
		Name:   problem.ProblemTitle,
		Value:  "\u200b",
		Inline: true,
	})
	*fs = append(*fs, &discordgo.MessageEmbedField{
		Name:   difficulty(problem.Difficulty),
		Value:  "\u200b",
		Inline: true,
	})
}

func ListReportsMessage(reports []model.ReportsResult, nick string) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	msg.Description = fmt.Sprintf("%s預計要報告的題目如下:\n", nick)

	fields := baseField()
	fields.setProblemHeader()

	for _, report := range reports {
		fields.appendReport(report)
	}
	msg.Fields = fields
	return msg
}

func AddReportMessage(problems []*model.Problem, nick string, month, day int) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	msg.Description = fmt.Sprintf("%s成功添加了以下題目:", nick)
	fields := baseField()
	fields.setProblemHeader()

	for _, problem := range problems {
		fields.appendProblem(problem)
	}

	fields.appendEmbedfield(&discordgo.MessageEmbedField{
		Name:   fmt.Sprintf("讀書會舉辦日期: %d月%d日", month, day),
		Value:  "--------------------",
		Inline: false,
	})
	msg.Fields = fields
	return msg
}

func DeleteReportMessage(problems []*model.Problem, nick string) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	msg.Description = fmt.Sprintf("%s 成功刪除了以下題目:", nick)
	fields := baseField()
	fields.setProblemHeader()

	for _, problem := range problems {
		fields.appendProblem(problem)
	}

	msg.Fields = fields
	return msg
}

func ClearReportMessage() *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	msg.Description = fmt.Sprintf("a")
	return msg
}

func SearchProblemsMsg(problem *model.Problem) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	fields := baseField()

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

	msg.Fields = fields
	return msg
}

func SearchProblemsDetailMsg(problem *model.Problem) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	fields := baseField()

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

	msg.Fields = fields
	msg.Timestamp = time.Now().Format("2006-0102 15:04")
	return msg
}

func SetStudyGroupStartTimeMessage(SID uuid.UUID, exist bool) *discordgo.MessageEmbed {
	sg := model.GetStudyGroup(SID)
	msg := baseEmbedMessage()
	fields := baseField()

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

func StudyGroupInfoMessage(sg *model.StudyGroup, cnt int64, reports []model.ReportsResult) *discordgo.MessageEmbed {
	msg := baseEmbedMessage()
	fields := baseField()

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

func HelpMessage() *discordgo.MessageEmbed {
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

func ClearMessage(username string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Color:  color,
		Author: author,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  fmt.Sprintf("已經成功把%s下次讀書會要報告的所有題目清空。", username),
				Value: "\u200b",
			},
		},
	}
}
