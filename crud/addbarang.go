package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) TambahBarang() {

	var newUser = new(model.Barang)
	fmt.Print("Nama: ")
	fmt.Scanln(&newUser.Nama)
	fmt.Print("Harga: ")
	fmt.Scanln(&newUser.Harga)
	fmt.Print("Stok: ")
	fmt.Scanln(&newUser.Stok_Barang)

	err := as.DB.Table("barang").Create(newUser).Error

	if err != nil {
		fmt.Println("gagal membuat barang baru")
		return
	} else {
		fmt.Println("\nBarang " + newUser.Nama + " berhasil dibuat.")
	}
}
