package main

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"f_Name"`
	LastName   string `json:"l_Name"`
	BudgetID   int    `json:"budget_id"`
	GoalItemID int    `json:"goal_item_id"`
}

func createUser(usr User) int {

	sqlStateusrnt := `INSERT INTO users (email, f_name, l_name, budget_id, goal_item_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;`

	err := db.QueryRow(sqlStateusrnt, usr.Email, usr.FirstName, usr.LastName, usr.BudgetID, usr.GoalItemID).Scan(&usr.ID)
	if err != nil {
		log.Println(err)
	}

	return usr.ID
}

func updateUser(usr User) (User, error) {
	sqlStmt := `UPDATE users
				SET f_name = Nul$2, l_name = $3, budget_id = $4, goal_item_id = $5
				WHERE email = $1
				RETURNING *;`

	err := db.QueryRow(sqlStmt, usr.Email, usr.FirstName, usr.LastName, usr.BudgetID, usr.GoalItemID).Scan(&usr.ID, &usr.Email, &usr.FirstName, &usr.LastName, &usr.BudgetID, &usr.GoalItemID)
	if err != nil {
		log.Println(err)
		return usr, err
	}

	// count, err := res.RowsAffected()
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Printf("The %v rows have been effected", count)

	return usr, nil
}

func deleteUser(id int) {
	stmt := `DELETE FROM users
		WHERE ID = $1;`

	_, err := db.Exec(stmt, id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The user has been deleted!")

}

func getUser(email int) User {
	var user User
	stmt := `SELECT id, email, f_name, l_name, budget_id, goal_item_id FROM users 
			WHERE email=$1;`
	row := db.QueryRow(stmt, email)
	err := row.Scan(&user.ID, &user.Email, &user.ID, &user.FirstName, &user.LastName, &user.BudgetID, &user.GoalItemID)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		return user
	default:
		panic(err)
	}

	return user
}
