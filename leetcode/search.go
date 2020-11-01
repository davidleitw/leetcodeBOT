package leetcode

import (
	"github.com/davidleitw/leetcodeBOT/model"
)

func SearchProblemWithID(ID int) (*model.Problem, error) {
	return model.SearchWithProblemID(ID)
}

func SearchProblemWithTitle(Title string) (*model.Problem, error) {
	return model.SearchWithProblemTitle(Title)
}

func SearchProblemsWithUserID(GuildID, UserID string) []*model.Problem {
	return model.SearchProblemsWithUserID(GuildID, UserID)
}
