package bot

import (
	"time"

	"github.com/davidleitw/leetcodeBOT/model"

	"github.com/bwmarrin/discordgo"
)

const TIME_LAYOUT = "2006-01-02 15:04"

func Set(GuildID string, command []string) (*discordgo.MessageEmbed, error) {
	if len(command) > 2 {
		return nil, FORMAT_ERROR
	}

	startTime, err := time.ParseInLocation(TIME_LAYOUT, command[1]+" "+command[2], time.Local)
	if err != nil {
		return nil, SET_TIME_FORMAT_ERROR
	}

	t1 := time.Now().Unix()
	t2 := startTime.Unix()

	if t1 >= t2 {
		return nil, SET_TIME_EARLY_ERROR
	}

	SID, studyGroupExist := model.SetStudyGroupTime(GuildID, startTime)
	message := SetStudyGroupStartTimeMessage(SID, studyGroupExist)

	return message, nil
}
