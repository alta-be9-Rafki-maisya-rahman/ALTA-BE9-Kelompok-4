package userbalance

import (
	"database/sql"
	_entities "be9/app-project/entities"
)

func UpdateUserBalance(db *sql.DB, inputUserBalance _entities.UserBalance) (int, error) {
	var query = (`UPDATE user_balance SET id_user= (?), total_balance= (?) WHERE id_user=(?)`)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	result, err := statement.Exec(inputUserBalance.IdUser, inputUserBalance.Saldo, inputUserBalance.IdUser)

	defer db.Close()

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}