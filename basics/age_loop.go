package main

import "fmt"

func main() {
	ages := []int{10, 15, 23, 40}

	for _, age := range ages {
		var status string

		if age < 13 {
			status = "Child"
		} else if age < 18 {
			status = "Teenager"
		} else {
			status = "Adult"
		}

		fmt.Printf("Age %d -> %s \n", age, status)
	}
}
