package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Report struct {
	UserID    string    `gorm:"not null;"`
	ProblemID int       `gorm:"not null;"`
	SID       uuid.UUID `gorm:"type: char(36); not null;"`
}

type StudyGroup struct {
	SID        uuid.UUID `gorm:"type: char(36); primaryKey; not null; unique;"`
	Turn       uint16
	GuildID    string
	Attendance uint16
	NextTurn   bool
	StartTime  Time `gorm:"not null;"`
}

func CreateNewReport(UserID string, ProblemID int, SID uuid.UUID) {
	DB.Create(&Report{UserID: UserID, ProblemID: ProblemID, SID: SID})
}

func SetStudyGroupTime(GuildID string, command []string) (time.Time, error) {
	//return nil, nil
}

// 更新study group 人數狀態並且回傳舉辦時間
func UpdateSgAttendance(SID uuid.UUID) (int, int) {
	var sg StudyGroup
	var attend int64
	DB.Where("s_id = ?", SID).Find(&sg)
	DB.Model(&Report{}).Distinct("user_id").Count(&attend)
	sg.Attendance = uint16(attend)
	DB.Save(&sg)

	t := time.Time(sg.StartTime)
	return int(t.Month()), int(t.Day())
}

// if studygroup update, remeber also update new studygroup turn && old studygroup next turn value.
func VerifyStudyGroup(GuildID string) uuid.UUID {
	var sg StudyGroup
	err := DB.Where("guild_id = ? AND next_turn = ?", GuildID, true).First(&sg).Error
	if err != nil {
		Sid := uuid.NewV4()
		var addDay int = (7 - int(time.Now().Weekday())) % 7

		// 獲得下個週日的晚上12:59
		endOfTime := time.Unix(GetTodayEnd(), 0).AddDate(0, 0, addDay)
		DB.Create(&StudyGroup{
			SID:        Sid,
			Turn:       1,
			GuildID:    GuildID,
			Attendance: 1,
			NextTurn:   true,
			StartTime:  Time(endOfTime),
		})

		return Sid
	}
	t1 := time.Now().Unix()
	t2 := time.Time(sg.StartTime).Unix()

	if t1 >= t2 {
		Sid := uuid.NewV4()
		var addDay int = (7 - int(time.Now().Weekday())) % 7

		// 獲得下個週日的晚上12:59
		endOfTime := time.Unix(GetTodayEnd(), 0).AddDate(0, 0, addDay)
		DB.Create(&StudyGroup{
			SID:        Sid,
			Turn:       sg.Turn + 1,
			GuildID:    GuildID,
			Attendance: 1,
			NextTurn:   true,
			StartTime:  Time(endOfTime),
		})
		sg.NextTurn = false
		DB.Save(&sg)
		return Sid
	}
	return sg.SID
}

func VerifyReport(UserID string, ProblemID int, SID uuid.UUID) bool {
	err := DB.Where("user_id = ? AND problem_id = ? AND s_id = ?", UserID, ProblemID, SID).First(&Report{}).Error
	// if err != nil 代表已經重複創立report
	return err != nil
}
