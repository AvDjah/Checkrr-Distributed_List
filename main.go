package main

import (
	"Checkrr/Helpers"
	"Checkrr/Router"
	"Checkrr/Services"
	"fmt"
)

func main() {

	fmt.Println("Heelo World")

	go Services.RunWorker()

	//_ := Db.InitDb()
	//users := Data.GetAllUser(db)
	//fmt.Println("Got Users: ", len(users))
	//
	r := Router.New()
	err := r.Run()
	if err != nil {
		Helpers.Log(err, "Error while listening")
	}

}
