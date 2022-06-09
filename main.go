package main

import (
	_config "be9/app-project/config"
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
			fmt.Println("registration failed", err.Error())
		} else {
			fmt.Println("registration succes")
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
			fmt.Println("login failed", err.Error(), "data tidak cocok")
		} else {
			fmt.Println("login succes")
			fmt.Println("row affect", row)
			var userMenu int
			fmt.Scanln(&userMenu)
			switch userMenu {
			case 1:

			}
		}

	}

}
