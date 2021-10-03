package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func UserRegister(ctx *gin.Context) {

	var requestUser model.User

	//check if right credeitals are sent
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	newUser, err := model.CreateUser(&requestUser)

	//if email still exists or insert problems
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//return 200
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created user with email:" + newUser.Email,
	})
	return
}

func UserLogin(ctx *gin.Context) {

	var requestUser model.User

	//check if right credeitals are sent
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//check if user with email exits
	resultUser := model.GetByEmail(&requestUser)

	if (model.User{}) == resultUser {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "No user with email:" + requestUser.Email,
		})
		return
	}

	//check if password is ok
	err := bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(requestUser.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials",
		})
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: requestUser.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Token for user is:" + tokenString,
	})
	return

}

/*func UserLogout(ctx *gin.Context) {

	claims := &Claims{}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(nil, claims)

	_, err := token.SignedString(jwtKey)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal ServerError",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged out",
	})
	return
}*/

func UserIndex(ctx *gin.Context) {

	// Get the JWT string from the cookie
	tknStr := ctx.Request.Header.Get("Authentication")

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	if !tkn.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Index",
	})
	return

}
