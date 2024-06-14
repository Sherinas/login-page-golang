package Handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SignIN(ctx *gin.Context) {

	errorMessage := ctx.Query("error")
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Error": errorMessage,
	})
}

func GetHome(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	_, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.HTML(http.StatusOK, "home.html", nil)
}

func RegForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}
