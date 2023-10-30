package crud

import (
	"Module/config"
	"fmt"
)

func (as *CrudSystem) FiturPegawai(pegawai string) {
	for {
		config.CallClear()
		fmt.Println("\n===============")
		fmt.Println("Login as:", pegawai)
		fmt.Println("===============")
		fmt.Println("\n===============")
		fmt.Println("Fitur Pegawai")
		fmt.Println("===============")
		input := ""
		fmt.Println("\n1. Menambahkan Barang Baru")
		fmt.Println("2. Edit Informasi Barang")
		fmt.Println("3. Update Stok Barang")
		fmt.Println("\n4. Menambahkan Daftar Customer")
		fmt.Println("\n5. Membuat Nota Transaksi")
		fmt.Println("6. Membuat Transaksi")
		fmt.Println("0. Menu Utama")
		fmt.Print("\n: ")
		fmt.Scanln(&input)
		if input == `1` {
			as.TambahBarang(pegawai)
		} else if input == `2` {
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
			input := ""
			fmt.Println("\n1. Nama")
			fmt.Println("2. Harga")
			fmt.Print("\n: ")
			fmt.Scanln(&input)

			if input == `1` {
				as.EditNamaBarang(pegawai)
			} else if input == `2` {
				as.EditHargaBarang(pegawai)
			} else if input == `3` {

			}

		} else if input == `3` {
			as.EditStokBarang(pegawai)
		} else if input == `4` {
			as.TambahCustomer(pegawai)
		} else if input == `5` {
			as.NotaTransaksi(pegawai)
		} else if input == `6` {
			as.OrderPesanan(pegawai)
		} else if input == `0` {
			break
		}
	}
}
