package model

type user struct {
	UserID      string
	UserName    string
	UserGuildID string
}

type StudyGroup struct {
	User user
}
