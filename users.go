package main

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"fName"`
	LastName   string `json:"lName"`
	BudgetID   int    `json:"budgetid"`
	GoalItemID int    `json:"goalitemid"`
}
