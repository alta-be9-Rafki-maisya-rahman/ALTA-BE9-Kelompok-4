package readprofile

import (
	_entities "be9/app-project/entities"
	"database/sql"
	"fmt"
)

func GetDataAccount(db *sql.DB, readAccount _entities.User) _entities.User {
	var query = ("SELECT user_name, telp, password, created_at FROM user WHERE telp=(?)")
	dataAccount := db.QueryRow(query, readAccount.Telp)

	var AccountKamu _entities.User
	err := dataAccount.Scan(&AccountKamu.Nama, &AccountKamu.Telp, &AccountKamu.Password, &AccountKamu.Tanggal)

	if err != nil {
		fmt.Println(err.Error())
	}
	return AccountKamu
}
