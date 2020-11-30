package models

//Admin model
type Admin struct {
	UserID   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Username string `json:"username"`
	Password string `json:"password"`
}
