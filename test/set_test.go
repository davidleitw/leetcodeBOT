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

func TestSet(t *testing.T) {
	model.SetDB(getTestDB())

	tests := []struct {
		guildID string
		command []string
	}{
		{
			guildID: test_guildID,
			command: []string{
				"set",
				"2021-03-08 14:29",
			},
		},
		{
			// Format error
			guildID: test_guildID,
			command: []string{
				"set",
				"20",
				"30",
			},
		},
		{
			// time format error
			guildID: test_guildID,
			command: []string{
				"set",
				"2021-03-46 13-28",
			},
		},
		{
			// set time early error
			guildID: test_guildID,
			command: []string{
				"set",
				"2009-03-08 19-24",
			},
		},
	}

	results := []struct {
		message *discordgo.MessageEmbed
		err     error
	}{
		{
			message: &discordgo.MessageEmbed{
				Color: 0xFFD700,
				Author: &discordgo.MessageEmbedAuthor{
					URL:     "https://github.com/davidleitw/leetcodeBOT",
					Name:    "薯條醬",
					IconURL: "https://img.icons8.com/doodle/2x/github.png",
				},
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "成功創建了一次新的讀書會，基本資訊如下:",
						Value:  "--------------------",
						Inline: false,
					},
					&discordgo.MessageEmbedField{
						Name:   "預計開始時間",
						Value:  "2021-03-08 14:29",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "目前參與人數",
						Value:  fmt.Sprintf("%d人", 1),
						Inline: true,
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
			err:     bot.FORMAT_ERROR,
		},
		{
			message: nil,
			err:     bot.SET_TIME_FORMAT_ERROR,
		},
		{
			message: nil,
			err:     bot.SET_TIME_EARLY_ERROR,
		},
	}

	for index, test := range tests {

		m, err := bot.Set(test.guildID, test.command)
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
