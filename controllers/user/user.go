package user

import (
	_entities "be9/app-project/entities"
	"database/sql"
	"fmt"
)

func GetAllUser(db *sql.DB) []_entities.User {
	results, err := db.Query("SELECT id, user_name, telp, password FROM user")
	if err != nil {
		fmt.Println("error", err.Error())
	}
	defer db.Close()

	var userAll []_entities.User
	for results.Next() {
		var user _entities.User
		err := results.Scan(&user.ID, &user.Nama, &user.Telp, &user.Password)
		if err != nil {
			fmt.Println("error scan", err.Error())
		}
		userAll = append(userAll, user)
	}
	return userAll
}
