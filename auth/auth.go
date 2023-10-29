package auth

import (
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
	fmt.Print("Password: ")
	fmt.Scanln(&pass)

	qry := as.DB.Table("users").Where("username = ? AND password = ?", user, pass).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("username / password salah.")
		return user, 0
	}
	if user == "admin" {
		return user, 1
	}
	return user, 2
}
