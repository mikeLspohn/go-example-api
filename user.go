package main

import (
	"encoding/json"
	"log"
	"os"
)

func (u *User) MarshalUserJSON() {
	if responseJSON, err := json.Marshal(u); err != nil {
		log.Println("Couldn't marshal response json for user")
	} else {
		log.Println("Response Date: " + string(responseJSON))
	}
}

type User struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  string `db:"age" json:"age"`
}

type UserErrors struct {
	Message string
}

type Users []User

func (u *User) SetID(id string) {
	u.ID = id
}

func MakeError(err error) ([]byte, error) {
	errJSON := UserErrors{Message: err.Error()}
	jsonResponse, err := json.Marshal(errJSON)
	if err != nil {
		log.Println("Error marshaling json for user errors", err)
	} else {
		log.Println("Error: ", errJSON.Message)
		os.Stdout.Write(jsonResponse)
	}

	return jsonResponse, err
}
