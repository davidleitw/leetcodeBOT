package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
)

func List(msg *discordgo.MessageCreate, command []string) (*discordgo.MessageEmbed, error) {
	sg, err := model.GetStudyGroupWithGID(msg.GuildID)
	if err != nil {
		return nil, STUDYGROUP_NOT_FOUND_ERROR
	}

	reports := model.GetUserReports(sg.SID, msg.Author.ID)

	if len(reports) == 0 {
		return nil, LIST_NO_DATA_ERROR
	}

	var message *discordgo.MessageEmbed
	if msg.Member.Nick != "" {
		message = ListReportsMessage(reports, msg.Member.Nick)
	} else {
		message = ListReportsMessage(reports, msg.Author.Username+"#"+msg.Author.Discriminator+" ")
	}
	return message, nil
}
