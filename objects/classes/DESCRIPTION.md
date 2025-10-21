# Classes in Go
Implementing classes in Go is a complicated topic because *technically* there is not a 'class' keyword in Go!

If you've used classes in other languages Go's implementation will feel weird to start. But over the course of this challenge you'll start to see where some similarities exist. And don't worry right now about traditional class terms like member variables/functions, public, private, and protected fields, etc. We'll discuss all of that and more in the package & modules module

So you may be asking...why is there a challenge about "Classes in Go" when there are no classes in Go!? We'll answer that and more in this challenge, so lets tuck in and get going!

## Structs and Methods
Since there isn't a 'class' keyword in Go, how do we make classes? We use `structs` and methods defined to only operate on those `structs`. If you haven't done the pointers challenge (where we introduce structs as well) please go back and do that challenge first.

Lets provide some code first and then we'll describe the functionality.
```go
type square struct {
	length int
	width  int
}

func (s square) area() int {
	return (s.length * s.width)
}

func main() {
	mySquare := square{length: 10, width: 5}
	fmt.Println(mySquare.area())
}
```
In the pointers challenge we introduced `structs`; what they are, how to make one, and some examples of using them. Here we define a `struct` and call it `square` with two variables `length` and `width` both `int`s. We then define a function with what is called a *receiver argument* in Go parlance; this is the `(s square)` component. When we define a function with a receiver component we now call it a *method.* *Methods* are distinguished from *functions* in this way: *methods* have receiver arguments where functions do not. While this seems like a small difference, it is the same thing that distinguishes *member functions* from *non-member functions* when interacting with classes in other languages.

You can create methods on any type ***that is defined within the same package.*** Since we haven't discussed packages and modules yet, for us this simply means any type defined within our code file. Significantly, this excludes the built-in types like `int` and `float.`

>__NOTE__: To complicate matters, you can *technically* use built-in types if you create a new type in your package. This is done by using `type <name> int` where `<name>` is whatever new name you want to give the `int` type. For example if you wanted to create a `negate` function on `int`s you could do the following:
```go
type myInt int
func (i myInt) negate() int{
    return int(-i)
}
func main(){
    negative := myInt(5)    //Note the use of '()' instead of '{}'.
    fmt.Println(negative.negate())
}
```

## Pointer Receivers
We discussed pointers in the challenge of the same name, so if you haven't done that challenge yet go back and complete it first.
Defining a pointer receiver is the most likely way you'll use receivers. This is because passing by value creates a new local copy to the method making it difficult, or in some cases impossible, to properly implement the intended functionality. So instead we use a pointer to pass by reference. What does that look like? We'll modify the `area` code as an example.
```go
type square struct {
	length int
	width  int
	area   int  // added area variable to the struct
}
//Notice the * in the receiver argument
func (s *square) calcArea() {
	s.area = s.length * s.width //Assigning the variable instead of returning the value.
}

func main() {
	mySquare := square{length: 10, width: 5}
	mySquare.calcArea()
	fmt.Println(mySquare.area)
}
```
Run the code and you should get the same result as before. However, if you remove the '\*' and rerun it, you'll notice a different value printed. This is due to the struct being passed by value (removing the '\*') instead of pass by reference (using the '\*').

>__NOTE__: Strictly speaking in the above code example, Go translates the method call `mySquare.calcArea()` as `(&mySquare).calcArea()` since the method is defined with a pointer receiver. The opposite conversion happens if a pointer is used to invoke the method `(*mySquare).calcArea()`. While potentially forgettable details, it may be important to remember when debugging sophisticated code.

