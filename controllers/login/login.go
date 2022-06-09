package login

import (
	_entities "be9/app-project/entities"
	"database/sql"
)

func UserLogin(db *sql.DB, loginUser _entities.User) (int, error) {
	var query = ("SELECT id FROM user WHERE telp = ? and password = ?")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}

	result := statement.QueryRow(loginUser.Telp, loginUser.Password)

	var id_user int
	err := result.Scan(&id_user)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}

}
