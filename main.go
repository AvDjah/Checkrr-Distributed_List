package main

import (
	"Checkrr/Db"
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"Checkrr/Router"
	"fmt"
)

func main() {

	fmt.Println("Heelo World")

	db := Db.InitDb()
	user := Models.User{
		Name:     "Arvind",
		Password: "Heelo",
		UserId:   "arvind20",
	}

	tx := db.Create(&user)
	fmt.Println(tx.Error, "\n", tx.RowsAffected)

	r := Router.New()
	err := r.Run()
	if err != nil {
		Helpers.Log(err, "Error while listening")
	}

}
