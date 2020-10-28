package leetcode

import (
	"github.com/davidleitw/leetcodeBOT/model"
)

func SearchProblemsWithID(ID int) (*model.Problem, error) {
	return model.SearchWithID(ID)
}

func SearchProblemsWithTitle(Title string) (*model.Problem, error) {
	return model.SearchWithTitle(Title)
}
