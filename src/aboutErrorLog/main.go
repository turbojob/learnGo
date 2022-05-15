package main

import (
	"fmt"
	"log"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func f1() {
	employee, err := getInformation(1001)
	if err != nil {
		// Something is wrong. Do something.
	} else {
		fmt.Print(employee)
	}
}
func main() {

	log.SetPrefix("from main():")
	log.Fatal("Hey, I'm an error log!")
	log.Panicln("A panic")
	fmt.Print("Can you see me?")
}

func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(1000)
	if err != nil {
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}
	return employee, nil
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
