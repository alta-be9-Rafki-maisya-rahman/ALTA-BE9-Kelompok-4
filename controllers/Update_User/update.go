package update_user

import (
	_entities "be9/app-project/entities"
	"database/sql"
)

func UpdateUser(db *sql.DB, updateAccount _entities.User) (int, error) {
	var query = ("update user set telp = ?, user_name=?, password = ? where telp=? ")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}

	result, err := statement.Exec(updateAccount.Nama, updateAccount.Telp, updateAccount.Password, updateAccount.Telp)

	defer db.Close()

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}
