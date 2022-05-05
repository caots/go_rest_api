//main.go
package main

import (
	"fmt"
	Config "rest_api/config"
	"rest_api/models"
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
	Config.DB.AutoMigrate(&models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
