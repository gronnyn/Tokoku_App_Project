package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) UpdateBarang() {
	fmt.Println()
	as.ListBarang()
	var currentUser = new(model.Barang)
	var kode, stok int

	fmt.Print("\nKode Barang: ")
	fmt.Scanln(&kode)

	qry := as.DB.Table("barangs").Where("ID = ?", kode).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nBarang tidak ditemukan.")
		return
	}

	fmt.Print("Stok: ")
	fmt.Scanln(&stok)
	as.DB.Table("barangs").Where("ID = ?", kode).Update("stok_barang", stok)
	currentUser.Stok_Barang = stok

	fmt.Println()
	as.ListBarang()
}
