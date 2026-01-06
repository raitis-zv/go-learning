package main

import (
	"fmt"
	"strconv"
)

func main() {

	name := []string{"Jane", "Peter", "Raul", "Paul"}
	goal := []int{1800, 2200, 2500, 2100}
	totalWater := []int{0, 0, 0, 0}

	for i := 0; i < 4; i++ {
		addWater(name[i], &totalWater[i])
	}

	for i := 0; i < 3; i++ {
		printSummary(name[i], goal[i], totalWater[i])
		seeNext(name[i+1])
	}

	printSummary(name[3], goal[3], totalWater[3])

	println("__________")
	leaderboard(totalWater, name)
	println("")
}

func addWater(name string, water *int) {
	var waterInput string

	for {
		fmt.Println("Enter water amount for ", name, " (ml):")
		fmt.Scanln(&waterInput)
		temp, err := strconv.Atoi(waterInput)

		if err != nil || temp < 0 {
			fmt.Println("Please enter a positive whole number")
			continue
		} else {
			fmt.Println("Water intake for", name, "has been updated!")
			*water += temp
			fmt.Println("")
			fmt.Println("Press Enter to continue...")
			waitForEnter()
		}
		break
	}
}

func printSummary(name string, goal int, totalWater int) {
	fmt.Println(name+"'s total water intake for today is:", totalWater, "ml")
	fmt.Println("")
	fmt.Println("Daily goal:", goal, "ml")
	fmt.Println("")
	fmt.Println("Progress:", totalWater, "ml /", goal, "ml")
	fmt.Println(goal-totalWater, "ml left to hit the target")
	fmt.Println("")
}

func waitForEnter() {
	fmt.Scanln()
}

func seeNext(nextName string) {
	fmt.Println("Press Enter to see the stats of", nextName)
	waitForEnter()
}

func leaderboard(total []int, names []string) {
	maxIndex := 0
	for i := 1; i < len(total); i++ {
		if total[i] > total[maxIndex] {
			maxIndex = i
		}
	}
	fmt.Println("Top drinker today:", names[maxIndex])
}
