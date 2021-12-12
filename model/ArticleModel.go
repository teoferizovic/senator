package model

import (
	"errors"
	"github.com/teoferizovic/senator/database"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID int `gorm:"column:id;primary_key:auto_increment"`
	UserID int `gorm:"column:user_id" json:"user_id"`
	Headline string `gorm:"column:headline" json:"headline"`
	Content string `gorm:"type:text;column:content" json:"content"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at;" json:"updated_at"`
	User *User `json:",omitempty"`
}

//get all articles
func GetArticles() (error error, articles []Article) {

	var resultArticles []Article

	err := database.DBCon.Preload("User").Find(&resultArticles).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return err, resultArticles
	}

	return nil, resultArticles

}

//find Article by id
func GetArticleById(id string) (error error, articles []Article) {

	var resultArticles []Article

	err := database.DBCon.Preload("User").Where("id = ?", id).Find(&resultArticles).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return err, resultArticles
	}

	return nil, resultArticles

}

//find Article by user_id
func GetArticleByUserId(userId string) (error error, articles []Article) {

	var resultArticles []Article

	err := database.DBCon.Preload("User").Where("user_id = ?", userId).Find(&resultArticles).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return err, resultArticles
	}

	return nil, resultArticles

}

//find Article by user_id
func GetArticleByIdAndUserId(id string, userId string) (error error, articles []Article) {

	var resultArticles []Article

	err := database.DBCon.Preload("User").Where("id = ? AND user_id = ?", id, userId).Find(&resultArticles).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return err, resultArticles
	}

	return nil, resultArticles

}