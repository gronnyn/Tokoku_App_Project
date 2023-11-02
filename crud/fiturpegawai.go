package crud

import (
	"Module/config"
	"fmt"
)

func (as *CrudSystem) FiturPegawai(pegawai string) {
	for {
		config.CallClear()
		input := ""
		fmt.Println("\n===============")
		fmt.Println("Login as:", pegawai)
		fmt.Println("===============")
		fmt.Println("\n===============")
		fmt.Println("Fitur Pegawai")
		fmt.Println("===============")
		fmt.Println("\n1. Menambahkan Barang Baru")
		fmt.Println("2. Edit Informasi Barang")
		fmt.Println("3. Update Stok Barang")
		fmt.Println("\n4. Menambahkan Daftar Customer")
		fmt.Println("\n5. Membuat Nota Transaksi")
		fmt.Println("\n6. Membuat Pesanan")
		fmt.Println("7. Tampilkan Data")
		fmt.Println("\n0. Back")
		fmt.Print("\n: ")
		fmt.Scanln(&input)
		if input == `1` {
			as.TambahBarang(pegawai)
		} else if input == `2` {
			for {
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
				fmt.Println("\n0. Back")
				fmt.Print("\n: ")
				fmt.Scanln(&input)

				if input == `0` {
					break
				} else if input == `1` {
					as.EditNamaBarang(pegawai)
				} else if input == `2` {
					as.EditHargaBarang(pegawai)
				} else {
					fmt.Println("Input Salah!")
				}
			}

		} else if input == `3` {
			as.EditStokBarang(pegawai)
		} else if input == `4` {
			as.TambahCustomer(pegawai)
		} else if input == `5` {
			as.NotaTransaksi(pegawai)
		} else if input == `6` {
			as.OrderPesanan(pegawai)
		} else if input == `7` {
			for {
				config.CallClear()
				input := ""
				fmt.Println("\n===============")
				fmt.Println("Login as:", pegawai)
				fmt.Println("===============")
				fmt.Println("\n===============")
				fmt.Println("Fitur Pegawai")
				fmt.Println("===============")
				fmt.Println("\n===============")
				fmt.Println("Tampilkan Data")
				fmt.Println("===============")
				fmt.Println("\n1. Barang")
				fmt.Println("2. Customer")
				fmt.Println("3. Pesanan")
				fmt.Println("\n0. Back")
				fmt.Print("\n: ")
				fmt.Scanln(&input)

				if input == `1` {
					as.ListBarang()
					fmt.Print("'press enter'")
					fmt.Scanln()
				} else if input == `2` {
					as.ListCustomer()
					fmt.Print("'press enter'")
					fmt.Scanln()
				} else if input == `0` {
					break
				} else if input == `3` {
					as.ListNota()
					fmt.Print("'press enter")
					fmt.Scanln()
				} else {
					fmt.Println("Input salah!")
					fmt.Print("'press enter")
					fmt.Scanln()
				}
			}
		} else if input == `0` {
			break
		} else {
			fmt.Println("Input Salah!")
			fmt.Print("'press enter'")
			fmt.Scanln()
		}
	}
}
