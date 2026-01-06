package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	for {
		fmt.Print("Enter number of people: ")
		var personCount string
		fmt.Scanln(&personCount)

		temp, err := strconv.Atoi(personCount)
		if err != nil || temp <= 0 {
			fmt.Println("Error: enter a positive whole number")
			continue
		}
		n = temp
		break
	}

	ages := make([]int, n)

	for i := 0; i < n; i++ {
		for {
			fmt.Printf("Enter age for person #%d: ", i+1)
			var input string
			fmt.Scanln(&input)

			personAge, err := strconv.Atoi(input)
			if err != nil || personAge <= 0 {
				fmt.Println("Error: enter a positive whole number")
				continue
			}

			ages[i] = personAge
			break
		}
	}

	for i, age := range ages {
		personNum := i + 1

		var status string
		status = classifyAge(age)

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
