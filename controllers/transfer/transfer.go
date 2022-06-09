package transfer

import (
	_entities "be9/app-project/entities"
	"database/sql"
	"fmt"
)

func CreateTransfer(db *sql.DB, newtransfer _entities.Transfer) (int, error) {
	var querybalance = (`SELECT user_balance.total_balance FROM user_balance
		INNER JOIN user ON user.telp = user_balance.telp
		INNER JOIN transfer ON transfer.transfer_user = user.telp where transfer.transfer_user = (?)`)
	balance := db.QueryRow(querybalance, newtransfer.TransferUser)
	var databalance int
	err := balance.Scan(&databalance)

	if err != nil {
		return 0, err
	}

	if newtransfer.NominalTransfer < databalance {
		var query = ("INSERT INTO transfer (transfer_user, transfer_receiver, nominal_transfer) VALUES(?, ?, ?)")
		statement, errPrepare := db.Prepare(query)
		if errPrepare != nil {
			return 0, errPrepare
		}
		_, err := statement.Exec(newtransfer.TransferUser, newtransfer.TransferReceiver, newtransfer.NominalTransfer)
		if err != nil {
			return 0, err
		}
	}

	var kueri = (`SELECT user_balance.total_balance FROM user_balance
									INNER JOIN user ON user.telp = user_balance.telp
									INNER JOIN transfer ON transfer.transfer_receiver = user.telp where transfer.transfer_receiver = (?)`)
	transReceive := db.QueryRow(kueri, newtransfer.TransferReceiver)
	var dataReceive int
	err3 := transReceive.Scan(&dataReceive)

	if err3 != nil {
		return 0, err3
	}

	var xueri = (`SELECT user_balance.total_balance FROM user_balance INNER JOIN user ON user.telp = user_balance.telp INNER JOIN transfer ON transfer.transfer_user = user.telp where transfer.transfer_user = (?)`)
	transUser := db.QueryRow(xueri, newtransfer.TransferUser)
	var dataTransUser int
	errx := transUser.Scan(&dataTransUser)

	if errx != nil {
		return 0, errx
	}

	if dataTransUser > newtransfer.NominalTransfer {
		lastbalance1 := dataReceive + newtransfer.NominalTransfer
		var kuerix = (`UPDATE user_balance SET total_balance= (?) WHERE telp=(?)`)
		statement1, errPrepare1 := db.Prepare(kuerix)
		if errPrepare1 != nil {
			return 0, errPrepare1
		}
		resulti, erri := statement1.Exec(lastbalance1, newtransfer.TransferReceiver)

		lastbalance := dataTransUser - newtransfer.NominalTransfer
		var kuerij = (`UPDATE user_balance SET total_balance= (?) WHERE telp=(?)`)
		statementts, errPreparees := db.Prepare(kuerij)
		if errPreparees != nil {
			return 0, errPreparees
		}
		resultj, _ := statementts.Exec(lastbalance, newtransfer.TransferUser)
		if resultj != nil {
			rowj, _ := resultj.RowsAffected()
			return int(rowj), nil
		}

		if erri != nil {
			return 0, erri
		} else {
			rowi, _ := resulti.RowsAffected()
			return int(rowi), nil
		}
	} else {
		fmt.Println("Saldo Tidak Cukup")
	}
	return 0, fmt.Errorf("")
}
