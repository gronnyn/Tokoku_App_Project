package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) DelBarang() {
	var isBarangExist int
	as.DB.Model(model.Barang{}).Select("ID").Find(&isBarangExist)
	if isBarangExist == 0 {
		fmt.Println("Barang tidak ditemukan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}
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
	var isCustExist int
	as.DB.Model(model.Customer{}).Select("ID").Find(&isCustExist)
	if isCustExist == 0 {
		fmt.Println("Customer tidak ditemukan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	fmt.Print("\nID User: ")
	var user int
	fmt.Scanln(&user)

	var currentCustomer = new(model.Customer)
	qry := as.DB.Table("customers").Where("ID = ?", user).Take(&currentCustomer)
	err := qry.Error

	if err != nil {
		fmt.Println("\nCustomer tidak ditemukan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	fmt.Println("\nCustomer", currentCustomer.Nama, "berhasil dihapus.")
	qry2 := as.DB.Table("customers").Where("ID = ?", user).Delete(&currentCustomer)
	err2 := qry2.Error

	if err2 != nil {
		fmt.Println("Customer", currentCustomer.Nama, "belum bayar.")
	}
	fmt.Print("'press enter'")
	fmt.Scanln()
}

func (as *CrudSystem) DelCustomerNota(customer int) {
	var currentUser = new(model.Customer)

	qry := as.DB.Table("customers").Where("ID = ?", customer).Take(&currentUser)
	err := qry.Error

	if err != nil {
		fmt.Println("\nCustomer tidak ditemukan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}
	as.DB.Table("customers").Where("ID = ?", customer).Delete(&currentUser)
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
		fmt.Println("\nUsername admin tidak dapat dihapus.")
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

	as.DB.Table("nota_transactions").Where("customer_id = ?", customers).Delete(&currentUser)
}
