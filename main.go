package main

import (
	_config "be9/app-project/config"
	_topUpBalanceController "be9/app-project/controllers/topupbalance"

	// _transferController "be9/app-project/controllers/transfer"

	// _entities "be9/app-project/entities"
	"database/sql"
	"fmt"
)

var DBconn *sql.DB

func init() {
	DBconn = _config.ConnectionDB()
}

func main() {
	var pilihan int
	fmt.Println("Pilih menu berikut: (1:Top Up) / // (2:Cek History Top Up) / (3: Transfer) / (4: Cek History Transfer) / (5 : Keluar)")
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		//Feature Top Up
		// newTopUpBalance := _entities.TopUpBalance{}
		// fmt.Println("Input Telp:")
		// fmt.Scanln(&newTopUpBalance.Telp)
		// fmt.Println("Input Nominal:")
		// fmt.Scanln(&newTopUpBalance.NominalTopUp)

		// row, err := _topUpBalanceController.CreateTopUpBalance(DBconn, newTopUpBalance)
		// if err != nil {
		// 	fmt.Println("error insert", err.Error())
		// } else {
		// 	fmt.Println("Insert Success")
		// 	fmt.Println("row affect", row)
		// }
	case 2:
		var cekTopUp string
		fmt.Println("input Telp:")
		fmt.Scanln(&cekTopUp)

		result := _topUpBalanceController.GetHistoryTopUp(DBconn, cekTopUp)
		for _, v := range result {
			fmt.Println("ID:", v.ID, "Nominal Top Up:", v.NominalTopUp, "Tanggal:", v.Tanggal, "Telp:", v.Telp)
		}
	}

}
