package Handler

import (
	"main/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDatabase(db *gorm.DB) {
	DB = db
}

func AddUser_frm_Reg(c *gin.Context, DB *gorm.DB) {
	var input models.User
	input.Name = c.PostForm("name")
	input.Email = c.PostForm("email")
	input.Password = c.PostForm("password")

	confpass := c.PostForm("pasconf")

	if input.Password != confpass {
		c.Redirect(http.StatusFound, "/register")
		c.HTML(http.StatusOK, "register.html", gin.H{"error": "Password not matching"})

		return

	}

	var user models.User
	if err := DB.Where("email = ?", input.Email).First(&user).Error; err == nil {
		c.Redirect(http.StatusFound, "/register?error=email_taken")
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	input.Password = string(hashedPassword)

	// Create user
	if err := DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func LogOut(c *gin.Context) {
	// Clear the token cookie by setting an expired cookie
	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	// Redirect the user to the login page
	c.Redirect(http.StatusFound, "/")
}

func UserDelete(c *gin.Context) {
	userID := c.PostForm("user_id")
	var user models.User

	if err := DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	// Check if the user is a superuser
	if user.IsSuperUser {
		c.Redirect(http.StatusFound, "/admin")
		return
	}

	// Delete the user
	if err := DB.Delete(&models.User{}, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
		return
	}
	c.Redirect(http.StatusFound, "/admin")
	//c.HTML(http.StatusOK, "admin.html", gin.H{})
}

func UserSerch(c *gin.Context) {

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
	/////////////////////////////////////////////////////////
	query := c.Query("query")
	var users []models.User
	if err := DB.Where("name LIKE ? OR email LIKE ?", "%"+query+"%", "%"+query+"%").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not search users"})
		return
	}

	c.HTML(http.StatusOK, "admin.html", gin.H{"users": users})
}

func UserEdit(c *gin.Context) {
	userID := c.PostForm("user_id")
	var user models.User
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	password := c.PostForm("password")

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
			return
		}
		user.Password = string(hashedPassword)
	}

	if err := DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}

	c.Redirect(http.StatusFound, "/admin")
}
