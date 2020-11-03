package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// 要報告的每一題
type User struct {
	UserID   string `gorm:"primaryKey;"`
	UserName string
}

type Report struct {
	UserID    string `gorm:"not null;"`
	ProblemID string `gorm:"not null;"`
	SID       string `gorm:"type: char(36);"`
}

type StudyGroup struct {
	SID        uuid.UUID `gorm:"type: char(36); primaryKey; not null;	"`
	GuildID    string
	Attendance uint16
	StartTime  time.Time
}

func SearchProblemsWithUserID(GuildID, UserID string) []*Problem {
	return nil
}
