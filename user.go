package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User is a model to create and fetch users from the db
type User struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  string `db:"age" json:"age"`
}

// UserErrors is used to hold and display user errors
type UserErrors struct {
	Message string
}

// Users is a collection of User
type Users []User

// MarshalUserJSON marshals a user struct into json
func (u *User) MarshalUserJSON() {
	if responseJSON, err := json.Marshal(u); err != nil {
		fmt.Println("Couldn't marshal response json for user")
	} else {
		fmt.Println("Response Date: " + string(responseJSON))
	}
}

// SetID sets a given users id to the passed in string (uuid)
func (u *User) SetID(id string) {
	u.ID = id
}

// MakeError creates a UserErrors struct and marshals it to json
// then writes the appropriate response
func MakeError(errObj error, w http.ResponseWriter) {
	errJSON := UserErrors{Message: errObj.Error()}
	_, err := json.Marshal(errJSON)
	if err != nil {
		fmt.Println("Error marshaling json for user errors", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Println("Error JSON: ", errJSON.Message)
		w.WriteHeader(422)
	}
}
