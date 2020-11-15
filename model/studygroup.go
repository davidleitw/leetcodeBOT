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

type ReportsResult struct {
	ProblemID  int
	Title      string
	Difficulty int
}

func DeleteUserReports(UserID, GuildID string) {
	s_id := VerifyStudyGroup(GuildID)
	DB.Where("user_id = ? AND s_id = ?", UserID, s_id).Delete(Report{})
	UpdateSgAttendance(s_id)
}

// use user_id and s_id search user's reports
func GetUserReports(SID uuid.UUID, UserID string) []ReportsResult {
	var results []ReportsResult
	DB.Table("problems").Select(
		"problems.problem_id as problem_id, problems.problem_title as title, problems.difficulty as difficulty").Joins(
		"left join reports on reports.problem_id = problems.problem_id where reports.user_id = ? and reports.s_id = ? order by reports.problem_id", UserID, SID).Scan(&results)
	return results
}

func GetStudyGroupReports(SID uuid.UUID) []ReportsResult {
	var results []ReportsResult
	DB.Table("problems").Select(
		"problems.problem_id as problem_id, problems.problem_title as title, problems.difficulty as difficulty").Joins(
		"left join reports on reports.problem_id = problems.problem_id where reports.s_id = ? order by reports.problem_id", SID).Scan(&results)
	return results
}

func CreateNewReport(UserID string, ProblemID int, SID uuid.UUID) {
	DB.Create(&Report{UserID: UserID, ProblemID: ProblemID, SID: SID})
}

func DeleteReport(UserID string, ProblemID int, SID uuid.UUID) {
	DB.Where("user_id = ? AND problem_id = ? AND s_id = ?", UserID, ProblemID, SID).Delete(&Report{})
}

func GetStudyGroupWithGID(GuildID string) (*StudyGroup, error) {
	var sg StudyGroup
	err := DB.Where("guild_id = ? AND next_turn = ?", GuildID, true).First(&sg).Error

	if err != nil {
		return nil, err
	}
	return &sg, nil
}

func GetStudyGroup(SID uuid.UUID) *StudyGroup {
	var sg StudyGroup
	DB.Where("s_id = ?", SID).Find(&sg)
	return &sg
}

// 總共有幾題會報告
func CountStudyGroupProblems(SID uuid.UUID) int64 {
	var count int64
	DB.Model(&Report{}).Where("s_id = ?", SID).Count(&count)
	return count
}

// true => study group is exist
func IsStudyGroupExist(GuildID string) bool {
	err := DB.Where("guild_id = ? AND next_turn = ?", GuildID, true).First(&StudyGroup{}).Error
	return err != nil
}

func SetStudyGroupTime(GuildID string, t time.Time) (uuid.UUID, bool) {
	var sgexist bool
	var sg StudyGroup

	err := DB.Where("guild_id = ? AND next_turn = ?", GuildID, true).First(&sg).Error
	if err != nil {
		sgexist = false

		Sid := uuid.NewV4()
		DB.Create(&StudyGroup{
			SID:        Sid,
			Turn:       1,
			GuildID:    GuildID,
			Attendance: 1,
			NextTurn:   true,
			StartTime:  Time(t),
		})

		return Sid, sgexist
	} else {
		sgexist = true
		sg.StartTime = Time(t)
		DB.Save(&sg)
		return sg.SID, sgexist
	}
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

// 不存在指定的report return true
func VerifyReport(UserID string, ProblemID int, SID uuid.UUID) bool {
	err := DB.Where("user_id = ? AND problem_id = ? AND s_id = ?", UserID, ProblemID, SID).First(&Report{}).Error
	return err != nil
}
