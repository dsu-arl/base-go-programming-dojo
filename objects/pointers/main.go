package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Function responsible for reading data from file, formatting,
and returning int values to user.
*/
func returnStudentData(fileName string) (int, [][]float64) {
	file, err := os.Open(fileName)
	checkErr(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	//Read in the first line; number of student scores in file
	fileScanner.Scan()
	studentCount, err := strconv.Atoi(fileScanner.Text())
	checkErr(err)

	studentGrades := make([][]float64, studentCount)
	j := 0
	//Read the remaining lines
	for fileScanner.Scan() {
		data := strings.Fields(fileScanner.Text())
		//convert data from string to int
		dataSlice := make([]float64, 10)
		for i := range data {
			temp, err := strconv.Atoi(data[i])
			checkErr(err)
			dataSlice[i] = float64(temp)
		}
		studentGrades[j] = dataSlice
		j++
	}

	return studentCount, studentGrades
}

/*
Function to quit if an unrecoverable error is detected
*/
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type challengeOutput struct {
	classAverageGrade    float64
	classLowestGrade     float64
	classHighestGrade    float64
	standardDeviation    float64
	studentCountSD_One   int
	studentCountSD_Two   int
	studentCountSD_Three int
}

//~*~*~*~*~*~*~*~* DO NOT MODIFY CODE ABOVE THIS LINE ~*~*~*~*~*~*~*~*

//~*~*~*~*~*~*~*~* START CUSTOM FUNCTIONS ~*~*~*~*~*~*~*~*

//~*~*~*~*~*~*~*~* END CUSTOM FUNCTIONS ~*~*~*~*~*~*~*~*

func main() {

	fileName := "/challenge/input.txt"
	if len(os.Args) == 2 {
		fileName = os.Args[1]
	}
	studentCount, studentGrades := returnStudentData(fileName)
	results := challengeOutput{
		classAverageGrade:    0.0,
		classLowestGrade:     0.0,
		classHighestGrade:    0.0,
		standardDeviation:    0.0,
		studentCountSD_One:   0,
		studentCountSD_Two:   0,
		studentCountSD_Three: 0}

	//~*~*~*~*~*~*~*~* DO NOT MODIFY CODE ABOVE THIS LINE ~*~*~*~*~*~*~*~*

	// INSERT YOUR CUSTOM CODE

	//~*~*~*~*~*~*~*~* DO NOT MODIFY CODE BELOW THIS LINE ~*~*~*~*~*~*~*~*

	fmt.Printf("%.02f, %.02f, %.02f, %.02f", results.classAverageGrade, results.classLowestGrade, results.classHighestGrade, results.standardDeviation)
	fmt.Printf("%d, %d, %d", results.studentCountSD_One, results.studentCountSD_Two, results.studentCountSD_Three)
}
