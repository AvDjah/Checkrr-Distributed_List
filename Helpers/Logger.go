package Helpers

import "fmt"

func Log(err error, msg string) {
	if err != nil {
		fmt.Println("Error at : ", msg)
	}
}
