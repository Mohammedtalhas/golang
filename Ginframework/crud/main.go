//main.go
package main

import (
	"fmt"

	"github.com/Mohammedtalhas/golang/tree/main/GinFramework/crud/Config"
	"github.com/Mohammedtalhas/golang/tree/main/GinFramework/crud/Models"
	"github.com/Mohammedtalhas/golang/tree/main/GinFramework/crud/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
