package crud

import (
	"Module/config"
	"Module/model"
	"bufio"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)

type CrudSystem struct {
	DB *gorm.DB
}

func (as *CrudSystem) TambahBarang(pegawai string) {
	config.CallClear()
	fmt.Println("\n===============")
	fmt.Println("Login as:", pegawai)
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println(" Fitur Pegawai")
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println(" Menambahkan Barang Baru")
	fmt.Println("===============")
	as.ListBarang()
	fmt.Println("\n===============")
	fmt.Println("Tambah Barang Baru")
	fmt.Println("===============")
	var newUser = new(model.Barang)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama: ")
	inputnama, _ := reader.ReadString('\n')
	inputnama = strings.TrimSpace(inputnama)
	newUser.Nama = inputnama

	if newUser.Nama == "" {
		fmt.Println("Nama tidak boleh kosong!")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}
	fmt.Print("Harga: ")
	fmt.Scanln(&newUser.Harga)
	if newUser.Harga == 0 {
		fmt.Println("Harga tidak boleh 0!")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}
	fmt.Print("Stok: ")
	fmt.Scanln(&newUser.Stok_Barang)
	if newUser.Stok_Barang == 0 {
		fmt.Println("Stok tidak boleh 0!")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}
	newUser.Pegawai = pegawai

	err := as.DB.Table("barangs").Create(newUser).Error

	if err != nil {
		fmt.Println("gagal membuat barang baru")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	} else {
		as.ListBarang()
		fmt.Println("\nBarang " + newUser.Nama + " berhasil dibuat.")
		fmt.Print("'press enter'")
		fmt.Scanln()
	}
}

func (as *CrudSystem) TambahCustomer(pegawai string) {
	config.CallClear()
	fmt.Println("\n===============")
	fmt.Println("Login as:", pegawai)
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Fitur Pegawai")
	fmt.Println("===============")
	fmt.Println("\n===============")
	fmt.Println("Menambahkan Customer")
	fmt.Println("===============")
	as.ListCustomer()
	reader := bufio.NewReader(os.Stdin)
	var newUser = new(model.Customer)
	fmt.Print("\nNama Customer: ")
	inputnama, _ := reader.ReadString('\n')
	inputnama = strings.TrimSpace(inputnama)
	newUser.Nama = inputnama
	if newUser.Nama == "" {
		fmt.Println("Nama tidak boleh kosong!")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}
	fmt.Print("Alamat Customer: ")
	inputalamat, _ := reader.ReadString('\n')
	inputalamat = strings.TrimSpace(inputalamat)
	newUser.Alamat_Customer = inputalamat
	if newUser.Alamat_Customer == "" {
		fmt.Println("Alamat tidak boleh kosong!")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	err := as.DB.Table("customers").Create(newUser).Error

	if err != nil {
		fmt.Println("gagal")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}

	as.DB.Table("customers_backup").Create(newUser)
	as.ListCustomer()
	fmt.Println("\nCustomer " + newUser.Nama + " berhasil ditambahkan.")
	fmt.Print("'press enter'")
	fmt.Scanln()

}

func (as *CrudSystem) TambahPegawai() {
	fmt.Println("\n===============")
	fmt.Println("Pegawai Baru")
	fmt.Println("===============")
	as.ListPegawai()
	var newUser = new(model.User)
	fmt.Print("Username: ")
	fmt.Scanln(&newUser.Username)
	if newUser.Username == "" {
		fmt.Println("username tidak boleh kosong!")
		fmt.Print("'press enter'")
		fmt.Scanln()
		return
	}
	fmt.Print("Password: ")
	fmt.Scanln(&newUser.Password)

	err := as.DB.Table("users").Create(newUser).Error

	if err != nil {
		fmt.Println("Pegawai sudah ada")
		fmt.Print("'press enter'")
		fmt.Scanln()
	} else {
		as.ListPegawai()
		fmt.Println("\nPegawai", newUser.Username, "berhasil ditambah.")
		fmt.Print("'press enter'")
		fmt.Scanln()
	}
}
