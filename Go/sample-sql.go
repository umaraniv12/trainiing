package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	EmpId int    `json:"emp_id"`
	Name  string `json:"emp_name"`
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/company")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT empno, name FROM emp")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// results, err := db.Query("INSERT INTO emp values(3,'Dana')")

	for results.Next() {
		var emp Emp
		// for each row, scan the result into our tag composite object
		err = results.Scan(&emp.EmpId, &emp.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(emp.Name)
	}

}
