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

	fmt.Print("ID , Nama Barang , Harga Stok , Barang")
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
	var id, nama, alamat []string

	as.DB.Model(model.Customer{}).Select("nama").Find(&nama)
	as.DB.Model(model.Customer{}).Select("alamat_customer").Find(&alamat)
	as.DB.Model(model.Customer{}).Select("ID").Find(&id)

	fmt.Print("ID , Nama , Alamat")
	fmt.Println()
	fmt.Println()
	for i := 0; i < len(id); i++ {
		fmt.Println(id[i], ",", nama[i], ",", alamat[i])
	}
}

func (as *CrudSystem) ListCustomrNota(cust string) {
	var customer []model.Customer
	as.DB.Find(&customer).Where("ID = nota_transactions.customer_id")
	fmt.Print(customer)
	fmt.Print("\n{ID, Nama, Alamat}")
	fmt.Println()
	fmt.Println()
	for i := 0; i < len(customer); i++ {
		fmt.Println(customer[i])
	}
}

func (as *CrudSystem) ListPegawai() {
	var user []model.User
	var listuser []string
	as.DB.Select("username").Find(&user)
	fmt.Println("\n===============")
	fmt.Println("List Pegawai")
	fmt.Println("===============")
	fmt.Println()
	as.DB.Model(model.User{}).Select("username").Find(&listuser)
	for i := 0; i < len(listuser); i++ {
		fmt.Println(listuser[i])
	}
}

func (as *CrudSystem) ListNota() {
	var nota []model.Nota_Transactions
	as.DB.Find(&nota)

	fmt.Print("\n{username, ID Barang, Total Barang, ID Customer}")
	fmt.Println()
	fmt.Println()
	var username []string
	var barang_id, qty, customer_id []int
	as.DB.Model(model.Nota_Transactions{}).Select("username").Find(&username)
	as.DB.Model(model.Nota_Transactions{}).Select("barang_id").Find(&barang_id)
	as.DB.Model(model.Nota_Transactions{}).Select("qty").Find(&qty)
	as.DB.Model(model.Nota_Transactions{}).Select("customer_id").Find(&customer_id)
	for i := 0; i < len(nota); i++ {
		fmt.Println(username[i], ",", barang_id[i], ",", qty[i], ",", customer_id[i])
	}
}
