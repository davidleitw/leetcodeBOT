package model

import (
	"encoding/json"
	"log"
	"net/http"
)

// leetcode api data struct
type LeetCode struct {
	UserName  string `json:"user_name"`
	NumSolved int    `json:"num_solved"`
	NumTotal  int    `json:"num_total"`
	AcEasy    int    `json:"ac_easy"`
	AcMedium  int    `json:"ac_medium"`
	AcHard    int    `json:"ac_hard"`

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

		Status interface{} `json:"status"` //

		Difficulty struct {
			Level int `json:"level"`
		} `json:"difficulty"`

		PaidOnly  bool `json:"paid_only"`
		IsFavor   bool `json:"is_favor"`
		Frequency int  `json:"frequency"`
		Progress  int  `json:"progress"`
	} `json:"stat_status_pairs"`
}

func CreateLeetCodeProblemsTable() {
	//res, err := http.Get("https://leetcode.com/api/problems/algorithms/")
	res, err := http.Get("https://leetcode.com/api/problems/all/")
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
			ProblemID:    problem.Stat.FrontendQuestionID,
			ProblemTitle: problem.Stat.QuestionTitle,
			ProblemURL:   leetcodeURL + problem.Stat.QuestionTitleSlug,
			Difficulty:   problem.Difficulty.Level,
			PaidOnly:     problem.PaidOnly,
			Submit:       problem.Stat.TotalSubmitted,
			Accept:       problem.Stat.TotalAcs,
		}
		log.Println("Problem ID = ", p.ProblemID, ", Title = ", p.ProblemTitle)
		AddLeetCodeProblem(p)
	}
}

func AddLeetCodeProblem(problem *Problem) {
	DB.Create(&problem)
}
