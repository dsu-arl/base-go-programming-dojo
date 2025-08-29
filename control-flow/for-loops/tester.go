package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

/*
Generates and returns a random number. DO NOT MODIFY THIS CODE
*/
func genRandomNumber() int {
	return rand.Int() % 1001
}

func main() {
	randomNumber := 0

	if len(os.Args) == 2 {
		randomNumber, _ = strconv.Atoi(os.Args[1])
	} else {
		randomNumber = genRandomNumber() // Use this variable in the random number challenge.
	}

	for i := 0; i <= 30; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println("")

	for i := 41; i < 70; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println("")

	for i := 22; i <= 130; i += 3 {
		fmt.Print(i, " ")
	}
	fmt.Println("")

	for i := 0; i <= randomNumber; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("")

	randCount := 0
	randomNumber = genRandomNumber()

	for randomNumber > 40 {
		if randCount >= 15 {
			break
		} else {
			randCount += 1
		}
		randomNumber = genRandomNumber()
	}
	fmt.Println(randCount)
}
