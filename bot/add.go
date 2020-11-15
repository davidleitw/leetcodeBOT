package bot

import (
	"sort"
	"strconv"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
)

func Add(msg *discordgo.MessageCreate, command []string) (*discordgo.MessageEmbed, error) {
	length := len(command)

	if length < 2 {
		return nil, FORMAT_ERROR
	}

	if length > 5 {
		return nil, ADD_TOO_MANY_REPORTS
	}

	if msg.Member.Nick != "" {
		model.IsRecorded(msg.Author.ID, msg.Member.Nick)
	} else {
		model.IsRecorded(msg.Author.ID, msg.Author.Username+"#"+msg.Author.Discriminator)
	}

	// 記得update study group 參加人數
	// get study group ID
	sid := model.VerifyStudyGroup(msg.GuildID)

	var failedCommand int = 0
	var wg sync.WaitGroup
	wg.Add(length - 1)

	problemsID := make([]int, 0)

	// 驗證每筆problem id 是否是整數而且是DB裡面有的題目
	for idx := 1; idx < length; idx++ {
		go func(idx int) {
			defer wg.Done()
			if pid, ok := strconv.Atoi(command[idx]); ok == nil && model.VerifyProblem(pid) {
				if model.VerifyReport(msg.Author.ID, pid, sid) {
					problemsID = append(problemsID, pid)
					model.CreateNewReport(msg.Author.ID, pid, sid)
				} else {
					failedCommand++
				}
			}
		}(idx)
	}
	wg.Wait()

	// 表示所有題目都是已經添加過的題目
	if failedCommand == length-1 {
		return nil, ADD_REPORT_REPEAT_ERROR
	}

	if len(problemsID) == 0 {
		return nil, ADD_REPORT_NUMBER_ERROR
	}

	sort.Ints(problemsID)
	problems := make([]*model.Problem, 0)

	for _, pid := range problemsID {
		problem, _ := model.SearchWithProblemID(pid)
		problems = append(problems, problem)
	}

	month, day := model.UpdateSgAttendance(sid)

	if msg.Member.Nick != "" {
		return AddReportMessage(problems, msg.Member.Nick, month, day), nil
	} else {
		return AddReportMessage(problems, msg.Author.Username+"#"+msg.Author.Discriminator+" ", month, day), nil
	}
}
