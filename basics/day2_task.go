package main

import (
	"fmt"
	"strconv"
)

func main() {
	var seniorCount int
	var childrenCount int
	var teenCount int
	var adultCount int
	var maaCount int

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

		if status == "a senior" {
			seniorCount++
		} else if status == "a child" {
			childrenCount++
		} else if status == "a teenager" {
			teenCount++
		} else if status == "an adult" {
			adultCount++
		} else {
			maaCount++
		}

		fmt.Printf("The person #%d is a %d-year-old %s\n", personNum, age, status)
	}

	fmt.Println("___ SUMMARY ___")
	fmt.Printf("Total people: %d\n", n)
	fmt.Printf("Children: %d\n", childrenCount)
	fmt.Printf("Teenagers %d\n", teenCount)
	fmt.Printf("Adults: %d\n", adultCount)
	fmt.Printf("Middle-aged adults: %d\n", maaCount)
	fmt.Printf("Seniors: %d\n", seniorCount)
}

func classifyAge(age int) string {
	if age >= 65 {
		return "a senior"
	} else if age >= 40 && age < 65 {
		return "a middle-aged adult"
	} else if age < 13 {
		return "a child"
	} else if age < 18 {
		return "a teenager"
	} else {
		return "an adult"
	}
}
