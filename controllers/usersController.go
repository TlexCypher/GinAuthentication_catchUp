package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/TlexCypher/ginAuthenticationCatchUp/initializes"
	"github.com/TlexCypher/ginAuthenticationCatchUp/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	type body struct {
		Email    string
		Password string
	}
	var b body
	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed to bind json.",
		})
		return
	}

	hashSalt := 10
	hp, err := bcrypt.GenerateFromPassword([]byte(b.Password), hashSalt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed to hash pass.",
		})
		return
	}

	user := models.User{Email: b.Email, Password: string(hp)}
	res := initializes.DB.Create(&user)
	if res != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed to register user into db.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	type body struct {
		Email    string
		Password string
	}
	var b body
	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed to bind json.",
		})
		return
	}

	var user models.User
	initializes.DB.First(&user, "email = ?", b.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed to find such user.",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(b.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid Password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString(os.Getenv("SECRET"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Failed to create token.",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}
