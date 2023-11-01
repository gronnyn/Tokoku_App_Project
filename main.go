package main

import (
	"Module/auth"
	"Module/config"
	"Module/crud"
	"fmt"
)

func main() {
	cfg, err := config.InitDB()

	inputMenu := ""

	if err != nil {
		fmt.Println("failed connect to database", err.Error())
	}

	var authSystem = auth.AuthSystem{DB: cfg}
	var crudSystem = crud.CrudSystem{DB: cfg}
	for {
		config.CallClear()
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
					fmt.Println("\n1. Menambahkan Pegawai")
					fmt.Println("2. Hapus Data")
					fmt.Println("3. Fitur Pegawai")
					fmt.Println("\n0. Logout")
					fmt.Print("\n: ")
					fmt.Scanln(&input)
					if input == `0` {
						break
					} else if input == `1` {
						config.CallClear()
						fmt.Println("\n===============")
						fmt.Println("Login as:", x)
						fmt.Println("===============")
						crudSystem.TambahPegawai()
					} else if input == `2` {
						config.CallClear()
						input := ""
						fmt.Println("\n===============")
						fmt.Println("Login as:", x)
						fmt.Println("===============")
						fmt.Println("\n===============")
						fmt.Println("Hapus Data")
						fmt.Println("===============")
						fmt.Println("\n1. Pegawai")
						fmt.Println("2. Barang")
						fmt.Println("3. Customer")
						fmt.Print("\n: ")
						fmt.Scanln(&input)
						if input == `1` {
							crudSystem.ListPegawai()
							crudSystem.DelUsers()
						} else if input == `2` {
							crudSystem.ListBarang()
							crudSystem.DelBarang()
						} else if input == `3` {
							crudSystem.ListCustomer()
							crudSystem.DelCustomer()
						} else {
							fmt.Print("\nInput Salah!")
							fmt.Print("\n'press enter")
							fmt.Scanln()
						}
					} else if input == `3` {
						crudSystem.FiturPegawai(x)
					} else {
						fmt.Println("Input Salah!")
						fmt.Print("''press enter''")
						fmt.Scanln()
					}
					config.CallClear()
				}
			} else if y == 2 {
				for {
					input := ""
					fmt.Println("\n===============")
					fmt.Println("Login as:", x)
					fmt.Println("===============")
					fmt.Println("\n1. Fitur Pegawai")
					fmt.Println("\n0. Logout")
					fmt.Print("\n: ")
					fmt.Scanln(&input)
					if input == `1` {
						crudSystem.FiturPegawai(x)
					} else if input == `0` {
						break
					} else {
						fmt.Println("Input Salah!")
						fmt.Print("''press enter''")
						fmt.Scanln()
					}
					config.CallClear()
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
