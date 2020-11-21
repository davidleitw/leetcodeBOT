package test

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestAdd(t *testing.T) {
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
			},
			command: []string{
				"add",
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
	fmt.Println(tests)
}
