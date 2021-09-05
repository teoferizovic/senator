package model

import (
	"fmt"
	"github.com/teoferizovic/senator/database"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID uint
	Email string
	Username  string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUsers() *gorm.DB {

	var users []User

	if err := database.DBCon.Find(&users).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(users)
	}
	return database.DBCon
}