package main

import (
	"Golang_udemy/todo_app/app/controllers"
	"Golang_udemy/todo_app/app/models"
	"fmt"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
