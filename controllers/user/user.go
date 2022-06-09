package user

import (
	_entities "be9/app-project/entities"
	"database/sql"
	"fmt"
)

func AddAccount(db *sql.DB, newAccount _entities.User) (int, error) {

	if newAccount.Nama == "" || newAccount.Password == "" || newAccount.Telp == "" {
		return 0, fmt.Errorf(" Data tidak boleh Kosong")
	}
	var query = ("INSERT INTO user(user_name, telp, password) VALUES(?,?,?)")
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}

	result, err := statement.Exec(newAccount.Nama, newAccount.Telp, newAccount.Password)

	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

func SearchUser(db *sql.DB, telp string) []_entities.User {
	querysearch := ("SELECT id, user_name, telp FROM user WHERE telp=(?)")
	defer db.Close()
	dataUser, err := db.Query(querysearch, telp)
	if err != nil {
		fmt.Println(err.Error())
	}

	var dataPengguna []_entities.User
	for dataUser.Next() {
		var dataPengguna1 _entities.User
		err := dataUser.Scan(&dataPengguna1.ID, &dataPengguna1.Nama, &dataPengguna1.Telp)
		if err != nil {
			fmt.Println("error scan", err.Error())
		}

		dataPengguna = append(dataPengguna, dataPengguna1)
	}
	return dataPengguna
}

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
