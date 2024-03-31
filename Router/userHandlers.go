package Router

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"Checkrr/auth"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var user Models.User

	// Bind user data from request body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user input (name, userID)
	if err := validateUserInput(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)

	// Store User
	db := Db.InitDb()

	exists := Data.GetUserByUsername(db, user.UserId)
	if exists.UserId != "" {
		c.JSON(400, "User Already Exists")
		return
	}

	result := Data.UpsertUser(db, user)

	if result == true {
		c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
	} else {
		c.JSON(http.StatusConflict, "Could not create user.")
	}

}

func validateUserInput(user *Models.User) error {
	// Implement your validation logic here
	// (e.g., check for empty fields, invalid characters)
	if user.Name == "" || user.UserId == "" {
		return errors.New("name and user ID are required")
	}
	// ... add more validations as needed
	return nil
}

func LoginUser(c *gin.Context) {
	var loginData struct {
		UserID   string `json:"userId"`
		Password string `json:"password"`
	}

	// Bind login data from request body
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user Models.User
	db := Db.InitDb()
	err := db.Where("user_id = ?", loginData.UserID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID or password"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID or password"})
		return
	}

	// Login successful, generate JWT token (implementation not included here)
	jwtToken, err := auth.GenerateJWT(user.Id, user.Name, "User")
	Helpers.Log(err, "Generating Token")

	if err != nil {
		c.JSON(401, "Trouble Logging In")
	}

	name, err := c.Cookie("Authorization")

	if err != nil {
		fmt.Println("err getting cookie", err)
	} else {
		fmt.Println("name:", name)
	}

	c.SetCookie("Authorization", jwtToken, 6*3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "login successful", "token": jwtToken})
}

func GetUserDetails(c *gin.Context) {

	authHeader, err := c.Cookie("Authorization")
	Helpers.Log(err, "Getting Authorization Cookie")

	if err != nil || authHeader == "" || len(authHeader) == 0 {
		c.JSON(401, "invalid authorization header format")
		return
	}
	jwtToken := authHeader

	claims, isAuthorized := auth.ParseJWT(jwtToken)

	userId := int64(claims["userId"].(float64))

	if !isAuthorized {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized to access user details"})
		return
	}

	var user Models.User
	db := Db.InitDb()
	err = db.Where("id = ?", userId).Preload("Categories").Preload("EventLists").Preload("Events").Preload("Notifications").Preload("Subscriptions").First(&user).Error

	res := Models.UserResponse{
		Id:            user.Id,
		Name:          user.Name,
		Subscriptions: user.Subscriptions,
		Categories:    user.Categories,
		UserId:        user.UserId,
		EventLists:    user.EventLists,
		Events:        user.Events,
		Notifications: user.Notifications,
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
	}
	c.JSON(200, res)
}
