package userbalance

import (
	_entities "be9/app-project/entities"
	"database/sql"
)

func UpdateUserBalance(db *sql.DB, ubahUserBalance _entities.UserBalance) (int, error) {

	var query = (`UPDATE user_balance SET total_balance= (?) WHERE id_user=(?)`)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	results, err := statement.Exec(ubahUserBalance.Saldo, ubahUserBalance.IdUser)

	defer db.Close()

	if err != nil {
		return 0, err
	} else {
		row, _ := results.RowsAffected()
		return int(row), nil
	}
}

// func GetUserBalance(db *sql.DB, iduser int) ([]_entities.UserBalance, error) {
// 	var query = fmt.Sprintf("SELECT id_user, total_balance FROM user_balance WHERE id_user=(?)")
// 	statement, errPrepare := db.Prepare(query)
// 	if errPrepare != nil {
// 		return []entities.UserBalance, errPrepare
// 	}
// 	result, err := statement.Exec(iduser)
// 	defer db.Close()

// 	if err != nil {
// 		return 0, err
// 	} else {
// 		row, _ := result.RowsAffected()
// 		return int(row), nil
// 	}
// }
