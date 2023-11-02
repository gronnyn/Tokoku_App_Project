package crud

import (
	"Module/config"
	"Module/model"
	"fmt"
)

func (as *CrudSystem) OrderPesanan(pegawai string) {
	var isCust int
	as.DB.Model(model.Customer{}).Select("ID").Find(&isCust)
	if isCust == 0 {
		fmt.Println("Customer tidak ada.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	config.CallClear()
	fmt.Println("\n===============")
	fmt.Println("Login as:", pegawai)
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Fitur Pegawai")
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Membuat Pesanan")
	fmt.Println("===============")
	var newTransactions = new(model.Nota_Transactions)
	var stok_barang, id_barang int
	newTransactions.Username = pegawai
	as.ListBarang()
	fmt.Print("\nID Barang: ")
	fmt.Scanln(&newTransactions.Barang_ID)
	as.DB.Model(model.Barang{}).Select("id").Where("id = ?", newTransactions.Barang_ID).Find(&id_barang)
	if id_barang == 0 {
		fmt.Println("Barang tidak tersedia.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	fmt.Print("Total Barang: ")
	fmt.Scanln(&newTransactions.Qty)

	if newTransactions.Qty == 0 {
		fmt.Println("Jumlah pesanan tidak boleh kosong!")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	as.DB.Model(model.Barang{}).Select("stok_barang").Where("ID = ?", newTransactions.Barang_ID).Take(&stok_barang)
	if stok_barang-newTransactions.Qty < 0 {
		fmt.Println("stok tidak cukup!")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	as.ListCustomer()
	fmt.Print("\nID Customer: ")
	fmt.Scanln(&newTransactions.Customer_ID)

	err := as.DB.Table("nota_transactions").Create(newTransactions).Error
	if err != nil {
		fmt.Println("gagal membuat pesanan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	x := stok_barang - newTransactions.Qty
	as.DB.Model(model.Barang{}).Where("ID = ?", newTransactions.Barang_ID).Update("stok_barang", x)
	as.ListBarang()
	as.DB.Table("nota_transactions_backup").Create(newTransactions)
	fmt.Println("\nPesanan atas customer dengan ID", newTransactions.Customer_ID, "berhasil dibuat.")
	fmt.Print("'press enter'")
	fmt.Scanln()
	return
}
