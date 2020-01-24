package main

import (
  "fmt"
)

func main() {
	//A map literal is a very convenient way to initialize a map with some data. 
	// You just need to pass the key-value pairs separated by colon 
	//inside curly braces like below
	var employees = map[int]string{
		1001: "Rajeev",
		1002: "Sachin",
		1003: "James",
	}

	printEmployee(employees, 1001)
	              // m[key] = value ~ using employees[1001] = Rajeev, to add and retrieve the value
	printEmployee(employees, 1010)

	if isEmployeeExists(employees, 1002) {
		fmt.Println("EmployeeId 1002 found")
	}
}

func printEmployee(employees map[int]string, employeeId int) {
	                  // m				   , [key][KeyType]
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
		                                   //[ValueType], ok
	} else {
		fmt.Printf("EmployeeId %d not found\n", employeeId)
	}
}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}