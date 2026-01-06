package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Grade float64
}

func main() {
	var studentCount int

	studentCount = sCount()
	students := make([]Student, studentCount)

	for i := 0; i < studentCount; i++ {
		newStudent(&students[i])
	}

	for i := 0; i < studentCount; i++ {
		printSummary(students[i])
	}

	avgGrade(students)
}

func sCount() int {
	var studentCount int
	var userInput string

	for {
		fmt.Println("How many students do you want to add?")
		fmt.Scanln(&userInput)

		value, valid := validation(userInput, "int")
		if !valid {
			println("Please enter a valid number of students!")
			continue
		}

		studentCount = int(value)

		break
	}

	return studentCount

}

func newStudent(s *Student) {
	var nameInput string
	var ageInput string
	var gradeInput string

	fmt.Println("Press Enter to create a new student...")
	fmt.Scanln()

	fmt.Println("Enter the name of the student:")
	fmt.Scanln(&nameInput)
	s.Name = nameInput
	fmt.Println("")

	for {
		fmt.Println("Enter the age of the student:")
		fmt.Scanln(&ageInput)

		value, valid := validation(ageInput, "int")
		if !valid {
			fmt.Println("Please enter a valid age!")
			continue
		}

		s.Age = int(value)
		fmt.Println("")

		break
	}

	for {
		fmt.Println("Enter the grade of the student:")
		fmt.Scanln(&gradeInput)

		value, valid := validation(gradeInput, "float")
		if !valid {
			fmt.Println("Please enter a valid grade (1.0 - 10.0)!")
			continue
		}

		s.Grade = value
		fmt.Println("")

		break
	}
}

func printSummary(s Student) {
	fmt.Println("")
	fmt.Println("Student:", s.Name)
	fmt.Println("Age: ", s.Age)
	fmt.Println("Grade: ", s.Grade)
}

func avgGrade(students []Student) {
	var sum float64

	for _, s := range students {
		sum += s.Grade
	}

	avgPrint := sum / float64(len(students))
	fmt.Println("")
	fmt.Print("The average grade of the students: ")
	fmt.Printf("%.1f", avgPrint)
}

func validation(input string, checkType string) (float64, bool) {
	if checkType == "float" {
		value, err := strconv.ParseFloat(input, 64)
		if err != nil || value < 1.0 || value > 10.0 {
			return 0, false
		} else {
			return value, true
		}
	} else {
		value, err := strconv.Atoi(input)
		if err != nil || value <= 0 {
			return 0, false
		} else {
			return float64(value), true
		}
	}
}
