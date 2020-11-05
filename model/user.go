package model

type User struct {
	UserID   string `gorm:"primaryKey;"`
	UserName string
}

func IsRecorded(userID, userName string) {
	var user User
	err := DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		DB.Create(&User{UserID: userID, UserName: userName})
	}
}
