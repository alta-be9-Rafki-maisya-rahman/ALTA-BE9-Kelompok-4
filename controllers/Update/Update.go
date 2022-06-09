package update

import (
	_entities "be9/app-project/entities"
	"database/sql"
)

func UpdateAccount(db *sql.DB, accountUpdate _entities.User) (int, error) {
	var query = ("UPDATE USER SET user_name = (?), password = (?) WHERE telp = ?")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}

	result, err := statement.Exec(accountUpdate.Nama, accountUpdate.Password, accountUpdate.Telp)

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}
