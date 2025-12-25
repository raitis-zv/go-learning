package main

import "fmt"

func main() {
	age := 23
	var status string

	if age < 13 {
		status = "Child"
	} else if age < 18 {
		status = "Teenager"
	} else {
		status = "Adult"
	}

	fmt.Println("You are an", status, ".")
}
