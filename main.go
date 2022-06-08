package main

import (
	_config "be9/app-project/config"
	_topUpBalanceController "be9/app-project/controllers/topupbalance"

	// _transferController "be9/app-project/controllers/transfer"

	_entities "be9/app-project/entities"
	"database/sql"
	"fmt"
)

var DBconn *sql.DB

func init() {
	DBconn = _config.ConnectionDB()
}

func main() {
	var pilihan int
	fmt.Println("Pilih menu berikut: (1:Top Up) / (2:Transfer) / (3: Cek Saldo) / (4: Lihat History Top Up) / (5 : Lihat History Transfer)")
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		//Feature Top Up
		newTopUpBalance := _entities.TopUpBalance{}
		fmt.Println("Input Telp:")
		fmt.Scanln(&newTopUpBalance.Telp)
		fmt.Println("Input Nominal:")
		fmt.Scanln(&newTopUpBalance.NominalTopUp)

		row, err := _topUpBalanceController.CreateTopUpBalance(DBconn, newTopUpBalance)
		if err != nil {
			fmt.Println("error insert", err.Error())
		} else {
			fmt.Println("Insert Success")
			fmt.Println("row affect", row)
		}

		// var ubahUserBalance _entities.UserBalance
		// fmt.Println("Input ID User:")
		// fmt.Scanln(&ubahUserBalance.IdUser)
		// fmt.Println("Input Saldo Top Up/Transfer:")
		// fmt.Scanln(&ubahUserBalance.Saldo)

		// row2_1, err := _userBalanceController.UpdateUserBalance(DBconn, ubahUserBalance)
		// if err != nil {
		// 	fmt.Println("error update", err.Error())
		// } else {
		// 	fmt.Println("update Success")
		// 	fmt.Println("row affect", row2_1)
		// }

		// case 2:
		//Transfer
		// newTransfer := _entities.Transfer{}
		// fmt.Println("Input ID Transfer Receiver:")
		// fmt.Scanln(&newTransfer.TransferReceiver)
		// fmt.Println("Input ID User Sent Transfer :")
		// fmt.Scanln(&newTransfer.TransferUser)
		// fmt.Println("Input Nominal Transfer :")
		// fmt.Scanln(&newTransfer.NominalTransfer)

		// row1, err := _transferController.CreateTransfer(DBconn, newTransfer)
		// if err != nil {
		// 	fmt.Println("error insert", err.Error())
		// } else {
		// 	fmt.Println("Insert Success")
		// 	fmt.Println("row affect", row1)
		// }
		// case 3:
		// 	var iduser int
		// 	fmt.Print("Masukkan id user:")
		// 	fmt.Scanln(&iduser)
		// 	row, err := _userBalanceController.GetUserBalance(DBconn, iduser)

		// 	if err != nil {
		// 		fmt.Println("error get", err.Error())
		// 	} else {
		// 		fmt.Println("delete Succes")
		// 		fmt.Println("row affect", row)
		// 	}
	}

}
