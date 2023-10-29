package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) TambahCustomer() {

	var newUser = new(model.Customer)
	fmt.Print("Nama Customer: ")
	fmt.Scanln(&newUser.Nama)
	fmt.Print("Alamat Customer: ")
	fmt.Scanln(&newUser.Alamat_Customer)

	err := as.DB.Table("customer").Create(newUser).Error

	if err != nil {
		fmt.Println("gagal menambah customer")
		return
	} else {
		fmt.Println("\nCustomer " + newUser.Nama + " berhasil ditambahkan.")
	}
}
