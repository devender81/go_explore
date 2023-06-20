package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Simple Calculator")

	var numbers []float64
	var userInput string

	// Read numbers from the user until they enter "done"
	for {
		fmt.Print("Enter a number (or 'done' to calculate): ")
		fmt.Scanln(&userInput)

		if userInput == "done" {
			break
		}
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}
		numbers = append(numbers, num)
	}

	sum := addition(numbers)
	fmt.Println(sum)

	sub := subtract(numbers)
	fmt.Println(sub)

	mul := multiply(numbers)
	fmt.Println(mul)

	div := divide(numbers)
	fmt.Println(div)

}

// Calculate the sum of the entered numbers
func addition(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func subtract(numbers []float64) float64 {

	sub := numbers[0]
	for i := 1; i < len(numbers); i++ {
		sub -= numbers[i]
	}
	return sub
}

func multiply(numbers []float64) float64 {
	var mul float64
	for _, num := range numbers {
		mul *= num
	}
	return mul
}

func divide(numbers []float64) float64 {
	var div float64
	for _, num := range numbers {
		div /= num
	}
	return div
}

// func subtract() float64 {
// 	var numbers []float64
// 	var userInput string

// 	// Read numbers from the user until they enter "done"
// 	for {
// 		fmt.Print("Enter a number (or 'done' to calculate): ")
// 		fmt.Scanln(&userInput)

// 		if userInput == "done" {
// 			break
// 		}
// 		num, err := strconv.ParseFloat(userInput, 64)
// 		if err != nil {
// 			fmt.Println("Invalid number. Please try again.")
// 			continue
// 		}
// 		numbers = append(numbers, num)
// 	}

// 	// Calculate the subtraction of the entered numbers
// 	var sub float64
// 	//for _, num := range numbers {
// 	//	sub -= num
// 	//}

// 	if len(numbers) > 0 {
// 		sub = numbers[0]
// 		for i := 1; i < len(numbers); i++ {
// 			sub -= numbers[i]
// 		}
// 	} else {
// 		fmt.Println("no number entered")
// 	}
// 	return sub

//}
