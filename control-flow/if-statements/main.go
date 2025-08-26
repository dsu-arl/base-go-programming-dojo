package main

import (
    "fmt"
    "os"
	"strconv"
)

/*
Description: Returns true if the integer is even, and false if odd. DO NOT MODIFY THIS CODE.
*/
func isEven(num int) bool {

	retValue = false

	if 0 == num%2 {
		retValue = true
	}

	return retValue
}

func main() {
	var userInput int

	if len(os.Args) == 2 {
		userInput, _ = strconv.Atoi(os.Args[1])
	} else {
		fmt.Println("Please enter a number less than or equal to 100: ")
		fmt.Scanln(&userInput)
	}

	/* ~*~*~*~*~*~*~*~* MODIFY CODE BELOW THIS LINE ONLY ~*~*~*~*~*~*~*~* */

    if /*fill in conditional statement */ {
        fmt.Println("") // Fill in the required text
        os.Exit(0) // Exit the program due to user giving improper value
    } else if /*fill in conditional statement */ {
        fmt.Println("") // Fill in the required text
        os.Exit(0) // Exit the program due to user giving improper value
    } else if /*fill in conditional statement */ {
		fmt.Println("") // Fill in the required text
		os.Exit(0) // Exit the program due to user giving improper value	
	} else {
        fmt.Println("") // Fill in the required text
    }

    /*
		Implement code that achieves the following:
		- If the number is even, print to the screen "Your number is even." If it isn't even, print "Your number is odd" to the screen instead.
		- If the number is GREATER THAN 75, but LESS THAN 90, print to the screen "The number is between 75 and 90." 
			Otherwise, if the number GREATER THAN 50 but LESS THAN 75, print to the screen "Your number is between 50 and 75." 
			Otherwise print "Your number is less than 50."
		- If the number is GREATER THAN 10, but LESS THAN 25, print "The number is between 10 and 25." 
			Otherwise, if the number is GREATER THAN 25 but LESS THAN 50, print "The number is between 25 and 50." 
			Otherwise print "You have a weird number..."
    */
}
