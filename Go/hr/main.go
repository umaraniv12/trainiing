package main

import (
	"log"

	"github.com/Eswaraa/hr/emp"
	"github.com/Eswaraa/hr/emp/empdb"
)

func main() {
	log.Println("Main")
	var emps []emp.Emp
	emps = empdb.FetchEmployees()
	for _, v := range emps {
		log.Printf("\n Emp: %v", v)
	}
}
