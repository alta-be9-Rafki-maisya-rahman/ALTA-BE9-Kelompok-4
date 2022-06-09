package main

import (
	_config "be9/app-project/config"
	_topUpBalanceController "be9/app-project/controllers/topupbalance"
	_transferController "be9/app-project/controllers/transfer"
	_user "be9/app-project/controllers/user"

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
	fmt.Println("Pilih menu berikut: (1:Top Up) / // (2:Cek History Top Up) / (3: Lihat Pengguna Lain) / (4 : Transfer) / (5 : Cek History Transfer) / (0 : Keluar)")
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
		fmt.Println("Input 0 untuk keluar")
		fmt.Scanln(&pilihan)

	case 2:
		var cekTopUp1 string
		fmt.Println("input Telp:")
		fmt.Scanln(&cekTopUp1)

		result := _topUpBalanceController.GetHistoryTopUp(DBconn, cekTopUp1)
		for _, v := range result {
			fmt.Println("ID:", v.ID, "Nominal Top Up:", v.NominalTopUp, "Tanggal:", v.Tanggal, "Telp:", v.Telp)
		}
		fmt.Println("Input 0 untuk keluar")
		fmt.Scanln(&pilihan)
	case 3:
		searchOtherUser := ""
		fmt.Print("Masukkan Telp:")
		fmt.Scanln(&searchOtherUser)

		row := _user.SearchUser(DBconn, searchOtherUser)
		if len(row) != 0 {
			for _, v := range row {
				fmt.Println("ID:", v.ID, "Nama:", v.Nama, "Telp:", v.Telp)
			}
		} else {
			fmt.Println(" Pengguna Tidak Ada")
		}
		fmt.Println("Input 0 untuk keluar")
		fmt.Scanln(&pilihan)

	case 4:
		newtransfer := _entities.Transfer{}
		fmt.Print("Masukkan Telp Pengirim:")
		fmt.Scanln(&newtransfer.TransferUser)
		fmt.Print("Masukkan Telp Penerima:")
		fmt.Scanln(&newtransfer.TransferReceiver)
		fmt.Print("Masukkan Nominal Transfer:")
		fmt.Scanln(&newtransfer.NominalTransfer)

		row, err := _transferController.CreateTransfer(DBconn, newtransfer)
		if err != nil {
			fmt.Println("error insert", err.Error())
		} else {
			fmt.Println("Insert Success")
			fmt.Println("row affect", row)
		}
		fmt.Println("Input 0 untuk keluar")
		fmt.Scanln(&pilihan)
	case 5:
		historyTransfer := _entities.Transfer{}
		fmt.Print("Masukkan Telp:")
		fmt.Scanln(&historyTransfer.TransferUser)

		results := _transferController.GetDataTransfer(DBconn, historyTransfer)
		for _, v := range results {
			fmt.Println("Nominal Transfer", v.NominalTransfer, "No.Pengirim", v.TransferUser, "No.Penerima", v.TransferReceiver, "Tanggal:", v.Tanggal)
		}
		fmt.Println("Input 0 untuk keluar")
		fmt.Scanln(&pilihan)
	}

	if pilihan == 0 {
		fmt.Println("Terima Kasih Telah Bertransaki")
	}
}
