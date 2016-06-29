package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pborman/uuid"
)

func getUserQuery(db *sqlx.DB, user User, id string, w http.ResponseWriter) {
	if err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}

// GetUser is the handler func for `GET /users/{id}`
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	db := DBConnection()

	getUserQuery(db, user, id, w)
	if userJSON, err := json.Marshal(&user); err != nil {
		fmt.Println("Error marshaling users data to json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Println("Retrieved users:", user)
		w.Write(userJSON)
	}
}

// GetUsers is the handler func for `GET /users`
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := DBConnection()
	var users Users

	if err := db.Select(&users, "SELECT * FROM users"); err != nil {
		MakeError(err, w)
		return
	}

	if usersJSON, err := json.Marshal(&users); err != nil {
		fmt.Println("Error: users json not valie")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJSON)
	}
}

// DeleteUser is the handler func for `DELETE /users/{id}`
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := DBConnection()
	vars := mux.Vars(r)
	id := vars["id"]

	tx := db.MustBegin()
	rows := tx.MustExec("DELETE FROM users WHERE id = $1", id)
	res, err := rows.RowsAffected()
	if err != nil {
		fmt.Println("Err deleting user. Error:", err)
	}
	if res > 0 {
		fmt.Println("Destroyed user with id:", id)
	} else {
		fmt.Println("User with that id wasn't found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusNoContent)
}

// CreateUser is the handler func for `POST /users`
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	id := uuid.New()
	db := DBConnection()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		fmt.Println("failed", err)
	}

	tx := db.MustBegin()
	_, err := db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3) RETURNING id", id, user.Name, user.Age)

	if err != nil {
		MakeError(err, w)
		return
	}

	tx.Commit()

	user.SetID(id)
	responseJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Couldn't marshal response json for user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Response Data: " + string(responseJSON))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(responseJSON); err != nil {
		fmt.Println("Error responsing with user json: ", err.Error())
	}
}
