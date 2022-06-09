package deleteaccount

import (
	_entities "be9/app-project/entities"
	"database/sql"
)

func DeleteAccount(db *sql.DB, userDelete _entities.User) (int, error) {
	var query = ("DELETE FROM user WHERE telp = ?")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}

	result := statement.QueryRow(userDelete.Telp)

	var id_user int
	err := result.Scan(&id_user)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}

}
