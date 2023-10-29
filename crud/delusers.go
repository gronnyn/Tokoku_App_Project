package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) DelUsers() {
	var currentUser = new(model.User)
	var user string
	fmt.Print("Username: ")
	fmt.Scanln(&user)

	qry := as.DB.Table("users").Where("username = ?", user).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nusername tidak ditemukan.")
		return
	} else if user == "admin" {
		fmt.Println("\nusername admin tidak dapat dihapus.")
		return
	}

	fmt.Println("\nusername " + user + " dihapus.")
	as.DB.Table("users").Where("username = ?", user).Delete(&currentUser)
}
