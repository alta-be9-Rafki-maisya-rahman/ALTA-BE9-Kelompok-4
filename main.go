package main

import (
	_config "be9/app-project/config"
	_login "be9/app-project/controllers/login"
	_topUpBalanceController "be9/app-project/controllers/topupbalance"
	_transferController "be9/app-project/controllers/transfer"
	_controllers "be9/app-project/controllers/user"
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
			fmt.Println("1. Read account /n 2. Update account /n 3. Delete account /n 4. top-up /n 5. transfer /n 6. history top-up /n 7. history transfer /n 8. lihat profil lain /n 0. keluar")

			fmt.Scanln(&userMenu)
			switch userMenu {
			case 1:
				//read account

			case 2:
				//update account

			case 3:
				//delete account

			case 4:
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

			case 5:
				//transfer
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

			case 6:
				//history top-up
				var cekTopUp1 string
				fmt.Println("input Telp:")
				fmt.Scanln(&cekTopUp1)

				result := _topUpBalanceController.GetHistoryTopUp(DBconn, cekTopUp1)
				for _, v := range result {
					fmt.Println("ID:", v.ID, "Nominal Top Up:", v.NominalTopUp, "Tanggal:", v.Tanggal, "Telp:", v.Telp)
				}

			case 7:
				//history transfer

			case 8:
				//lihat profil user lain
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

			case 0:
				//keluar & ucapan terima kasih

			}
		}
	}
}
