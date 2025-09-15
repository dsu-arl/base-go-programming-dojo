package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//~*~*~*~*~*~*~*~* DO NOT MODIFY THIS CODE ~*~*~*~*~*~*~*~*

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

//~*~*~*~*~*~*~*~* DO NOT MODIFY THIS CODE ~*~*~*~*~*~*~*~*

/*
Function to quit if an unrecoverable error is detected
*/
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	fileName := "/challenge/input.txt"
	if len(os.Args) == 2 {
		fileName = os.Args[1]
	}
	studentCount, studentGrades := returnStudentData(fileName)
	classAverageGrade := 0.0
	classLowestGrade := 0.0
	classHighestGrade := 0.0
	standardDeviation := 0.0
	//~*~*~*~*~*~*~*~* DO NOT MODIFY CODE ABOVE THIS LINE ~*~*~*~*~*~*~*~*

	//~*~*~*~*~*~*~*~* DO NOT MODIFY CODE BELOW THIS LINE ~*~*~*~*~*~*~*~*

	fmt.Printf("%.02f, %.02f, %.02f, %.02f", classAverageGrade, classLowestGrade, classHighestGrade, standardDeviation)
}
