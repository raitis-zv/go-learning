package main

import "fmt"

func main() {
	ages := []int{10, 15, 18, 23, 40, 65}

	for index, age := range ages {
		personNum := index + 1
		status := classifyAge(age)

		fmt.Printf("The person #%d is %s, age: %d\n", personNum, status, age)
	}
}

func classifyAge(age int) string {
	if age >= 65 {
		return "a senior"
	} else if age < 13 {
		return "a child"
	} else if age < 18 {
		return "a teenager"
	} else {
		return "an adult"
	}
}
