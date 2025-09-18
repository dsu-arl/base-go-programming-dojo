package main

import (
	"math/rand"
	"os"
	"strconv"
)

/*
Generates and returns a random number. DO NOT MODIFY THIS CODE
*/
func genRandomNumber() int {
	return rand.Int() % 101
}

func main() {
	randomNumber := 0

	if len(os.Args) == 2 {
		randomNumber, _ = strconv.Atoi(os.Args[1])
	} else {
		randomNumber = genRandomNumber() // Use this variable in the random number challenge.
	}

	/* ~*~*~*~*~*~*~*~* MODIFY CODE BELOW THIS LINE ONLY ~*~*~*~*~*~*~*~* */
}
