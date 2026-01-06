package main

import "fmt"

func main() {
	/* Only one main function is allowed per package;
	   ignore the error for now, run files individually (not the whole package)
	*/
	age := 23

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are under 18.")
	}
}
