package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) NotaTransaksi(pegawai string) {

	var id int
	var Nota []model.Nota_Transactions
	var barangID []int
	var qty []int
	var listpembelian []int
	var currentlist int
	var nama_barang []string
	var currentstring string
	var customer, alamat string
	var hargatotal int
	var bayar string
	var created_at []string
	var currstok int
	var oldstok, newstok []int

	as.DB.Joins("JOIN barangs ON barangs.ID = nota_transactions.barang_id").Where("customer_id = ?", id).Find(&Nota)

	as.ListCustomer()
	fmt.Print("\nID customer: ")
	fmt.Scanln(&id)

	as.DB.Model(model.Nota_Transactions{}).Select("barang_id").Where("customer_id = ?", id).Find(&barangID)
	as.DB.Model(model.Nota_Transactions{}).Select("qty").Where("customer_id = ?", id).Find(&qty)
	as.DB.Model(model.Nota_Transactions{}).Where("customer_id = ?", id).Find(&Nota)
	as.DB.Model(model.Nota_Transactions{}).Select("created_at").Where("customer_id = ?", id).Find(&created_at)
	as.DB.Model(model.Customer{}).Select("nama").Where("ID = ?", id).Find(&customer)
	as.DB.Model(model.Customer{}).Select("alamat_customer").Where("ID = ?", id).Find(&alamat)

	for x := 0; x < len(barangID); x++ {
		as.DB.Model(model.Barang{}).Where("ID = ?", barangID[x]).Select("harga").Take(&currentlist)
		listpembelian = append(listpembelian, currentlist)
	}

	for x := 0; x < len(barangID); x++ {
		as.DB.Model(model.Barang{}).Where("ID = ?", barangID[x]).Select("nama").Take(&currentstring)
		nama_barang = append(nama_barang, currentstring)
	}

	for x := 0; x < len(barangID); x++ {
		as.DB.Model(model.Barang{}).Where("ID = ?", barangID[x]).Select("stok_barang").Take(&currstok)
		oldstok = append(oldstok, currstok)
	}

	for x := 0; x < len(barangID); x++ {
		y := 0
		y = oldstok[x] - qty[x]
		newstok = append(newstok, y)
	}
	fmt.Println(listpembelian)
	for i := 0; i < len(Nota); i++ {
		fmt.Println(Nota[i])
	}

	for i := 0; i < len(qty); i++ {
		hargatotal += qty[i] * listpembelian[i]
	}
	fmt.Println(barangID)
	fmt.Println(qty)
	fmt.Print(nama_barang)
	fmt.Print(customer)
	fmt.Print(created_at)
	fmt.Println(newstok)
	as.DB.Model(model.Barang{}).Select("stok_barang")
	fmt.Println()
	fmt.Println()
	if hargatotal == 0 {
		fmt.Println("sudah lunas")
		return
	}
	fmt.Println("===============")
	fmt.Print("Nama\t: ", customer)
	fmt.Println()
	fmt.Print("Alamat\t: ", alamat)
	fmt.Println()
	fmt.Print("Pegawai\t: ", pegawai)
	fmt.Println("\n===============")
	fmt.Println("\nPesanan\t:")
	fmt.Println()
	fmt.Println("Total\t Barang\t\t Harga_Satuan\t Harga_Total\t Waktu")
	fmt.Println()
	for i := 0; i < len(nama_barang); i++ {
		fmt.Println(qty[i], "\t", nama_barang[i], "\t", listpembelian[i], "\t\t", listpembelian[i]*qty[i], "\t\t", created_at[i])
	}
	fmt.Println()
	fmt.Println("Total\t:\t\t\t\t", hargatotal)

	fmt.Print("Apakah costumer ", customer, "alamat", alamat, "sudah bayar(y/n): ")
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