## Interfaces
To discuss interfaces, lets discuss a wrinkle with structs and receivers. First lets modify the code above slightly.
```go
type square struct {
	side  int   // Don't need length and width for a square!
	area  int
	perim int   // New perimeter value
}

type rect struct {
	length int
	width  int
	area   int
	perim  int  // New perimeter value
}

func (s *square) calcArea() {
	s.area = s.side * s.side
}
func (s *square) calcPerim() {
	s.perim = s.side * 4
}

func (r *rect) calcArea() {
	r.area = r.length * r.width
}

func (r *rect) calcPerim() {
	r.perim = (r.length * 2) + (r.width * 2)
}

func printRect(r rect) {
	fmt.Printf("length: %d, width: %d, area: %d, perim: %d\n", r.length, r.width, r.area, r.perim)
}

func printSquare(s square) {
	fmt.Printf("side: %d, area: %d, perim: %d\n", s.side, s.area, s.perim)
}

func main() {
	sq := square{side: 7}
	rt := rect{width: 10, length: 3}

	sq.calcArea()
	sq.calcPerim()

	rt.calcArea()
	rt.calcPerim()

	printRect(rt)
	printSquare(sq)
}
```
All of the receivers and methods are defined as needed, however look at the redundancy in the `main` function. Since a `square` is not a `rect` separate functions are required for everything! With this simple example it's clear this isn't a sustainable programming practice. This is where the `interface` keyword in Go helps. Interfaces provide a set of methods that a type must implement to participate as a 'type' of that interface. For example we can create an interface of type geometry `type geometry interface` that specifies the functions `calcPerim`, `calcArea`, and new method `printMeasure()` must be defined to use the interface. All of the new code is provided below.
```go
type geometry interface{
    calcArea()
    calcPerim()
    printMeasure()
}
...
func (s *square) printMeasure() {
	fmt.Printf("side: %d, area: %d, perim: %d\n", s.side, s.area, s.perim)
}

func (r *rect) printMeasure() {
	fmt.Printf("length: %d, width: %d, area: %d, perim: %d\n", r.length, r.width, r.area, r.perim)
}

func calculate(g geometry) {
	g.calcArea()
	g.calcPerim()
	g.printMeasure()
}
func main() {
	sq := square{side: 7}
	rt := rect{width: 10, length: 3}

    // Notice the '&' is required to satisfy the receiver requirements for a pointer but the interface does not have a '*'.
	calculate(&sq)  
	calculate(&rt)
}
```
The simplicity! Because the structs `rect` and `square` implement the necessary functions we can create a new function that accepts the `geometry` interface and executes the functions! Notice that the interface itself cannot have a receiver argument.

## Embedding Interfaces
For the advanced developers, notice how this behavior is similar to polymorphism. Lets assume the code we've completed so-far is part of a built-in library. Now we want to create a circle struct that we can use with, and eventually expand, the `geometry` interface. First we need to define the 3 methods (`calcArea`, `calcPerim`, `printMeasure`) required by the interface. Once that's done we want to implement a unique function for circles: `calcDiameter`. We could create a brand new interface containing all 4 functions, but why when an already, perfectly good interface exists? This is where embedding comes in.

First lets define our structure.
```go
type circle struct{
    radius      float64
    diameter    float64
    area        float64
    perim       float64
}
```

Next we implement our 4 functions. The capitalization ***is very significant*** but we'll come back to it when we discuss packages and modules. For now know that the functions that implement the `geometry` interface must be exported (first word of the function is capitalized.) Note in contrast our `calcDiameter` function is not exported (first word of the function is ***not*** capitalized.)
```go
func (c *circle) CalcArea() {
    c.area = c.radius * c.radius * math.Pi
}

func (c *circle) CalcPerim() {
    c.perim = 2 * math.Pi * c.radius
}

func (c *circle) PrintMeasure() {
    fmt.Printf("radius: %f, diameter: %f, area: %f, perim: %f\n", c.radius, c.diameter, c.area, c.perim)
}

func (c *circle) calcDiameter() {
    c.diameter = 2 * c.radius
}
```

