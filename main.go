package main

import (
	_config "be9/app-project/config"
	_topupbalanceController "be9/app-project/controllers/topupbalance"

	// _transferController "be9/app-project/controllers/transfer"
	// _userbalanceController "be9/app-project/controllers/userbalance"
	_entities "be9/app-project/entities"
	"database/sql"
	"fmt"
)

var DBconn *sql.DB

func init() {
	DBconn = _config.ConnectionDB()
}

func main() {
	//Top Up
	newTopUpBalance := _entities.TopUpBalance{}
	fmt.Println("Input ID USER:")
	fmt.Scanln(&newTopUpBalance.UserId)
	fmt.Println("Input Nominal:")
	fmt.Scanln(&newTopUpBalance.NominalTopUp)

	row, err := _topupbalanceController.CreateTopUpBalance(DBconn, newTopUpBalance)
	if err != nil {
		fmt.Println("error insert", err.Error())
	} else {
		fmt.Println("Insert Success")
		fmt.Println("row affect", row)
	}

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

	// //User Balance
	// var inputUserBalance _entities.UserBalance
	// fmt.Println("Input ID User:")
	// fmt.Scanln(&inputUserBalance.IdUser)
	// fmt.Println("Input Saldo Top Up:")
	// fmt.Scanln(&inputUserBalance.Saldo)

	// row2, err := _userbalanceController.UpdateUserBalance(DBconn, inputUserBalance)
	// if err != nil {
	// 	fmt.Println("error update", err.Error())
	// } else {
	// 	fmt.Println("update Success")
	// 	fmt.Println("row affect", row2)
	// }

}
