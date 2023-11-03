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

	as.ListCustomer()
	var idCustomer, isCustomer int
	fmt.Print("\nID Customer: ")
	fmt.Scanln(&idCustomer)
	as.DB.Model(model.Customer{}).Select("id").Where("id = ?", idCustomer).Find(&isCustomer)

	if isCustomer == 0 {
		fmt.Println("Customer dengan ID", idCustomer, "tidak ada.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	for {
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
		var stok_barang, id_barang int
		var newTransactions = new(model.Nota_Transactions)
		newTransactions.Username = pegawai
		newTransactions.Customer_ID = idCustomer
		as.ListBarang()
		fmt.Println("\n0. Selesai")
		fmt.Print("\nID Barang: ")
		fmt.Scanln(&newTransactions.Barang_ID)
		if newTransactions.Barang_ID == 0 {
			return
		}
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

		as.DB.Table("nota_transactions").Create(newTransactions)
		x := stok_barang - newTransactions.Qty
		as.DB.Model(model.Barang{}).Where("ID = ?", newTransactions.Barang_ID).Update("stok_barang", x)
		as.ListBarang()

		var qty []int
		as.DB.Model(model.Nota_Transactions{}).Select("qty").Where("customer_id = ?", idCustomer).Find(&qty)
		var customer string
		as.DB.Model(model.Customer{}).Select("nama").Where("ID = ?", idCustomer).Find(&customer)
		var barangID []int
		as.DB.Model(model.Nota_Transactions{}).Select("barang_id").Where("customer_id = ?", idCustomer).Find(&barangID)
		var alamat string
		as.DB.Model(model.Customer{}).Select("alamat_customer").Where("ID = ?", idCustomer).Find(&alamat)
		var currentstring string
		var nama_barang []string
		for x := 0; x < len(barangID); x++ {
			as.DB.Model(model.Barang{}).Where("ID = ?", barangID[x]).Select("nama").Take(&currentstring)
			nama_barang = append(nama_barang, currentstring)
		}
		var currentlist int
		var listpembelian []int
		for x := 0; x < len(barangID); x++ {
			as.DB.Model(model.Barang{}).Where("ID = ?", barangID[x]).Select("harga").Take(&currentlist)
			listpembelian = append(listpembelian, currentlist)
		}
		var hargatotal int
		for i := 0; i < len(qty); i++ {
			hargatotal += qty[i] * listpembelian[i]
		}

		fmt.Println("\nPesanan\t:")
		fmt.Println()
		fmt.Println("Total\t Barang\t\t Harga_Satuan\t Harga_Total")
		fmt.Println()
		for i := 0; i < len(nama_barang); i++ {
			fmt.Println(qty[i], "\t", nama_barang[i], "\t", listpembelian[i], "\t\t", listpembelian[i]*qty[i])
		}
		fmt.Println()
		fmt.Println("\n===============")
		fmt.Print("Customer : ", customer, "\nTotal    : ", hargatotal)
		fmt.Println("\n===============")
		as.DB.Table("nota_transactions_backup").Create(newTransactions)
		fmt.Println("\nPesanan berhasil dibuat.")
		fmt.Print("'press enter'")
		fmt.Scanln()
	}

}
