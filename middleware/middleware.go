package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte("secret"))
}

func SetTokenCookie(c *gin.Context, tokenString string) {
	c.SetCookie("token", tokenString, 3600*72, "/", "localhost", false, true)
}
