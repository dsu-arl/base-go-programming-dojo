package main

import (
    "fmt"
    "math"
)

/*
Function name: checkError
Arguments: error
Purpose: Checks if there is an error and quits the program if one exists.
Return Value(s):
  - None
*/
func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

/*
Function name: validateNumber
Arguments: integer pointer
Purpose: Ensures the number is not 0 or negative, and loops until one is entered
Return value(s): None
*/
func validateNumber(num *int) {

    for (*num) <= 0 {
        fmt.Print("Number must not be 0 or negative; try gain: ")
        _, err := fmt.Scanf("%d", num)
        checkError(err)
    }
}

type cylinder struct {
    length      int
    diameter    int
    height      int
    volume      float64
    surfaceArea int
}

/*
Function name: setDimensions
Arguments: integer slice
Purpose: Assigns the user provided values to the elements of the structure.
  - 0 = length
  - 1 = diameter
  - 2 = height

Return Value(s):
  - None
*/
func (c *cylinder) setDimensions(dimSlice []int) {
    c.length = dimSlice[0]
    c.diameter = dimSlice[1]
    c.height = dimSlice[2]

    return
}

/*
Function name: calcvolume
Arguments: None
Purpose: Calculates the volume for a cylinder and stores the result.
Return Value(s):
  - None
*/
func (c *cylinder) calcVolume() {
    fmt.Println("Calculating Volume for Cylinder...")
    c.volume = math.Pi * math.Pow(float64(c.diameter)/2.0, 2.0) * float64(c.height)
}

/*
Function name: calcSurfaceArea
Arguments: None
Purpose: Calculates the surface area for a cylinder and stores the result.
Return Value(s):
  - None
*/
func (c *cylinder) calcSurfaceArea() {
    fmt.Println("Calculating Surface Area for Cylinder...")
    c.surfaceArea = 2 * ((c.length * c.diameter) + (c.length * c.height) + (c.diameter * c.height))
}

/*
Function name: printValues
Arguments: None
Purpose: Prints the cylinder's values, including the calculated ones.
Return Value(s):
  - None
*/
func (c *cylinder) printValues() {
    fmt.Printf("Cylinder: Length = %d, diameter = %d, height = %d, volume = %.02f, surface area = %d\n", c.length, c.diameter, c.height, c.volume, c.surfaceArea)
}

/*
Function name: requestDimensions
Arguments: integer slice
Purpose: Asks the user to ender the dimensions for a cylinder: Length, Diameter, Height.
Return Value(s):
  - None
*/
func (c *cylinder) requestDimensions(dimSlice []int) {
    var ln, wd, ht int

    fmt.Print("Length: ")
    _, err := fmt.Scanf("%d", &ln)
    checkError(err)

    validateNumber(&ln)
    dimSlice[0] = ln

    fmt.Print("Width/Diameter: ")
    _, err = fmt.Scanf("%d", &wd)
    checkError(err)

    validateNumber(&wd)
    dimSlice[1] = wd

    fmt.Print("Height: ")
    _, err = fmt.Scanf("%d", &ht)
    checkError(err)

    validateNumber(&ht)
    dimSlice[2] = ht

}

type geometry interface {
    setDimensions([]int)
    requestDimensions([]int)
    calcVolume()
    calcSurfaceArea()
    printValues()
}

/*
Function name: initializeShape
Arguments: geometry interface
Purpose: Requests values from the user and sets those values within the structure.
Return Value(s):
  - None
*/
func initializeShape(g geometry) {
    dimSlice := make([]int, 3) // 0 = length, 1 = width/diameter, 2 = height for cylinder

    fmt.Println("Please enter the following dimensions for your shape: ")
    g.requestDimensions(dimSlice)
    g.setDimensions(dimSlice)
}

/*
Function name: calculate
Arguments: geometry interface
Purpose: Calculates the volume and surface area for the shape.
Return Value(s):
  - None
*/
func calculate(g geometry) {
    fmt.Println("Running calculations...")
    g.calcVolume()
    g.calcSurfaceArea()
}

func main() {
    cyl := cylinder{}
    initializeShape(&cyl)
    calculate(&cyl)
    cyl.printValues()

}
