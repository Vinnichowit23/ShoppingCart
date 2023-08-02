package library

import (
	"encoding/json"
	"fmt"
	"net/http"
	db "qw/db/flyway"
	model1 "qw/model"
	"strconv"

	"github.com/gorilla/mux"
)

var users []model1.Users

func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := db.ConnectToDB()
	users = GetUsersDB(db, users)
	fmt.Println(users)
	json.NewEncoder(w).Encode(users)
}

func deleteusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for index, item := range users {
		fmt.Println(item.User_id)
		fmt.Println(params["id"])
		if item.User_id == id {
			users = append(users[:index], users[index+1:]...)
			fmt.Println("Delete from Database")
			db := db.ConnectToDB()
			DelUsersDB(db, item.User_id)
			break
		}
	}
	fmt.Println(users)
	json.NewEncoder(w).Encode(users)
}

func getuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, item := range users {
		if item.User_id == id {
			db := db.ConnectToDB()
			GetUserDB(db, users, item.User_id)
			json.NewEncoder(w).Encode(item)
		}
	}
}

func addusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model1.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := db.ConnectToDB()
	AddUsersDB(db, user)
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

// func updateusers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println(users)
// 	for index, item := range users {
// 		if item.User_id == id {
// 			users = append(users[:index], users[index+1:]...)
// 			var user model1.Users
// 			err := json.NewDecoder(r.Body).Decode(&user)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusBadRequest)
// 				return
// 			}
// 			db := db.ConnectToDB()
// 			UpdateUsersDB(db, user)
// 			users = append(users, user)
// 			json.NewEncoder(w).Encode(users)
// 		}
// 	}
// }

func updateusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Find the index of the user with the given ID in the users slice
	index := -1
	for i, item := range users {
		if item.User_id == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the updated user from the request body
	var updatedUser model1.Users
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the user in the database
	db := db.ConnectToDB()
	// Update the user in the users slice
	users[index] = updatedUser
	fmt.Println("Values", users)
	UpdateUsersDB(db, updatedUser)

	// Send the updated users slice as the response
	json.NewEncoder(w).Encode(users)
}
