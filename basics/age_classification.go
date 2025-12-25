package main

import "fmt"

func main() {
	age := 23

	if age < 13 {
		fmt.Println("Child")
	} else if age < 18 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Adult")
	}
}
