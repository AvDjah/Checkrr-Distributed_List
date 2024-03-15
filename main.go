package main

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Helpers"
	"Checkrr/Router"
	"fmt"
)

func main() {

	fmt.Println("Heelo World")

	db := Db.InitDb()
	users := Data.GetAllUser(db)
	fmt.Println("Got Users: ", len(users))

	r := Router.New()
	err := r.Run()
	if err != nil {
		Helpers.Log(err, "Error while listening")
	}

}
