package main

import (
	"fmt"

	"github.com/iamDushu/lenslocked/models"
)

func main() {

	cfg := models.DefaultPostgresConfig()

	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	userService := models.UserService{
		DB: db,
	}

	user, err := userService.Create("dushyanth1@gmail.com", "1234")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
