package model

type User struct {
	Model
	Name  string `json:"name" gorm:"not null;unique;"`
	Email string `json:"email" gorm:"not null;unique;"`
	AppID string `json:"app_id" gorm:"not null;unique;"`
}
