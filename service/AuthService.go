package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/teoferizovic/senator/model"
	"time"
)

func CreateToken(requestUser model.User) (string,error){

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(10 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &model.Claims{
		Email: requestUser.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(model.JwtKey)

	if err != nil {
		return "",err
	}

	return tokenString, nil

}