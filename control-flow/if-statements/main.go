package main

import (
    "fmt"
    "math/rand"
    "os"
)

/*
Description: Returns an integer divided by 2, rounded down if odd. DO NOT MODIFY THIS CODE.
*/
func roundDown(num int) int {
    return num / 2
}

/*
Description: Generates a random integer between 1 and 100. DO NOT MODIFY THIS CODE.
*/
func generateNumber() int {
    return (rand.Int() % 100) + 1
}

func main() {
    var userInput int
    var randomValue int

    fmt.Println("Please enter a number less than or equal to 100: ")
    fmt.Scanln(&userInput)
    randomValue = generateNumber()

    if /*fill in conditional statement */ {
        fmt.Println("")
        os.Exit(0) // Exit the program due to user giving improper value
    } else if /*fill in conditional statement */ {
        fmt.Println("")
        os.Exit(0) // Exit the program due to user giving improper value
    } else {
        fmt.Println("")
        os.Exit(1) // Exit the program due to something unexpected happening
    }

    /*
    	  Implement code that achieves the following:
			- If the number is even print to the screen "Your number is even." If it isn't even, print "Your number is odd" to the screen instead.
			- If the number is greater than 50, print to the screen "Your number is greater than 50". If it isn't, print "Your number is not greater than 50" instead.
			- If the number is GREATER than 50, is it greater than 75? If it isn't, print to the screen "The number is between 50 and 75." 
			- If the number is also GREATER than 75, is it greater than 90? If it isn't, print to the screen "The number is between 75 and 90."
			- If the number is not greater than 50, is it greater than 25? If it is, print to the screen "The number is between 25 and 50."
			- If the number is also LESS than 25, is it greater than 10? if it is, print to the screen "The number is between 10 and 25."
    */
}
