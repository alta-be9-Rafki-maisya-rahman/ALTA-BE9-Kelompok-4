package user

import (
	_entities "be9/app-project/entities"
	"database/sql"
)

func AddAccount(db *sql.DB, newAccount _entities.User) (int, error) {
	var query = ("INSERT INTO user(id, user_name, telp, password) VALUES(?,?,?,?)")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}

	result, err := statement.Exec(newAccount.ID, newAccount.Nama, newAccount.Telp, newAccount.Password)

	defer db.Close()

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}
