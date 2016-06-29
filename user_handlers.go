package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	db := DBConnection()

	if err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id); err != nil {
		log.Println("Error: ", err)
		errResponse, err := MakeError(err)
		if err != nil {
			log.Println("Error creating error json response")
		}

		errjson, err := json.Marshal(&errResponse)
		if err != nil {
			log.Println("Can't marshal error to json")
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(errjson)
		return
	}

	if userJSON, err := json.Marshal(&user); err != nil {
		log.Println("Error marshaling users data to json")
		w.WriteHeader(500)
	} else {
		log.Println("Retrieved users:", user)
		w.Write(userJSON)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := DBConnection()
	var users Users

	if err := db.Select(&users, "SELECT * FROM users"); err != nil {
		log.Println("Error:", err)

		errResponse, err := MakeError(err)
		if err != nil {
			log.Println("Error ", err)
		}

		value, err := json.Marshal(&errResponse)
		if err != nil {
			log.Println("Can't marshal error to json")
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(value)
		return
	}

	if usersJSON, err := json.Marshal(&users); err != nil {
		log.Println("Error: users json not valie")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJSON)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := DBConnection()
	vars := mux.Vars(r)
	id := vars["id"]

	tx := db.MustBegin()
	rows := tx.MustExec("DELETE FROM users WHERE id = $1", id)
	res, err := rows.RowsAffected()
	if err != nil {
		log.Println("Err deleting user. Error:", err)
	}
	if res > 0 {
		log.Println("Destroyed user with id:", id)
	} else {
		log.Println("User with that id wasn't found")
		w.WriteHeader(404)
		return
	}

	tx.Commit()
	w.WriteHeader(204)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	id := uuid.New()
	db := DBConnection()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		log.Println("failed", err)
	}

	tx := db.MustBegin()
	_, err := db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3) RETURNING id", id, user.Name, user.Age)

	if err != nil {
		if res, err := MakeError(err); err != nil {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(422)
			w.Write(res)
		}
		return
	}

	tx.Commit()

	user.SetID(id)
	responseJSON, err := json.Marshal(user)
	if err != nil {
		log.Println("Couldn't marshal response json for user")
		w.WriteHeader(500)
		return
	}
	log.Println("Response Data: " + string(responseJSON))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	if err := json.NewEncoder(w).Encode(responseJSON); err != nil {
		log.Println("Error responsing with user json: ", err.Error())
	}
}
