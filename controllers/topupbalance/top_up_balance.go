package topupbalance

import (
	_entities "be9/app-project/entities"
	"database/sql"
)

func CreateTopUpBalance(db *sql.DB, newTopUpBalance _entities.TopUpBalance) (int, error) {
	var query = (`INSERT INTO top_up_balance (id_user, nominal_top_up) VALUES (?, ?)`)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	results, err := statement.Exec(newTopUpBalance.UserId, newTopUpBalance.NominalTopUp)

	// defer db.Close()

	if err != nil {
		return 0, err
	} else {
		row, _ := results.RowsAffected()
		return int(row), nil
	}
}
