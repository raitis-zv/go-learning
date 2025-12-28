package main

import "fmt"

func main() {
	ages := []int{10, 15, 18, 23, 40, 65}

	for index, age := range ages {
		personNum := index + 1
		var status string

		if age >= 65 {
			status = "a senior"
		} else if age < 13 {
			status = "a child"
		} else if age < 18 {
			status = "a teenager"
		} else {
			status = "an adult"
		}

		fmt.Printf("The person #%d is %s, age: %d\n", personNum, status, age)
	}
}
