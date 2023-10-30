package crud

import (
	"Module/config"
	"Module/model"
	"fmt"
)

func (as *CrudSystem) OrderPesanan(pegawai string) {
	config.CallClear()
	fmt.Println("\n===============")
	fmt.Println("Login as:", pegawai)
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Fitur Pegawai")
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Membuat Transaksi")
	fmt.Println("===============")
	var newUser = new(model.Nota_Transactions)
	var stok_barang int
	newUser.Username = pegawai
	as.ListBarang()
	fmt.Print("\nID Barang: ")
	fmt.Scanln(&newUser.Barang_ID)
	as.DB.Model(model.Barang{}).Select("stok_barang").Where("ID = ?", newUser.Barang_ID).Take(&stok_barang)
	fmt.Print("Total Barang: ")
	fmt.Scanln(&newUser.Qty)
	if stok_barang-newUser.Qty < 0 {
		fmt.Println("stok tidak cukup!")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}
	x := stok_barang - newUser.Qty
	as.DB.Model(model.Barang{}).Where("ID = ?", newUser.Barang_ID).Update("stok_barang", x)

	as.ListCustomer()
	fmt.Print("\nID Customer: ")
	fmt.Scanln(&newUser.Customer_ID)

	err := as.DB.Table("nota_transactions").Create(newUser).Error
	as.DB.Table("nota_transactions_backup").Create(newUser)

	if err != nil {
		fmt.Println("gagal membuat nota transaksi.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	} else {
		fmt.Println("\nNota transaksi atas customer dengan ID ", newUser.Customer_ID, " berhasil dibuat.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}
}
