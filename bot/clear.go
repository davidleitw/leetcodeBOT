package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
)

func Clear(msg *discordgo.MessageCreate, command []string) (*discordgo.MessageEmbed, error) {
	if len(command) > 2 {
		return nil, FORMAT_ERROR
	}
	model.DeleteUserReports(msg.Author.ID, msg.GuildID)
	return ClearMessage(msg.Author.Username + "#" + msg.Author.Discriminator), nil
}
