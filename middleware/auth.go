package auth

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDatabase(db *gorm.DB) {
	DB = db
}

func LoginHandler(c *gin.Context) {
	var input models.User
	input.Email = c.PostForm("email")
	input.Password = c.PostForm("password")

	var user models.User
	if err := DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.Redirect(http.StatusFound, "/?error=User not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.Redirect(http.StatusFound, "/?error=invalid Password. please enter a valid password")
		return
	}
	tokenString, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	SetTokenCookie(c, tokenString)

	if user.IsSuperUser {
		c.Redirect(http.StatusFound, "/admin")
	} else {
		c.Redirect(http.StatusFound, "/home")
	}

}
