package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) DelBarang() {
	var currentUser = new(model.Barang)
	var barang string
	fmt.Print("\nKode Barang: ")
	fmt.Scanln(&barang)

	qry := as.DB.Table("barangs").Where("ID = ?", barang).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nBarang tidak ditemukan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	fmt.Println("\nBarang dengan kode " + barang + " berhasil dihapus.")
	as.DB.Table("barangs").Where("ID = ?", barang).Delete(&currentUser)
	fmt.Print("'press enter'")
	fmt.Scanln()
}

func (as *CrudSystem) DelCustomer() {
	var currentUser = new(model.Customer)
	var user int
	fmt.Print("\nID User: ")
	fmt.Scanln(&user)

	qry := as.DB.Table("customers").Where("ID = ?", user).Take(&currentUser)
	err := qry.Error

	if err != nil {
		fmt.Println("\nCustomer tidak ditemukan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	fmt.Println("\nCustomer", currentUser.Nama, "berhasil dihapus.")
	as.DB.Table("customers").Where("ID = ?", user).Delete(&currentUser)
	fmt.Print("'press enter'")
	fmt.Scanln()
}

func (as *CrudSystem) DelUsers() {
	var currentUser = new(model.User)
	var user string
	fmt.Print("Hapus Pegawai: ")
	fmt.Scanln(&user)

	qry := as.DB.Table("users").Where("username = ?", user).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nPegawai tidak ditemukan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	} else if user == "admin" {
		fmt.Println("\nusername admin tidak dapat dihapus.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	fmt.Println("\nusername " + user + " berhasil dihapus.")
	as.DB.Table("users").Where("username = ?", user).Delete(&currentUser)
	fmt.Print("'press enter'")
	fmt.Scanln()
}

func (as *CrudSystem) DelNota(customers int) {
	var currentUser = new(model.Nota_Transactions)

	qry := as.DB.Table("nota_transactions").Where("customer_id = ?", customers).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nNota tidak dapat ditemukan.")
		return
	}

	fmt.Println("\nNota", customers, "dihapus.")
	as.DB.Table("nota_transactions").Where("customer_id = ?", customers).Delete(&currentUser)
}
