package bot

import "github.com/bwmarrin/discordgo"

func Set(GuildID string, command []string) (*discordgo.MessageEmbed, error) {
	if len(command) == 1 {
		return nil, FORMAT_ERROR
	}

	return nil, nil
}
