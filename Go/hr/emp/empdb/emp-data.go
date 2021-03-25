package empdb

import (
	"database/sql"
	"log"

	"github.com/Eswaraa/hr/emp"
	_ "github.com/go-sql-driver/mysql"
)

func FetchEmployees() []emp.Emp {
	employees := make([]emp.Emp, 0)
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")

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
		var emp emp.Emp
		// for each row, scan the result into our tag composite object
		err = results.Scan(&emp.EmpId, &emp.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(emp.Name)
		employees = append(employees, emp)
	}
	return employees

}
