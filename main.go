package main

import (
	_config "be9/app-project/config"
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
	fmt.Println("Pilih Menu: (1: Register Account)")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		newAccount := _entities.User{}
		fmt.Println("Input ID")
		fmt.Scanln(&newAccount.ID)
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

	}

}
