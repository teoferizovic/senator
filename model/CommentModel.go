package model

import (
	"github.com/teoferizovic/senator/database"
	"time"
)

type Comment struct {
	ID int `gorm:"column:id;primary_key:auto_increment"`
	UserID int `gorm:"column:user_id;not null" json:"user_id"`
	ArticleID int `gorm:"column:article_id;not null" json:"article_id"`
	Content string `gorm:"type:text;column:content" json:"content"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at;" json:"updated_at"`
	User *User `json:",omitempty"`
	Article *Article `json:",omitempty"`
}

//get all articles
func GetComments(comments *[]Comment) (err error) {

	if err := database.DBCon.Preload("User").Preload("Article.User").Find(&comments).Error; err != nil {
		return err
	}

	return nil

}

//find comment by id
func GetCommentById(comments *[]Comment, id string) (err error) {

	if err := database.DBCon.Preload("User").Preload("Article.User").Where("id = ?", id).Find(&comments).Error; err != nil {
		return err
	}

	return nil

}