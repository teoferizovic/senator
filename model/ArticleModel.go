package model

import (
	"github.com/teoferizovic/senator/database"
	"time"
)

type Article struct {
	ID int `gorm:"column:id;primary_key:auto_increment"`
	UserID int `gorm:"column:user_id;not null" json:"user_id"`
	Headline string `gorm:"type:text;column:headline;not null" json:"headline"`
	Content string `gorm:"type:text;column:content;not null" json:"content"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at;" json:"updated_at"`
	User *User `json:",omitempty"`
	Comments []*Comment `json:",omitempty"`
}

//get all articles
func GetArticles(articles *[]Article) (err error) {

	if err := database.DBCon.Preload("User").Find(&articles).Error; err != nil {
		return err
	}
	return nil

}

//find Article by id
func GetArticleById(articles *[]Article, id string) (err error) {

	if err := database.DBCon.Preload("User").Where("id = ?", id).Find(&articles).Error; err != nil {
		return err
	}

	return nil

}

//find Article by user_id
func GetArticleByUserId(articles *[]Article, userId string) (err error) {

	if err := database.DBCon.Preload("User").Where("user_id = ?", userId).Find(&articles).Error; err != nil {
		return err
	}

	return nil

}

//find Article by user_id
func GetArticleByIdAndUserId(articles *[]Article, id string, userId string) (error error) {

	if err := database.DBCon.Preload("User").Where("id = ? AND user_id = ?", id, userId).Find(&articles).Error; err != nil {
		return err
	}

	return nil

}