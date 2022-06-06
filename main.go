package main

import (
	_config "be9/app-project/config"
	_userController "be9/app-project/controllers/user"
	"database/sql"
	"fmt"
)

var DBconn *sql.DB

func init() {
	DBconn = _config.ConnectionDB()
}

func main() {
	fmt.Println("Pilih Menu: (1: Lihat Data)")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		results := _userController.GetAllUser(DBconn)

		for _, v := range results {
			fmt.Println("ID:", v.ID, "Nama", v.Nama, "Telp", v.Telp)
		}

	}
}
