package Helpers

import (
	"fmt"
	"strconv"
	"time"
)

func GetTimeFromUnixTime(unixTimeStr string) string {
	unixTimeInt, err := strconv.ParseInt(unixTimeStr, 10, 64)
	if err != nil {
		panic(err) // Handle the error appropriately in a real application
	}

	unixTime := time.Unix(unixTimeInt, 0)
	unixTime = unixTime.Truncate(time.Minute)
	formattedTime := unixTime.Format("15:04:05") // HH:MM:SS format

	fmt.Println("Formatted time:", formattedTime)
	return formattedTime
}
