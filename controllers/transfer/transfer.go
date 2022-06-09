package transfer

import (
	"database/sql"
	_entities "be9/app-project/entities"
)

func CreateTransfer(db *sql.DB, newTransfer _entities.Transfer) (int, error) {
	var query =(`INSERT INTO transfer`)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	results, err := statement.Exec(newTransfer.TransferReceiver, newTransfer.TransferUser, newTransfer.NominalTransfer)

	defer db.Close()

	if err != nil {
		return 0, err
	} else {
		row, _ := results.RowsAffected()
		return int(row), nil
	}
}