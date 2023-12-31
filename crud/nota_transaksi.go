package crud

import (
	"Module/config"
	"Module/model"
	"fmt"
	"time"
)

func (as *CrudSystem) NotaTransaksi(pegawai string) {
	var isNotaExist int
	as.DB.Model(model.Nota_Transactions{}).Select("ID").Find(&isNotaExist)
	if isNotaExist == 0 {
		fmt.Println("Tidak ada yang pesan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		config.CallClear()
		return
	}

	as.ListCustomer()
	fmt.Print("\nID customer: ")
	var id int
	fmt.Scanln(&id)

	var isCust int
	as.DB.Model(model.Customer{}).Select("id = ?", id).Where("id = ?", id).Find(&isCust)
	if isCust == 0 {
		fmt.Println("Costumer tidak ada.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		config.CallClear()
		return
	}
	var ListNota int
	as.DB.Model(model.Nota_Transactions{}).Select("customer_id = ?", id).Where("customer_id = ?", id).Find(&ListNota)
	if ListNota == 0 {
		fmt.Println("Costumer belum pesan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		config.CallClear()
		return
	}
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
	fmt.Println("===============")
	fmt.Print("Pegawai\t: ", pegawai)
	fmt.Println("\n===============")
	fmt.Println("\nPesanan\t:")
	fmt.Println()
	fmt.Println("Total\t Barang\t\t Harga_Satuan\t Harga_Total")
	fmt.Println()
	for i := 0; i < len(nama_barang); i++ {
		fmt.Println(qty[i], "\t", nama_barang[i], "\t", listpembelian[i], "\t\t", listpembelian[i]*qty[i])
	}
	fmt.Println()
	fmt.Println("\n===============")
	ttl := time.Now()
	y, m, d := ttl.Date()
	fmt.Print("Pegawai : ", pegawai, "\nCostumer: ", customer, " \nAlamat  : ", alamat, "\nTotal   : ", hargatotal, "\nTanggal : ", d, m, y)
	fmt.Println("\n===============")
	fmt.Print("\nSudah bayar(y/n): ")
	var bayar string
	fmt.Scanln(&bayar)
	if bayar == `y` || bayar == `Y` {

		for i := 0; i < len(barangID); i++ {
			as.DelNota(id)
		}
		as.DelCustomerNota(id)
		fmt.Println("Terima kasih!")
		fmt.Print("''press enter''")
		fmt.Scanln()
		return
	}
	fmt.Println("belum bayar")
	fmt.Print("''press enter''")
	fmt.Scanln()
}
