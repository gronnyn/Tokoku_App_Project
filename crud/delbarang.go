package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) DelBarang() {
	var currentUser = new(model.Barang)
	var barang string
	fmt.Print("\nNama Barang: ")
	fmt.Scanln(&barang)

	qry := as.DB.Table("barang").Where("nama = ?", barang).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nBarang tidak ditemukan.")
		return
	}

	fmt.Println("\nbarang " + barang + " berhasil dihapus.")
	as.DB.Table("barang").Where("nama = ?", barang).Delete(&currentUser)
}
