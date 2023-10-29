package main

import (
	"Module/auth"
	"Module/config"
	"Module/crud"
	"fmt"
)

func main() {
	cfg, err := config.InitDB()

	inputMenu := "0"

	if err != nil {
		fmt.Println("failed connect to database", err.Error())
	}

	var authSystem = auth.AuthSystem{DB: cfg}
	var crudSystem = crud.CrudSystem{DB: cfg}
	for {
		fmt.Printf("\t\tTokoku Project Start")
		fmt.Println("\n\n1. Login")
		fmt.Println("0. Exit")
		fmt.Print("\n: ")
		fmt.Scanln(&inputMenu)

		if inputMenu == `1` {
			fmt.Println()
			x, y := authSystem.Login()
			if y == 1 {
				for {
					input := ""
					fmt.Println("\n===============")
					fmt.Println("Login as:", x)
					fmt.Println("===============")
					fmt.Println("\n1. Menambahkan User")
					fmt.Println("2. Menambahkan Barang")
					fmt.Println("\n3. Menghapus User")
					fmt.Println("4. Menghapus Barang")
					fmt.Println("0. Logout")
					fmt.Print("\n: ")
					fmt.Scanln(&input)
					if input == `0` {
						break
					} else if input == `1` {
						crudSystem.TambahUser()
					} else if input == `2` {
						crudSystem.TambahBarang()
					} else if input == `3` {
						crudSystem.DelUsers()
					} else if input == `4` {
						crudSystem.DelBarang()
					} else if input == `5` {
						crudSystem.DelCustomer()
					}
				}
			} else if y == 2 {
				for {
					input := ""
					fmt.Println("\n===============")
					fmt.Println("Login as: ", x)
					fmt.Println("===============")
					fmt.Println("\n1. Menambahkan Barang")
					fmt.Println("2. Edit Informasi Barang")
					fmt.Println("3. Update Stok Barang")
					fmt.Println("\n4. Menambahkan Daftar Customer")
					fmt.Println("5. Membuat Nota Transaksi")
					fmt.Println("0. Logout")
					fmt.Print("\n: ")
					fmt.Scanln(&input)
					if input == `0` {
						break
					} else if input == `1` {
						crudSystem.TambahBarang()
					} else if input == `2` {
					} else if input == `3` {

					} else if input == `4` {
						crudSystem.TambahCustomer()
					} else if input == `5` {
					} else if input == `6` {
					}
				}
			}
		} else if inputMenu == `0` {
			break
		} else {
			fmt.Println("input salah!")
		}
	}
	fmt.Println("\t\tTokoku Project End")
}
