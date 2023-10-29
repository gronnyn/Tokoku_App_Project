package crud

import (
	"Module/model"
	"fmt"

	"gorm.io/gorm"
)

type CrudSystem struct {
	DB *gorm.DB
}

func (as *CrudSystem) TambahUser() (model.User, bool) {

	var newUser = new(model.User)
	fmt.Print("Username: ")
	fmt.Scanln(&newUser.Username)
	fmt.Print("Password: ")
	fmt.Scanln(&newUser.Password)

	err := as.DB.Table("users").Create(newUser).Error

	if err != nil {
		fmt.Println("gagal membuat user baru")
		return model.User{}, false
	} else {
		fmt.Println("\nUser successfully created.")
	}
	return *newUser, true
}
