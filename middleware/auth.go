package middleware

import (
	"Checkrr/Helpers"
	"Checkrr/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Check if user is authorized for the data being asked

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get Auth Cookie
		cookie, err := c.Cookie("Authorization")
		Helpers.Log(err, "Getting auth cookie")
		if err != nil {
			c.JSON(401, gin.H{
				"result": "unauthorized",
			})
			return
		}

		var request Helpers.Request
		// Get UserID from Body
		err = c.ShouldBind(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		jwtClaims, result := auth.ParseJWT(cookie)
		if result == false {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		userId := int64(jwtClaims["userId"].(float64))
		if userId != request.UserId {
			c.JSON(401, "Unauthorized Request")
			return
		}

		// Successful Auth
		c.Next()
	}
}

func ValidAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("Authorization")
		Helpers.Log(err, "Getting auth cookie")
		if err != nil {
			c.JSON(401, gin.H{
				"result": "unauthorized",
			})
			return
		}
		_, result := auth.ParseJWT(cookie)
		if result == false {
			c.JSON(http.StatusUnauthorized, "")
			return
		}
		c.Next()
	}
}
