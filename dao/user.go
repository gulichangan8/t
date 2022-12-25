package dao

import "t/model"

func TakeUserDate() model.Users {
	rows, _ := dB.Query("select * from ?", "user")
	var u model.User
	var U model.Users
	for rows.Next() {
		err := rows.Scan(&u.Username, &u.Password)
		if err != nil {
			panic(err)
		}
		U = append(U, u)
	}
	return U
}
