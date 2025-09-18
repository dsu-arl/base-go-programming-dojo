package main

import (
	"fmt"
	"math/rand"
)

/*
Generates hints based on the provited hintLevel. DO NOT MODIFY THIS CODE
*/
func hintGenerator(number, guess, hintLevel int) {
	switch hintLevel {
	case 0:
		if number >= 50 {
			fmt.Println("The number is greater than or equal to 50")
		} else {
			fmt.Println("The number is less than 50.")
		}
	case 1:
		if 0 == (number % 2) {
			fmt.Println("The number is even.")
		} else {
			fmt.Println("The number is odd.")
		}
	case 2:
		tensValue := number / 10
		valueOne := tensValue * 10
		valueTwo := valueOne + 9
		fmt.Println("The number is between", valueOne, "and", valueTwo)
	default:
		if guess > number {
			fmt.Println("You've guessed too high")
		} else {
			fmt.Println("You've guessed too low")
		}
	}
}

func main() {
	randValue := rand.Int() % 101
	var userGuess int
	var userHint string

	fmt.Println("Guess a number between 0 and 100")
	fmt.Scanln(&userGuess)

	/* ~*~*~*~*~*~*~*~* MODIFY CODE BELOW THIS LINE ONLY ~*~*~*~*~*~*~*~* */

	if userGuess == randValue {

	} else {

	}
}
