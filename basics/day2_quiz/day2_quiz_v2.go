package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("___ MATH QUIZ ___\n")
	fmt.Println("Rules:")
	fmt.Println("1. Get 1 point for each correct answer")
	fmt.Println("2. Enter only whole numbers for the answers")
	fmt.Println("3. One try per question\n")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()

	var totalPoints = questionCheck()
	var incorrectAnsw = 5 - totalPoints
	var totalScore = (5 - incorrectAnsw) * 20

	fmt.Printf("Your score: %d%%\n", totalScore)
	fmt.Printf("Correct answers: %v\n", totalPoints)
	fmt.Printf("Incorrect answers: %v\n", incorrectAnsw)
}

func questionCheck() int {
	var points int
	var answer string
	var userAnswer int

	questions := []string{
		"What is 6 x 6?",
		"What is 99 + 0 ?",
		"What is 10 / 10?",
		"What is 42 + 1?",
		"What is 100 - 67?"}

	answers := []int{36, 99, 1, 43, 33}

	for i, q := range questions {
		for {
			fmt.Printf("Question %d: %s\n", i+1, q)
			fmt.Scanln(&answer)
			temp, err := strconv.Atoi(answer)

			if err != nil || temp < 0 {
				fmt.Println("Please enter a positive whole number")
				continue
			}
			userAnswer = temp
			break
		}

		if userAnswer == answers[i] {
			fmt.Println("Correct!")
			points++
		} else {
			fmt.Println("Incorrect :(")
		}

		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
	}

	return points
}
