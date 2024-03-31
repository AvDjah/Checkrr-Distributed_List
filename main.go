package main

import (
	"Checkrr/Db"
	"Checkrr/Helpers"
	"Checkrr/Router"
	"Checkrr/Services"
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	_ = Db.InitDb()
	viper.SetConfigName("")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		Helpers.Log(err, "Reading Env")
	}
	fmt.Println("Heelo World")

	go Services.RunWorker()

	//_ := Db.InitDb()
	//users := Data.GetAllUser(db)
	//fmt.Println("Got Users: ", len(users))
	//
	r := Router.New()
	err = r.Run()
	if err != nil {
		Helpers.Log(err, "Error while listening")
	}

}
