package model

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/teoferizovic/senator/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID uint `gorm:"column:ID;primary_key:auto_increment"`
	Email string `gorm:"column:email" json:"email" validate:"min=1,max=16,regexp=^[a-zA-Z]*$" binding:"required"`
	Password string `gorm:"column:password" json:"password" validate:"min=1,max=16" binding:"required"`
	Active bool `gorm:"column:active;default:0"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:null" json:"deleted_at"`
}

//create user
func CreateUser(requestUser *User)  (User, error) {

	err := database.DBCon.Create(requestUser)

	if err != nil {
		return *requestUser, err.Error
	}

	return *requestUser, nil

}

//find user by email
func GetByEmail(requestUser *User) (User) {

	var resultUser User

	database.DBCon.Where("email = ?", requestUser.Email).First(&resultUser)

	return resultUser

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

// BeforeSave : hook before a user is saved
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	hash, err := MakePassword(u.Password)
	if err != nil {
		return nil
	}
	u.Password = hash
	return nil
}

// MakePassword : Encrypt user password
func MakePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

var JwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}