package Handler

import (
	"main/models"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserCreateFrmAdmin(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "adduser.html", gin.H{})
}

func UserCreate(c *gin.Context) {
	var input models.User
	input.Name = c.PostForm("name")
	input.Email = c.PostForm("email")
	input.Password = c.PostForm("password")
	confpass := c.PostForm("pasconf")

	if !utils.ValidateEmailAddress(input.Email) {
		c.HTML(http.StatusOK, "register.html", gin.H{"error": "Invalid email address"})
		return
	}

	if !utils.CheckPasswordComplexity(input.Password) {
		c.HTML(http.StatusOK, "register.html", gin.H{"error": "Password should be at least 4 characters long and include a mix of uppercase and lowercase letters"})
		return
	}

	if input.Password != confpass {
		c.Redirect(http.StatusFound, "/register")
		c.HTML(http.StatusOK, "register.html", gin.H{"error": "Password not matching"})
		return
	}

	var user models.User
	if err := DB.Where("email = ?", input.Email).First(&user).Error; err == nil {
		c.HTML(http.StatusOK, "register.html", gin.H{"error": "This Email already registered"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	input.Password = string(hashedPassword)

	if err := DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.Redirect(http.StatusFound, "/admin")
}
