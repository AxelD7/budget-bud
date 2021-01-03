package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func dbTest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var usr User
	err := decoder.Decode(&usr)
	fmt.Println("User is decoding")
	if err != nil {
		panic(err)
	}

	fmt.Printf("!!!!!The user is %#v", usr)

	user, err := updateUser(usr)
	if err != nil {
		fmt.Printf("error updating user - %s", err)
		return
	}

	fmt.Printf("the user has been updated - id %v", user)
	w.WriteHeader(http.StatusOK)
}
