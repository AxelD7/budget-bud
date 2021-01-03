package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error

	pgInfo := fmt.Sprintf("host=localhost port=%v user=postgres password=%s dbname=%s sslmode=disable", PostgresPort, PostgresPW, PostgresDb)

	db, err = sql.Open("postgres", pgInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("WE FUCKED UP")
		panic(err)
	}

	fmt.Println("Successfully connected to database!")

	// Handlers
	http.HandleFunc("/", root)

	http.HandleFunc("/c/usr", dbTest)

	fmt.Printf("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	defer db.Close()
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root page!!")
}
