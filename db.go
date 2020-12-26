package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	sqlStatement := `INSERT INTO users (fname, lname, budgetid, goalitemid)
	VALUES ($1, $2, $3, $4);`

	err := db.QueryRow(sqlStatement)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("The user has been created")

}
