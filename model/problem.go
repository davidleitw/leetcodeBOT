package model

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const leetcodeURL = "https://leetcode.com/problems/"

type Problem struct {
	ProblemID    int    `gorm:"primaryKey"`
	ProblemTitle string `gorm:"not null"`
	ProblemURL   string
	Difficulty   int
	PaidOnly     bool
	Submit       int
	Accept       int
}

func SearchWithProblemID(ID int) (*Problem, error) {
	if ID <= 0 {
		return nil, errors.New("ID有誤， 請確認輸入是否正確.")
	}
	var problem Problem
	err := DB.Where("problem_id = ?", ID).First(&problem).Error

	// no data in our database.
	if err != nil {
		return nil, err
	}

	return &problem, nil
}

func SearchWithProblemTitle(title string) (*Problem, error) {
	return &Problem{}, nil
}

func SearchWithProblemIDTest(ID int) (*Problem, error) {
	res, err := http.Get("https://leetcode.com/api/problems/algorithms/")
	if err != nil {
		log.Println("http get leetcode api data failed.")
		return nil, err
	}
	defer res.Body.Close()

	var leetcodeAPIData LeetCode
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&leetcodeAPIData)
	if err != nil {
		log.Println("decoder leetcode api data failed.")
		return nil, err
	}

	if ID <= 0 || ID > len(leetcodeAPIData.StatStatusPairs) {
		return nil, errors.New("ID有誤， 請確認輸入是否正確.")
	}

	for _, problem := range leetcodeAPIData.StatStatusPairs {
		if problem.Stat.QuestionID == ID {
			return &Problem{
				ProblemID:    problem.Stat.QuestionID,
				ProblemTitle: problem.Stat.QuestionTitle,
				ProblemURL:   leetcodeURL + problem.Stat.QuestionTitleSlug,
				Difficulty:   problem.Difficulty.Level,
				PaidOnly:     problem.PaidOnly,
				Submit:       problem.Stat.TotalSubmitted,
				Accept:       problem.Stat.TotalAcs,
			}, nil
		}
	}
	return nil, nil
}
