package main

import (
	"fmt"
	"strconv"
)

func waitInput() {
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
}

func main() {
	fmt.Println("___ MATH QUIZ ___\n")
	fmt.Println("Rules:")
	fmt.Println("1. Get 1 point for each correct answer")
	fmt.Println("2. Enter only whole numbers for the answers")
	fmt.Println("3. One try per question\n")
	waitInput()

	var totalPoints = questionCheck()

	fmt.Printf("Your score: %d%%\n", totalPoints*20)
	fmt.Printf("Correct answers: %v\n", totalPoints)
	fmt.Printf("Incorrect answers: %v\n", 5-totalPoints)
}

func questionCheck() int {
	var points int
	var userAnswer string

	qOperations := []string{"6 x 6", "99 + 0", "10 / 10", "42 + 1", "100 - 67"}

	answers := []string{"36", "99", "1", "43", "33"}

	for i := range qOperations {
		for {
			fmt.Printf("Question %d: What is %s?\n", i+1, qOperations[i])
			fmt.Scanln(&userAnswer)
			temp, err := strconv.Atoi(userAnswer)

			if err != nil || temp < 0 {
				fmt.Println("Please enter a positive whole number")
				continue
			} else if userAnswer == answers[i] {
				fmt.Println("Correct!")
				points++
			} else {
				fmt.Println("Incorrect :(")
			}
			break
		}
		waitInput()
	}

	return points
}