Now for the cool part. We need to define a new `interface` that combines the `geometry` interface and our new `calcdiameter` function and a function that does all the calculations like `geometry`'s `calculate` function. This is what that looks like:
```go
type circleMath interface{
    shape.Geometry  //shape is the package containing all our original work. Hang tight we'll discuss packages and modules soon!
    calcDiameter()
}

func circleCalculate(g circleMath) {
    g.calcDiameter()
    shape.Calculate(g)
}
```

First we define a new interface `circleMath` which requires the `calcDiameter` method and also the `geometry` interface, which subsequently requires all its methods. We've imposed the requirements of `geometry` on `circleMath` by *embedding* it within the new interface! Now within `circleCalculate`, since its function argument is an interface of type `circleMath` we have access to not only the `calcDiameter` method but can invoke the `calculate` function from before!

> __NOTE__: Embedding can also be done with structures, but is not discussed here and left as an exercise for the reader.

# Challenge
For this challenge, instead of a list of strict and relatively straight-forward requirements you are given a problem to solve. It'll be up to you to choose the best solution. The only requirement is that the topic(s) presented **must** be used to solve the challenge.

## Description
This challenge will continue the shapes examples, but now we're switching to three dimensions; instead of circles and squares we'll be building spheres and cubes. A template has been provided with code implementing a cylinder and calculations to determine its volume and surface area. Your task is to use this code as a springboard to create structures, methods, functions, and if necessary interfaces, to implement a sphere.

Your sphere must, at minimum, adhere to the `geometry` interface. This means your code must, at minimum, implement these methods:
```go
- setDimensions([]int)
- requestDimensions([]int)
- calcVolume()
- calcSurfaceArea()
- printValues()
```

This gives you access to the `initializeShape` and `calculate` functions, which you should utilize in your code rather than developing your own.

## Hints
Your `requestDimensions` and `setDimensions` functions must use a slice.

## Required Packages
The built-in `math` package is used by the `cylinder` code and already imported and should be the only additional package you need. You are free to use any additional packages you wish however.

## Directions
In other challenges you have been given the code that prints the solutions. This time you must write your own. This is also part of the challenge, so pay close attention to the requirements below!

Below is a list of requirements and formatting restrictions that your program ***must follow*** in order to retrieve the flag. These requirements are reflected by the `cylinder` object in the supplied template.
- You ***must*** print all your shape's values and calculations to the screen. Use the provided `printValues()` function as the template for your own print function.
    - You ***must*** print the name of your shape first followed by a ':'.
    - You ***must*** print all of the shape's measurements. The order is not significant, but you must include the measurement's name and value separated by an '=' sign.
    - You ***must*** print all of the shape's calculations. Again the order is not significant, but you must include the name and value separated by an '=' sign.
    - All of this output ***must*** be on 1 line: ***do not*** use the `Println` function or the `\n` newline character as these will insert new lines. It's OK if the text naturally wraps due to the terminal size.
    - All values ***must*** be greater than 0: 0 and negative numbers are not allowed.
- All other output from the program will be ignored and is for your own debugging or user interaction purposes. However, following good programming practices you should minimize 

A template has been provided; use the provided functions and add to the existing code where indicated. The file is located in "/challenges." Issue the following command to move the file to your local directory. ***IT WILL DELETE ANY OTHER FILE NAMED*** `main.go` ***IN THE DESTINATION. BE CAREFUL!***
- `cp /challenge/main.go /home/hacker/`
- If you want to organize your code into folders, instead use the command `cp /challenge/main.go /home/hacker/yourFolder` where "yourFolder" is the name of the folder you want to move the file to.

1. Open a new VSCode Workspace environment and open the folder "/home/hacker/".
    - If you want to organize your code into different folders, you will need to include that folder in subsequent commands.
2. Modify the provided template to complete the challenge
3. Open a terminal in VSCode to build and run your code with the commands `go build main.go` and `./main`.
4. Verify your solution by running the command `cd /challenge` and `./verify main`.
    `main` must be the absolute path to your built Go program, not your `.go` source code file. This will likely be "/home/hacker/main" unless you organized your code differently.