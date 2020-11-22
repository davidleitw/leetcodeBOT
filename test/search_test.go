package test

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/bot"
	"github.com/davidleitw/leetcodeBOT/model"
	"gopkg.in/go-playground/assert.v1"
)

func TestSearch(t *testing.T) {
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
			},
			command: []string{
				"search", // normal
				"100",
			},
		},
		{
			msg: &discordgo.Message{
				Author: &discordgo.User{
					ID: test_userID,
				},
				GuildID: test_guildID,
			},
			command: []string{
				"search",
				"一百五十", // SEARCH_NUMBER_ERROR
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
				"search",
				"48763", // KIRITO_ERROR
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
						Name:   "Problem",
						Value:  "No.100\n",
						Inline: false,
					},
					&discordgo.MessageEmbedField{
						Name:   "Title: ",
						Value:  "Same Tree",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "URL: ",
						Value:  "https://leetcode.com/problems/same-tree",
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name:   "Difficulty: ",
						Value:  "Eazy",
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
			err:     bot.SEARCH_NUMBER_ERROR,
		},
		{
			message: nil,
			err:     bot.KIRITO_ERROR,
		},
	}
	for index, test := range tests {
		m, err := bot.Search(test.command)
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
