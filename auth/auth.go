package auth

import (
	"Module/config"
	"Module/crud"
	"Module/model"
	"fmt"

	"gorm.io/gorm"
)

type AuthSystem struct {
	DB *gorm.DB
}

var authSystem = crud.CrudSystem{}

func (as *AuthSystem) Login() (string, int) {
	var currentUser = new(model.User)
	var user, pass string
	fmt.Print("Username: ")
	fmt.Scanln(&user)
	if user == "" {
		fmt.Println("Username tidak boleh kosong!")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return user, 0
	}
	fmt.Print("Password: ")
	fmt.Scanln(&pass)

	qry := as.DB.Table("users").Where("username = ? AND password = ?", user, pass).Take(&currentUser)

	err := qry.Error

	if err != nil {
		config.CallClear()
		fmt.Printf("\t\tTokoku Project Start")
		fmt.Println("\n\n1. Login")
		fmt.Println("0. Exit")
		fmt.Print("\n ")
		fmt.Println("Username / password salah.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return user, 0
	}

	if user == "admin" {
		config.CallClear()
		return user, 1
	}

	config.CallClear()
	return user, 2
}
