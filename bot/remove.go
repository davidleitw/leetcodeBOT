package bot

import (
	"sort"
	"strconv"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/model"
)

func Remove(msg *discordgo.MessageCreate, command []string) (*discordgo.MessageEmbed, error) {
	length := len(command)

	if length < 2 {
		return nil, FORMAT_ERROR
	}

	if length >= 5 {
		return nil, REMOVE_TOO_MANY_REPORTS
	}

	sid := model.VerifyStudyGroup(msg.GuildID)

	var failedCommand int = 0
	var wg sync.WaitGroup
	wg.Add(length - 1)

	problemsID := make([]int, 0)
	for idx := 1; idx < length; idx++ {
		go func(idx int) {
			defer wg.Done()
			if pid, ok := strconv.Atoi(command[idx]); ok == nil && model.VerifyProblem(pid) {
				if !model.VerifyReport(msg.Author.ID, pid, sid) {
					problemsID = append(problemsID, pid)
					model.DeleteReport(msg.Author.ID, pid, sid)
				} else {
					failedCommand++
				}
			}
		}(idx)
	}
	wg.Wait()

	if failedCommand == (length-1) || len(problemsID) == 0 {
		return nil, REMOVE_REPROT_ERROR
	}

	sort.Ints(problemsID)

	problems := make([]*model.Problem, 0)
	for _, pid := range problemsID {
		problem, _ := model.SearchWithProblemID(pid)
		problems = append(problems, problem)
	}

	_, _ = model.UpdateSgAttendance(sid)

	if msg.Member.Nick != "" {
		return DeleteReportMessage(problems, msg.Member.Nick), nil
	} else {
		return DeleteReportMessage(problems, msg.Author.Username+"#"+msg.Author.Discriminator+" "), nil
	}
}
