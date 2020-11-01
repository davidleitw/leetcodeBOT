package model

// 要報告的每一題
type User struct {
	UserID    string
	GuildID   string
	ProblemID int
}

type StudyGroup struct {
	GuildID   string `gorm:"primary_key"`
	StudyTurn int
}

func SearchProblemsWithUserID(GuildID, UserID string) []*Problem {
	return nil
}
