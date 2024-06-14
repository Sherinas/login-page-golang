package Handler

import (
	"main/models"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {
	var input models.User
	input.Name = c.PostForm("name")
	input.Email = c.PostForm("email")
	input.Password = c.PostForm("password")

	confpass := c.PostForm("pasconf")

	var errors = make(map[string]string)

	// Validate email address
	if !utils.ValidateEmailAddress(input.Email) {
		errors["error1"] = "Invalid email address"
	}

	// Validate password complexity
	if !utils.CheckPasswordComplexity(input.Password) {
		errors["error2"] = "Password must be at least 4 characters long and include a mix of uppercase and lowercase letters"
	}

	// Check if passwords match
	if input.Password != confpass {
		errors["error3"] = "Passwords do not match"
	}

	// If there are errors, render the registration page with error messages
	if len(errors) > 0 {
		c.HTML(http.StatusOK, "register.html", errors)
		return
	}

	// Check if user already exists

	var user models.User
	if err := DB.Where("email = ?", input.Email).First(&user).Error; err == nil {
		// c.Redirect(http.StatusFound, "/register?error=email_taken")
		c.HTML(http.StatusOK, "register.html", gin.H{"error": "This Email already registered"})
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
