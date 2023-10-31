package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) NotaTransaksi(pegawai string) {

	// as.DB.Joins("JOIN barangs ON barangs.ID = nota_transactions.barang_id").Where("customer_id = ?", id).Find(&Nota)

	as.ListCustomer()
	fmt.Print("\nID customer: ")
	var id int
	fmt.Scanln(&id)

	var barangID []int
	as.DB.Model(model.Nota_Transactions{}).Select("barang_id").Where("customer_id = ?", id).Find(&barangID)

	var qty []int
	as.DB.Model(model.Nota_Transactions{}).Select("qty").Where("customer_id = ?", id).Find(&qty)

	var Nota []model.Nota_Transactions
	as.DB.Model(model.Nota_Transactions{}).Where("customer_id = ?", id).Find(&Nota)

	var created_at []string
	as.DB.Model(model.Nota_Transactions{}).Select("created_at").Where("customer_id = ?", id).Find(&created_at)

	var customer string
	as.DB.Model(model.Customer{}).Select("nama").Where("ID = ?", id).Find(&customer)

	var alamat string
	as.DB.Model(model.Customer{}).Select("alamat_customer").Where("ID = ?", id).Find(&alamat)

	var currentlist int
	var listpembelian []int
	for x := 0; x < len(barangID); x++ {
		as.DB.Model(model.Barang{}).Where("ID = ?", barangID[x]).Select("harga").Take(&currentlist)
		listpembelian = append(listpembelian, currentlist)
	}

	var currentstring string
	var nama_barang []string
	for x := 0; x < len(barangID); x++ {
		as.DB.Model(model.Barang{}).Where("ID = ?", barangID[x]).Select("nama").Take(&currentstring)
		nama_barang = append(nama_barang, currentstring)
	}

	var currstok int
	var oldstok, newstok []int
	for x := 0; x < len(barangID); x++ {
		as.DB.Model(model.Barang{}).Where("ID = ?", barangID[x]).Select("stok_barang").Take(&currstok)
		oldstok = append(oldstok, currstok)
	}
	for x := 0; x < len(barangID); x++ {
		y := 0
		y = oldstok[x] - qty[x]
		newstok = append(newstok, y)
	}

	var hargatotal int
	for i := 0; i < len(qty); i++ {
		hargatotal += qty[i] * listpembelian[i]
	}
	as.DB.Model(model.Barang{}).Select("stok_barang")
	if hargatotal == 0 {
		fmt.Println("sudah lunas")
		fmt.Print("''press enter''")
		fmt.Scanln()
		return
	}
	fmt.Println("===============")
	fmt.Print("Customer: ", customer)
	fmt.Println()
	fmt.Print("Alamat\t: ", alamat)
	fmt.Println()
	fmt.Print("Pegawai\t: ", pegawai)
	fmt.Println("\n===============")
	fmt.Println("\nPesanan\t:")
	fmt.Println()
	fmt.Println("Total\t Barang\t\t Harga_Satuan\t Harga_Total\t Waktu Pemesanan")
	fmt.Println()
	for i := 0; i < len(nama_barang); i++ {
		fmt.Println(qty[i], "\t", nama_barang[i], "\t", listpembelian[i], "\t\t", listpembelian[i]*qty[i], "\t\t", created_at[i])
	}
	fmt.Println()
	fmt.Println("\n===============")
	fmt.Print("Costumer: ", customer, " \nAlamat  : ", alamat, "\nTotal   : ", hargatotal)
	fmt.Println("\n===============")
	fmt.Print("\nSudah bayar(y/n): ")
	var bayar string
	fmt.Scanln(&bayar)
	if bayar == `y` || bayar == `Y` {

		for i := 0; i < len(barangID); i++ {
			as.DelNota(id)
		}

		as.ListBarang()
		fmt.Println("Terima kasih!")
		fmt.Print("''press enter''")
		fmt.Scanln()
		return
	}
	fmt.Println("belum bayar")
	fmt.Print("''press enter''")
	fmt.Scanln()
}
