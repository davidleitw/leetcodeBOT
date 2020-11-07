package bot

import (
	"sort"
	"strconv"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
)

func Add(msg *discordgo.MessageCreate, command []string) (*discordgo.MessageEmbed, error) {
	if len(command) < 2 {
		return nil, FORMAT_ERROR
	}
	model.IsRecorded(msg.Author.ID, msg.Member.Nick)
	// 記得update study group 參加人數
	// get study group ID
	SID := model.VerifyStudyGroup(msg.GuildID)

	var wg sync.WaitGroup
	var cnt int = 0
	wg.Add(len(command) - 1)

	problemsID := make([]int, 0)

	// 驗證每筆problem id 是否是整數而且是DB裡面有的題目
	for idx := 1; idx < len(command); idx++ {
		go func(idx int) {
			defer wg.Done()
			if id, ok := strconv.Atoi(command[idx]); ok == nil && model.VerifyProblem(id) {
				if model.VerifyReport(msg.Author.ID, id, SID) {
					problemsID = append(problemsID, id)
					model.CreateNewReport(msg.Author.ID, id, SID)
				} else {
					cnt++
				}
			}
		}(idx)
	}
	wg.Wait()

	// 表示所有題目都是已經添加過的題目
	if cnt == len(command)-1 {
		return nil, ADD_REPORT_REPEAT_ERROR
	}

	if len(problemsID) == 0 {
		return nil, ADD_REPORT_NUMBER_ERROR
	}

	sort.Ints(problemsID)
	problems := make([]*model.Problem, 0)

	for _, id := range problemsID {
		problem, _ := model.SearchWithProblemID(id)
		problems = append(problems, problem)
	}

	month, day := model.UpdateSgAttendance(SID)
	message := AddReportMessage(problems, msg.Member.Nick, month, day)

	return message, nil
}
