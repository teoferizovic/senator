package model

import "time"

type Post struct {
	//gorm.Model
	ID uint `gorm:"column:Id;primary_key:auto_increment"`
	Body string `gorm:"column:body" json:"body"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp;column:created_at;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp;column:updated_at;" json:"updated_at"`
}
