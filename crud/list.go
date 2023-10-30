package crud

import (
	"Module/model"
	"fmt"
)

func (as *CrudSystem) ListBarang() {
	fmt.Println("\n===============")
	fmt.Println("List Barang")
	fmt.Println("===============")
	var barang []model.Barang
	var id, harga, stok []int
	var nama, pegawai []string
	as.DB.Find(&barang)
	as.DB.Model(model.Barang{}).Select("ID").Find(&id)
	as.DB.Model(model.Barang{}).Select("nama").Find(&nama)
	as.DB.Model(model.Barang{}).Select("harga").Find(&harga)
	as.DB.Model(model.Barang{}).Select("stok_barang").Find(&stok)
	as.DB.Model(model.Barang{}).Select("pegawai").Find(&pegawai)

	fmt.Print("\nID , Nama Barang , Harga Stok , Barang")
	fmt.Println()
	fmt.Println()
	for i := 0; i < len(barang); i++ {
		fmt.Println(id[i], ",", nama[i], ",", harga[i], ",", stok[i], "added by", pegawai[i])
	}
}

func (as *CrudSystem) ListCustomer() {
	fmt.Println("\n===============")
	fmt.Println("List Costumer")
	fmt.Println("===============")
	var customer []model.Customer
	var id, nama, alamat []string
	as.DB.Find(&customer)
	as.DB.Model(model.Customer{}).Select("nama").Find(&nama)
	as.DB.Model(model.Customer{}).Select("alamat_customer").Find(&alamat)
	as.DB.Model(model.Customer{}).Select("ID").Find(&id)

	fmt.Print("\nID , Nama , Alamat")
	fmt.Println()
	for i := 0; i < len(customer); i++ {
		fmt.Println(id[i], ",", nama[i], ",", alamat[i])
	}
}

func (as *CrudSystem) ListCustomerNota(cust string) {
	var customer []model.Customer
	// fmt.Print(nota)
	as.DB.Find(&customer).Where("ID = nota_transactions.customer_id")
	// as.DB.Model(model.Nota_Transactions{}).Select("barang_id").Where("customer_id = ?", id).Find(&cust)
	fmt.Print(customer)
	fmt.Print("\n{ID, Nama, Alamat}")
	fmt.Println()
	fmt.Println()
	for i := 0; i < len(customer); i++ {
		fmt.Println(customer[i])
	}
}

func (as *CrudSystem) ListUser() {
	var user []model.User
	var listuser []string
	as.DB.Select("username").Find(&user)
	fmt.Println()
	fmt.Println("\n===============")
	fmt.Println(" List Pegawai")
	fmt.Println("===============")
	as.DB.Model(model.User{}).Select("username").Find(&listuser)
	for i := 0; i < len(listuser); i++ {
		fmt.Println(listuser[i])
	}
}

func (as *CrudSystem) ListNota() {
	var nota []model.Nota_Transactions

	as.DB.Find(&nota)
	fmt.Print("\n{ID, username, ID Barang, Total Barang, ID Customer}")
	fmt.Println()
	fmt.Println()
	for i := 0; i < len(nota); i++ {
		fmt.Println(nota[i])
	}
}
