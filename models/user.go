package models

import "time"

type User struct {
	ID          int       `gorm:"primary_key"`
	Name        string    `gorm:"type:varchar(100)"`
	IsSuperUser bool      `gorm:"default:false"`
	Email       string    `json:"email" gorm:"type:varchar(100);unique_index"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
