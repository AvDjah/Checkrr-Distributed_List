package Helpers

import (
	"Checkrr/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type Request struct {
	UserId int64 `form:"userId"`
}

func GetUserIdFromQuery(c *gin.Context) int64 {
	var request Request

	err := c.ShouldBind(&request)
	if err != nil {
		Log(err, "Binding Query String")
		return 0
	} else {
		return request.UserId
	}
}

func GetUserIdFromJWTClaim(c *gin.Context) int64 {

	cookie, err := c.Cookie("Authorization")
	if err != nil {
		Log(err, "Getting Cookie from request")
		return 0
	}
	// Get Claims
	claims, result := auth.ParseJWT(cookie)
	if result == false {
		return 0
	}

	return int64(claims["userId"].(float64))
}

func ExtractCookieFromRequest(c *gin.Context) (string, bool) {

	cookie, err := c.Cookie("Authorization")
	if err != nil {
		Log(err, "Extracting Cookie from request")
		return "", false
	} else {
		return cookie, true
	}
}

func GetPriorityId(value string) int64 {
	if value == "Low" {
		return 1
	} else if value == "Medium" {
		return 2
	} else if value == "High" {
		return 3
	} else {
		return 0
	}
}

func addDurationToTime(timeStr string, durationStr string) (string, error) {
	// Parse the input time string
	inputTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return "", err
	}

	// Parse the duration string
	var dur time.Duration
	if strings.Contains(durationStr, "h") {
		dur, err = time.ParseDuration(strings.ReplaceAll(durationStr, "h", "h0m0s"))
	} else if strings.Contains(durationStr, "m") {
		dur, err = time.ParseDuration(strings.ReplaceAll(durationStr, "m", "m0s"))
	} else if strings.Contains(durationStr, "s") {
		dur, err = time.ParseDuration(durationStr)
	} else {
		return "", fmt.Errorf("duration format not recognized")
	}
	if err != nil {
		return "", err
	}

	// Add the duration to the input time
	resultTime := inputTime.Add(dur)

	// Format the result time to hh:mm:ss
	resultTimeStr := resultTime.Format("15:04:05")

	return resultTimeStr, nil
}
