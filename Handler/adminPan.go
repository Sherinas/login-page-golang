package Handler

import (
	"main/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AdminGet(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	claims := &jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !parsedToken.Valid {
		c.Redirect(http.StatusFound, "/")
		return
	}

	userID := (*claims)["user_id"].(float64)

	var user models.User
	if err := DB.Where("id = ?", uint(userID)).First(&user).Error; err != nil || !user.IsSuperUser {
		c.Redirect(http.StatusFound, "/")
		return
	}

	// Fetch all users
	var users []models.User
	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch users"})
		return
	}

	c.HTML(http.StatusOK, "admin.html", gin.H{"users": users})
}
