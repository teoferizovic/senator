package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/teoferizovic/senator/model"
	"net/http"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// Get the JWT string from the cookie
		tknStr := ctx.Request.Header.Get("Authentication")

		// Initialize a new instance of `Claims`
		claims := &model.Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				respondWithError(ctx, http.StatusUnauthorized, "Unauthorized")
				return
			}
			respondWithError(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		if !tkn.Valid {
			respondWithError(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		ctx.Next()

	}
}