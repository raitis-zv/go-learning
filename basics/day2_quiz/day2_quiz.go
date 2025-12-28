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
	var answer1 string
	var check1 int
	var answer2 string
	var check2 int
	var answer3 string
	var check3 int
	var answer4 string
	var check4 int
	var answer5 string
	var check5 int

	var points int

	fmt.Println("What is 6 x 6?")
	for {
		fmt.Scanln(&answer1)
		temp, err := strconv.Atoi(answer1)
		if err != nil || temp < 0 {
			fmt.Println("Please enter a positive whole number")
			continue
		}
		check1 = temp
		break
	}

	if check1 == 36 {
		fmt.Println("Correct!")
		points++
	} else {
		fmt.Println("Incorrect :(")
	}
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()

	fmt.Println("What is 99 + 0 ?")
	for {
		fmt.Scanln(&answer2)
		temp, err := strconv.Atoi(answer2)
		if err != nil || temp < 0 {
			fmt.Println("Please enter a positive whole number")
			continue
		}
		check2 = temp
		break
	}

	if check2 == 99 {
		fmt.Println("Correct!")
		points++
	} else {
		fmt.Println("Incorrect :(")
	}
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()

	fmt.Println("What is 10 / 10?")
	for {
		fmt.Scanln(&answer3)
		temp, err := strconv.Atoi(answer3)
		if err != nil || temp < 0 {
			fmt.Println("Please enter a positive whole number")
			continue
		}
		check3 = temp
		break
	}

	if check3 == 1 {
		fmt.Println("Correct!")
		points++
	} else {
		fmt.Println("Incorrect :(")
	}
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()

	fmt.Println("What is 42 + 1?")
	for {
		fmt.Scanln(&answer4)
		temp, err := strconv.Atoi(answer4)
		if err != nil || temp < 0 {
			fmt.Println("Please enter a positive whole number")
			continue
		}
		check4 = temp
		break
	}

	if check4 == 43 {
		fmt.Println("Correct!")
		points++
	} else {
		fmt.Println("Incorrect :(")
	}
	fmt.Println("Press any button to continue...")
	fmt.Scanln()

	fmt.Println("What is 100 - 67?")
	for {
		fmt.Scanln(&answer5)
		temp, err := strconv.Atoi(answer5)
		if err != nil || temp < 0 {
			fmt.Println("Please enter a positive whole number")
			continue
		}
		check5 = temp
		break
	}

	if check5 == 33 {
		fmt.Println("Correct!")
		points++
	} else {
		fmt.Println("Incorrect :(")
	}
	fmt.Println("Press any button to continue...")
	fmt.Scanln()

	return points
}
