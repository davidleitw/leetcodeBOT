package model

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const leetcodeURL = "https://leetcode.com/problems/"

// leetcode api data struct
type LeetCode struct {
	UserName        string `json:"user_name"`
	NumSolved       int    `json:"num_solved"`
	NumTotal        int    `json:"num_total"`
	AcEasy          int    `json:"ac_easy"`
	AcMedium        int    `json:"ac_medium"`
	AcHard          int    `json:"ac_hard"`
	StatStatusPairs []struct {
		Stat struct {
			QuestionID          int         `json:"question_id"`
			QuestionArticleLive interface{} `json:"question__article__live"` //
			QuestionArticleSlug interface{} `json:"question__article__slug"` //
			QuestionTitle       string      `json:"question__title"`
			QuestionTitleSlug   string      `json:"question__title_slug"`
			QuestionHide        bool        `json:"question__hide"`
			TotalAcs            int         `json:"total_acs"`
			TotalSubmitted      int         `json:"total_submitted"`
			FrontendQuestionID  int         `json:"frontend_question_id"`
			IsNewQuestion       bool        `json:"is_new_question"`
		} `json:"stat"`
		Status     interface{} `json:"status"` //
		Difficulty struct {
			Level int `json:"level"`
		} `json:"difficulty"`
		PaidOnly  bool `json:"paid_only"`
		IsFavor   bool `json:"is_favor"`
		Frequency int  `json:"frequency"`
		Progress  int  `json:"progress"`
	} `json:"stat_status_pairs"`
}

type Problem struct {
	ProblemID    int
	ProblemTitle string
	ProblemURL   string
	Difficulty   int
	PaidOnly     bool
	Submit       int
	Accept       int
}

func CreateLeetCodeProblemsTable() {
	res, err := http.Get("https://leetcode.com/api/problems/algorithms/")
	if err != nil {
		log.Println("http get leetcode api data failed.")
		panic(err)
	}
	defer res.Body.Close()

	var leetcodeAPIData LeetCode
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&leetcodeAPIData)
	if err != nil {
		log.Println("decoder leetcode api data failed.")
		panic(err)
	}

	log.Println(len(leetcodeAPIData.StatStatusPairs))
	for _, problem := range leetcodeAPIData.StatStatusPairs {
		p := &Problem{
			ProblemID:    problem.Stat.QuestionID,
			ProblemTitle: problem.Stat.QuestionTitle,
			ProblemURL:   leetcodeURL + problem.Stat.QuestionTitleSlug,
			Difficulty:   problem.Difficulty.Level,
			PaidOnly:     problem.PaidOnly,
			Submit:       problem.Stat.TotalSubmitted,
			Accept:       problem.Stat.TotalAcs,
		}
		log.Println(p)
		AddLeetCodeProblem(p)
	}
}

func AddLeetCodeProblem(problem *Problem) {

}

func SearchWithID(ID int) (*Problem, error) {
	if ID <= 0 || ID >= 2500 {
		return nil, errors.New("ID有誤， 請確認輸入是否正確.")
	}

	return &Problem{}, nil
}

func SearchWithTitle(title string) (*Problem, error) {
	return &Problem{}, nil
}
