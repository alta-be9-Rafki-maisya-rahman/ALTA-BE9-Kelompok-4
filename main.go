package main

import (
	_config "be9/app-project/config"
	_updateUser "be9/app-project/controllers/Update_User"
	_login "be9/app-project/controllers/login"
	_controllers "be9/app-project/controllers/user"
	_entities "be9/app-project/entities"
	"database/sql"
	"fmt"
)

var DBconn *sql.DB

func init() {
	DBconn = _config.ConnectionDB()
}

func main() {
	fmt.Println("Pilih Menu: (1: Register Account) (2: Login)")
	var pilihan int
	fmt.Scanln(&pilihan)
	defer DBconn.Close()

	switch pilihan {
	case 1:
		newAccount := _entities.User{}

		fmt.Println("Input NAMA")
		fmt.Scanln(&newAccount.Nama)
		fmt.Println("Input Password")
		fmt.Scanln(&newAccount.Password)
		fmt.Println("Input Telp")
		fmt.Scanln(&newAccount.Telp)

		row, err := _controllers.AddAccount(DBconn, newAccount)

		if err != nil {
			fmt.Println("registrasi gagal", err.Error())
		} else {
			fmt.Println("registrasi berhasil")
			fmt.Println("row affect", row)
		}

	case 2:
		loginUser := _entities.User{}
		fmt.Println("Input Telp")
		fmt.Scanln(&loginUser.Telp)
		fmt.Println("Input Password")
		fmt.Scanln(&loginUser.Password)

		row, err := _login.UserLogin(DBconn, loginUser)

		if err != nil {
			fmt.Println("login Berhasil", err.Error(), "data tidak cocok")
		} else {
			fmt.Println("login Sukses")
			fmt.Println("row affect", row)

			fmt.Println("  WALLET USER MENU ")
			fmt.Println(" (1 UPDATE DATA) (2 CARI KONTAK) (3 !! HAPUS AKUN ANDA !!)")
			var userMenu int
			fmt.Scanln(&userMenu)
			switch userMenu {
			case 1:
				updateUser := _entities.User{}
				fmt.Println("input no telp baru")
				fmt.Scanln(&updateUser.Telp)
				fmt.Println("input user name baru")
				fmt.Scanln(&updateUser.Nama)
				fmt.Println("input password baru")
				fmt.Scanln(&updateUser.Password)
				fmt.Println("input Telp saat ini")
				fmt.Scanln(&updateUser.Telp)

				row, err := _updateUser.UpdateUser(DBconn, updateUser)
				if err != nil {
					fmt.Println("update gagal", err.Error())
				} else {
					fmt.Println("update berhasil")
					fmt.Println("row affect", row)
				}

			}
		}

	}

}
