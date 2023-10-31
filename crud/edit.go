package crud

import (
	"Module/config"
	"Module/model"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (as *CrudSystem) EditNamaBarang(pegawai string) {
	config.CallClear()
	fmt.Println("\n===============")
	fmt.Println("Login as:", pegawai)
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Fitur Pegawai")
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Edit Informasi Barang")
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Edit Nama Barang")
	fmt.Println("===============")
	as.ListBarang()
	var currentUser = new(model.Barang)
	reader := bufio.NewReader(os.Stdin)
	var kode int

	fmt.Print("\nKode Barang: ")
	fmt.Scanln(&kode)

	qry := as.DB.Table("barangs").Where("ID = ?", kode).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nBarang tidak ditemukan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	fmt.Print(currentUser.Nama, " > ")
	inputnama, _ := reader.ReadString('\n')
	inputnama = strings.TrimSpace(inputnama)

	as.DB.Table("barangs").Where("ID = ?", kode).Update("nama", inputnama)

	as.ListBarang()
	fmt.Println("Nama barang dengan kode", kode, "berhasil diubah menjadi", inputnama)
	fmt.Print("'press enter'")
	fmt.Scanln()
}

func (as *CrudSystem) EditHargaBarang(pegawai string) {
	config.CallClear()
	fmt.Println("\n===============")
	fmt.Println("Login as:", pegawai)
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Fitur Pegawai")
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Edit Informasi Barang")
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Edit Nama Barang")
	fmt.Println("===============")
	as.ListBarang()
	var currentUser = new(model.Barang)
	var kode, harga int

	fmt.Print("\nKode Barang: ")
	fmt.Scanln(&kode)

	qry := as.DB.Table("barangs").Where("ID = ?", kode).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nBarang tidak ditemukan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	fmt.Print("Harga: ")
	fmt.Scanln(&harga)
	as.DB.Table("barangs").Where("ID = ?", kode).Update("harga", harga)

	fmt.Println()
	config.CallClear()
	as.ListBarang()
	fmt.Print("'press enter'")
	fmt.Scanln()
}

func (as *CrudSystem) EditStokBarang(pegawai string) {
	config.CallClear()
	fmt.Println("\n===============")
	fmt.Println("Login as:", pegawai)
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Fitur Pegawai")
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Update Stok Barang")
	fmt.Println("===============")
	as.ListBarang()
	var currentUser = new(model.Barang)
	var kode, stok int

	fmt.Print("\nKode Barang: ")
	fmt.Scanln(&kode)

	qry := as.DB.Table("barangs").Where("ID = ?", kode).Take(&currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nBarang tidak ditemukan.")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	fmt.Print(currentUser.Stok_Barang, " > ")
	fmt.Scanln(&stok)
	as.DB.Table("barangs").Where("ID = ?", kode).Update("stok_barang", stok)

	fmt.Println()
	config.CallClear()
	as.ListBarang()
	fmt.Println("Stok barang dengan kode", kode, "berhasil diubah dari ", currentUser.Stok_Barang, "menjadi", stok)
	fmt.Print("'press enter'")
	fmt.Scanln()
}
