package main

import (
	"Checkrr/Services"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Heelo World")

	go Services.RunWorker()

	time.Sleep(time.Minute * 2)

	//db := Db.InitDb()
	//users := Data.GetAllUser(db)
	//fmt.Println("Got Users: ", len(users))
	//
	//r := Router.New()
	//err := r.Run()
	//if err != nil {
	//	Helpers.Log(err, "Error while listening")
	//}

}
