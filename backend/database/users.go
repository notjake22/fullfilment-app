package database

import "log"

func GetAllUsers() (Users, error) {
	rows, err := DBConn.Query(`SELECT id, "name", phone, balance FROM "user-table"`)
	if err != nil {
		return Users{}, err
	}
	defer rows.Close()

	var users Users
	var user User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Balance)
		if err != nil {
			return users, nil
		}
		users.Users = append(users.Users, user)
	}

	return users, nil
}

func GetUserByID(id string) (*User, error) {
	row := DBConn.QueryRow(`SELECT "name", "phone", "balance" FROM "user-table" WHERE "id" = $1`, id)
	if row.Err() != nil {
		return &User{}, row.Err()
	}

	var user User
	err := row.Scan(&user.Name, &user.Phone, &user.Balance)
	if err != nil {
		return &User{}, err
	}
	user.ID = id

	return &user, nil
}

func SetNewUser(user User) error {
	log.Printf("User Submitted: %v\n", user)
	_, err := DBConn.Exec(`insert into "user-table" ("name", "phone", "balance") values($1, $2, $3)`, user.Name, user.Phone, user.Balance)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func EditUser(user User) error {
	_, err := DBConn.Exec(`UPDATE "user-table" SET "name" = $1, "phone" = $2, "balance" = $3 WHERE id = $4`, user.Name, user.Phone, user.Balance, user.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
