package user

import (
	_entities "be9/app-project/entities"
	"database/sql"
	"fmt"
)

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
