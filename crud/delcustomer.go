package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) DelCustomer() {
	var currentUser = new(model.Customer)
	var user int
	fmt.Print("ID User: ")
	fmt.Scanln(&user)

	qry := as.DB.Table("customer").Where("ID = ?", user).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nCustomer tidak ditemukan.")
		return
	}

	fmt.Println("\nCustomer ", currentUser.Nama, " berhasil dihapus.")
	as.DB.Table("customer").Where("ID = ?", user).Delete(&currentUser)
}
