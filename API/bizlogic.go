package library

import (
	"database/sql"
	"fmt"
	"log"
	model1 "qw/model"
)

func GetUsersDB(db *sql.DB, users []model1.Users) []model1.Users {
	rows, err := db.Query(`SELECT user_id, email, password, username, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model1.Users
		err := rows.Scan(&user.User_id, &user.Email, &user.Username, &user.Password, &user.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

func DelUsersDB(db *sql.DB, user_id int) {
	fmt.Println("---in Db---")
	_, err := db.Query("Delete from users where user_id = ?", user_id)
	if err != nil {
		log.Fatal(err)
		fmt.Println("------In err-------------")
	}
	fmt.Println("Deleted from users table successfully")
}

func GetUserDB(db *sql.DB, users []model1.Users, user_id int) {
	fmt.Println("----In DB--------")
	var user model1.Users
	query := `select user_id, email, password, username, created_at FROM users where user_id = ?`
	fmt.Println(query)
	err := db.QueryRow(query, user_id).Scan(&user.User_id, &user.Email, &user.Username, &user.Password, &user.Created_at)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Got a particular users  from table successfully")
}

func AddUsersDB(db *sql.DB, users model1.Users) {
	query := `INSERT INTO users (user_id, email, password, username,created_at)
              VALUES (?, ?, ?, ?, ?)`
	fmt.Println(query)
	_, err := db.Exec(query, users.User_id, users.Username, users.Email, users.Password, users.Created_at)
	fmt.Println(users.User_id, users.Username, users.Email, users.Password, users.Created_at)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Addd new user to table successfully")
}

func UpdateUsersDB(db *sql.DB, users model1.Users) {
	query := `UPDATE users SET user_id = ? , username = ?, password = ?, email = ?, created_at = ? WHERE user_id = ?`
	fmt.Println(query)
	// fmt.Println(users.User_id, users.Username, users.Email, users.Password, users.Created_at)
	u, err := db.Exec(query, users.User_id, users.Username, users.Email, users.Password, users.Created_at) // check err
	fmt.Println(u)
	if err != nil {
		log.Fatal(err)
		fmt.Println("----IN DB ERROR-----")
	}
	fmt.Println("Record updated successfully.")
}
