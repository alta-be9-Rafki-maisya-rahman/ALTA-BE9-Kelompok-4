package userbalance

import (
	_entities "be9/app-project/entities"
	"database/sql"
)

func InputUserBalance(db *sql.DB, inputUserBalance _entities.UserBalance) (int, error) {

	var query = (`INSERT INTO user_balance (id_user, total_balance) VALUES(?, ?) `)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	result, err := statement.Exec(inputUserBalance.IdUser, inputUserBalance.Saldo)

	defer db.Close()

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}
