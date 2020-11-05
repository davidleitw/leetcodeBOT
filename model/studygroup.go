package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Report struct {
	UserID    string    `gorm:"not null;"`
	ProblemID string    `gorm:"not null;"`
	SID       uuid.UUID `gorm:"type: char(36); not null;"`
}

type StudyGroup struct {
	SID        uuid.UUID `gorm:"type: char(36); primaryKey; not null;"`
	Turn       uint16
	GuildID    string
	Attendance uint16
	NextTurn   bool
	StartTime  time.Time `gorm:"not null;"`
}

func SearchProblemsWithUserID(GuildID, UserID string) []*Problem {

	return nil
}
