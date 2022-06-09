package topupbalance

import (
	_entities "be9/app-project/entities"
	"database/sql"
	"fmt"
)

func CreateTopUpBalance(db *sql.DB, newTopUpBalance _entities.TopUpBalance) (int, error) {
	var query = (`INSERT INTO top_up_balance (telp, nominal_top_up) VALUES (?, ?)`)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	_, err := statement.Exec(newTopUpBalance.Telp, newTopUpBalance.NominalTopUp)

	defer db.Close()

	if err != nil {
		return 0, err
	} else {
		var queryselect = (`SELECT user_balance.total_balance, user_balance.telp FROM user_balance 
		INNER JOIN user ON user.telp = user_balance.telp WHERE user.telp = ?`)
		var data1, data2 int
		dataSaldo := db.QueryRow(queryselect, newTopUpBalance.Telp)
		err := dataSaldo.Scan(&data1, &data2)

		if err != nil {
			return 0, err
		} else {
			balanceUpdate := data1 + newTopUpBalance.NominalTopUp
			var query = (`UPDATE user_balance SET total_balance= (?) WHERE telp=(?)`)
			statement, errPrepare := db.Prepare(query)
			if errPrepare != nil {
				return 0, errPrepare
			}
			results, err := statement.Exec(balanceUpdate, data2)

			if err != nil {
				return 0, err
			} else {
				row, _ := results.RowsAffected()
				return int(row), nil
			}

		}
	}

}

func GetHistoryTopUp(db *sql.DB, cekTopUp string) []_entities.TopUpBalance {
	query1 := ("SELECT id, nominal_top_up, telp, created_at FROM top_up_balance WHERE telp = (?) ORDER BY created_at DESC")
	results, err := db.Query(query1, cekTopUp)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	defer db.Close()

	var dataTopUpAll []_entities.TopUpBalance
	for results.Next() {
		var dataTopUp _entities.TopUpBalance
		err := results.Scan(&dataTopUp.ID, &dataTopUp.NominalTopUp, &dataTopUp.Telp, &dataTopUp.Tanggal)
		if err != nil {
			fmt.Println("error scan", err.Error())
		}
		dataTopUpAll = append(dataTopUpAll, dataTopUp)
	}
	return dataTopUpAll
}
