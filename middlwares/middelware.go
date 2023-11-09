package middlware

import (
	"fmt"
	"golang/gin/model"
	"golang/gin/response"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtClaim struct {
	Email string `json:"user_email"`
	ID    int    `json:"user_id"`
	jwt.StandardClaims
}

var secret_key = "devendra_secretkey"

func Login(r *gin.Context) {

	var user response.Users
	r.ShouldBindJSON(&user)
	userData := model.LoginUser(user.Email, user.Password)
	if userData.Password == "" && userData.Email == "" {
		r.JSON(http.StatusBadGateway, gin.H{
			"error": "email and password not match",
		})
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JwtClaim{Email: user.Email, ID: userData.ID, StandardClaims: jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

	}
	data := map[string]interface{}{"User_Data": token.Claims, "Tokan": tokenString}

	r.JSON(http.StatusAccepted, gin.H{
		"data": data,
	})

}

func AuthMiddleware() gin.HandlerFunc {
	return func(r *gin.Context) {
		tokenString := r.GetHeader("Authorization")
		if tokenString == "" {
			r.JSON(http.StatusBadRequest, gin.H{
				"Error": "Token Missing Please Provide Valide Token",
			})
		}
		tokenString = tokenString[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method")
			}
			return secret_key, nil
		})
		if err != nil || !token.Valid {
			r.JSON(http.StatusBadRequest, gin.H{
				"error": "some occured at the time",
			})

		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			r.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalid claims Token",
			})
		}
		r.Set("login_user", claims)
		r.Next()
	}
}
