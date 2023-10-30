package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) ListBarang() {
	var barang []model.Barang

	as.DB.Find(&barang)
	fmt.Print("{ID, Nama Barang, Harga, Stok Barang}")
	fmt.Println()
	fmt.Println()
	for i := 0; i < len(barang); i++ {
		fmt.Println(barang[i])
	}
}
