package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/bot"
	"github.com/davidleitw/leetcodeBOT/model"
	"gopkg.in/go-playground/assert.v1"
)

func TestAdd(t *testing.T) {
	model.SetDB(getTestDB())
	tests := []struct {
		msg     *discordgo.Message
		command []string
	}{
		{
			msg: &discordgo.Message{
				Author: &discordgo.User{
					ID: test_userID,
				},
				GuildID: test_guildID,
				Member: &discordgo.Member{
					Nick: "user",
				},
			},
			command: []string{
				"add",
				"120",
				"130",
			},
		},
		{
			msg: &discordgo.Message{
				Author: &discordgo.User{
					ID: test_userID,
				},
				GuildID: test_guildID,
				Member: &discordgo.Member{
					Nick: "user",
				},
			},
			command: []string{
				"add",
				"測試文字應該要被忽略",
				"120", // 重複的資料應該也要被忽略
				"150",
			},
		},
		{
			// it should be format error.
			msg: &discordgo.Message{
				Author: &discordgo.User{
					ID: test_userID,
				},
				GuildID: test_guildID,
				Member: &discordgo.Member{
					Nick: "user",
				},
			},
			command: []string{ // ADD_REPORT_REPEAT_ERROR
				"add",
				"120",
				"130",
			},
		},
		{
			// it should be format error.
			msg: &discordgo.Message{
				Author: &discordgo.User{
					ID: test_userID,
				},
				GuildID: test_guildID,
			},
			command: []string{
				"add",
				"abc",
			},
		},
	}

	results := []struct {
		message *discordgo.MessageEmbed
		err     error
	}{
		{
			message: &discordgo.MessageEmbed{
				Color:       0xFFD700,
				Description: fmt.Sprintf("%s成功添加了以下題目:", "user"),
				Author: &discordgo.MessageEmbedAuthor{
					URL:     "https://github.com/davidleitw/leetcodeBOT",
					Name:    "薯條醬",
					IconURL: "https://img.icons8.com/doodle/2x/github.png",
				},
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "Problem",
						Value:  "----------",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Title: ",
						Value:  "----------------------------",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Difficulty: ",
						Value:  "----------",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   fmt.Sprintf("No.%d\n", 120),
						Value:  "\u200b",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Triangle",
						Value:  "\u200b",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Medium",
						Value:  "\u200b",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   fmt.Sprintf("No.%d\n", 130),
						Value:  "\u200b",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Surrounded Regions",
						Value:  "\u200b",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Medium",
						Value:  "\u200b",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   fmt.Sprintf("讀書會舉辦日期: %d月%d日", 11, 22),
						Value:  "--------------------",
						Inline: false,
					},
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "  一起刷leetcode 一起進步 (･ω´･ )",
					IconURL: "https://cdn.icon-icons.com/icons2/2389/PNG/512/leetcode_logo_icon_145113.png",
				},
			},
			err: nil,
		},
		{
			message: &discordgo.MessageEmbed{
				Color:       0xFFD700,
				Description: fmt.Sprintf("%s成功添加了以下題目:", "user"),
				Author: &discordgo.MessageEmbedAuthor{
					URL:     "https://github.com/davidleitw/leetcodeBOT",
					Name:    "薯條醬",
					IconURL: "https://img.icons8.com/doodle/2x/github.png",
				},
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "Problem",
						Value:  "----------",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Title: ",
						Value:  "----------------------------",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Difficulty: ",
						Value:  "----------",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   fmt.Sprintf("No.%d\n", 150),
						Value:  "\u200b",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Evaluate Reverse Polish Notation",
						Value:  "\u200b",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Medium",
						Value:  "\u200b",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   fmt.Sprintf("讀書會舉辦日期: %d月%d日", 11, 22),
						Value:  "--------------------",
						Inline: false,
					},
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "  一起刷leetcode 一起進步 (･ω´･ )",
					IconURL: "https://cdn.icon-icons.com/icons2/2389/PNG/512/leetcode_logo_icon_145113.png",
				},
			},
			err: nil,
		},
		{
			message: nil,
			err:     bot.ADD_REPORT_REPEAT_ERROR,
		},
		{
			message: nil,
			err:     bot.ADD_REPORT_NUMBER_ERROR,
		},
	}

	for index, test := range tests {
		m, err := bot.Add(&discordgo.MessageCreate{
			test.msg,
		}, test.command)
		log.Printf("test case%d: ", index)
		if m != nil && err == nil {
			assert.Equal(t, m.Author, results[index].message.Author)
			assert.Equal(t, m.Color, results[index].message.Color)
			assert.Equal(t, m.Footer, results[index].message.Footer)
			// assert.Equal(t, m.Fields, results[index].message.Fields)
			for i := 0; i < len(m.Fields); i++ {
				assert.Equal(t, m.Fields[i].Name, results[index].message.Fields[i].Name)
				assert.Equal(t, m.Fields[i].Value, results[index].message.Fields[i].Value)
			}
		}

		if err != nil {
			assert.Equal(t, err, results[index].err)
		}
	}
}
