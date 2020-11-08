package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
)

func GetNextStudyGroupInfo(GuildID string, command []string) (*discordgo.MessageEmbed, error) {
	sg, err := model.GetStudyGroupWithGID(GuildID)

	if err != nil {
		return nil, STUDYGROUP_NOT_FOUND_ERROR
	}

	cnt := model.CountStudyGroupProblems(sg.SID)
	return StudyGroupInfoMessage(sg, cnt), nil
}
