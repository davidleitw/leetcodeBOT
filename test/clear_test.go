package test

import (
	"fmt"
	"testing"

	"github.com/davidleitw/leetcodeBOT/bot"
	"gopkg.in/go-playground/assert.v1"

	"github.com/bwmarrin/discordgo"
)

func TestClear(t *testing.T) {
	tests := []struct {
		msg     *discordgo.Message
		command []string
	}{
		{
			// it should be pass.
			msg: &discordgo.Message{
				Author: &discordgo.User{
					ID:            test_userID,
					Username:      test_userName,
					Discriminator: "48763",
				},
				GuildID: test_guildID,
				Member: &discordgo.Member{
					Nick: "user",
				},
			},
			command: []string{
				"clear",
			},
		},
		{

			msg: &discordgo.Message{
				Author: &discordgo.User{
					ID:            test_userID,
					Username:      test_userName,
					Discriminator: "48763",
				},
				GuildID: test_guildID,
				Member: &discordgo.Member{
					Nick: "user",
				},
			},
			command: []string{
				"clear",
				"123",
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
						Name:  fmt.Sprintf("已經成功把%s下次讀書會要報告的所有題目清空。", test_userName+"#"+"48763"),
						Value: "\u200b",
					},
				},
			},
			err: nil,
		},
		{
			message: nil,
			err:     bot.FORMAT_ERROR,
		},
	}
	for index, test := range tests {
		m, err := bot.Clear(&discordgo.MessageCreate{
			test.msg,
		}, test.command)

		if m != nil && err == nil {
			assert.Equal(t, m.Author, results[index].message.Author)
			assert.Equal(t, m.Color, results[index].message.Color)
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
