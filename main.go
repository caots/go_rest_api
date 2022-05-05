//main.go
package main

import (
	"fmt"
	"rest_api/Models"
	Config "rest_api/config"
	Routes "rest_api/routers"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDbConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
