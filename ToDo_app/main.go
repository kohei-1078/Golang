package main

import (
	"fmt"
	// "log"
	// "ToDo_app/config"
	"ToDo_app/app/controllers"
	"ToDo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}