package topupbalance

import (
	_entities "be9/app-project/entities"
	"database/sql"
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
		var queryselect = (`SELECT user_balance.total_balance, user_balance.id_user FROM user_balance INNER JOIN user ON user.id = user_balance.id_user WHERE user.telp = ?`)
		var data1, data2 int
		dataSaldo := db.QueryRow(queryselect, newTopUpBalance.Telp)
		err := dataSaldo.Scan(&data1, &data2)

		if err != nil {
			return 0, err
		} else {
			balanceUpdate := data1 + newTopUpBalance.NominalTopUp
			var query = (`UPDATE user_balance SET total_balance= (?) WHERE id_user=(?)`)
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
		// row, _ := results.RowsAffected()
		// return int(row), nil
	}

}
