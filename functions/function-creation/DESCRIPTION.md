# Functions
Within your Go file, you can define as many functions that implement whatever functionality you want. Eventually we'll discuss packages, but for now all of your functions will live in just one file.

One reason you build functions is to extract repeated code and place it in one location. There are many reasons to do this:
- Develop once and reuse
- One location to conduct code reviews
- One debug location
- Simplifies codebase and readability

Lets use the below code for an example.
```go
fmt.Println("Please enter a number:")
fmt.Scanln(&myNumber)
fmt.Println("You've entered:", myNumber)
```
A simple enough series of instructions that you could write numerous times without much problem. But imagine if you wanted to ask the user for multiple numbers. You could do this in a `for` loop, doing something like:
```go
myNumOne := 0
myNumTwo := 0
myNumThree := 0

for (i := 0; i < 3; ++i){
    fmt.Println("Please enter a number:")
    if 0 == i {
        fmt.Scanln(&myNumOne)
    } else if 1 == i{
        fmt.Scanln(&myNumTwo)
    } else if 2 == i{
        fmt.Scanln(&myNumThree)
    }
}

fmt.Println("You've entered the numbers:", myNumOne, myNumTwo, myNumThree)
```

But notice how complicated its already become. What happens if you want to add another value? Another 10 values? An unspecified amount of values? Or what if you need to ensure the user's number meets certain criteria? This is where functions come to the rescue. Below is an example function declaration
```go
func myPrintFunction(userNum int, iterations int)(int)
```

To define a function, you use the keyword `func` followed by the name you want to give the function. For now, function names must start with a lowercase letter. After the function name you provide the arguments to the function surrounded by parentheses. You can have as many arguments to a function as you need, however the fewer the better. These arguments can be named any way that you want, the only requirement is that the type must be provided after the name. At the end you must specify the return type, if there is one. If there is no return type then you don't have to specify anything! In the example above, there is only one return value and it is an integer. Unlike some other programming languages, if you want to return multiple values you just add its type to the return list. For example, if you wanted to return two `ints` your return line would be `(int, int)` or an `int`, `string`, and `float` would be `(int, string, float64)`. It's that easy!

> __NOTE__: The naming convention surrounding functions is based on public/private functions and their visibility within and outside of packages. This will be discussed more when we introduce packages.

> __NOTE__: When there is only one return value, the parentheses are not necessary. If there is more than one however, they are required. Including them will not prevent compilation, but some editors will automatically remove them; so don't be surprised!

Using our example above, one possible way we can define a function would be:
```go
func getUserNumber()(int){
    userNumber := 0
    fmt.Println("Please enter a number:")
    fmt.Scanln(&userNumber)
    fmt.Println("You've entered the number:", userNumber)

    return userNumber
}
```

Note at the end of the function we use the `return` keyword to "return" the `int` we said we'd return in the function definition. If we wanted to return two `int`s we'd simply say `return (intOne, intTwo)` (or whatever we've named our return variables.)

We can now use this function to get multiple values from the user in a much simplier form.
```go
numOne := getUserNumber()
numTwo := getUserNumber()
numThree := getUserNumber()

fmt.Println(numOne, numTwo, numthree)
```

To really make our point, lets say instead of just one number we wanted to get two numbers from the user, but if the first number was bigger than 50 then we don't need the second number. Lets modify the function
```go
func getUserNumber()(int, int){
    useNumOne := 0
    useNumTwo := 0
    fmt.Println("Please enter a number:")
    fmt.Scanln(&userNumOne)
    if userNumber <= 50 {
        fmt.Println("Please enter your second number:")
        fmt.Scanln(&userNumTwo)
    }

    return useNumONe, useNumTwo
}
```

Without functions, we'd have to rewrite that code *every time* we utilized that functionality in our code. Even worse, we'd have to *remember where* we implemented that functionality! As we said earlier, functions not only simplify our code but make it easier to fix and update.

## Best Practices
When it comes to returning from functions, there are varying opinions. Take the two function definitions below as examples:
```go
func functionOne(numOne int, numTwo int)(int){
    retValue := 0
    if numOne > numTwo {
        retValue = 1
    } else {
        retValue = 2
    }

    return retValue
}

func functionTwo(numOne int, numTwo int)(int){
    if numOne > numTwo {
        return 1
    } else {
        return 2
    }
}
```
Functionally, these two functions do the same thing; if `numOne` is greater than `numTwo`, return 1, otherwise return 2. However, the exact way this happens is different between the two. `functionOne` utilizes a variable to contain the return value, whereas `functionTwo` simply returns the value. For simple functions like this, there really isn't much to discuss. But imagine a more complicated scenario. Imagine we're tasked with processing data and only allowing text through that adheres to some rules. The below code is an example of one such way to implement this.
```go
func processText(input string) string{
    // Rule 1: empty string
    if input == '' {
        return "Error: empty input"
    }

    // Rule 2: too short
    if len(input) < 3 {
        return "Error: too short"
    }

    // Rule 3: contains numbers
    for _, r := range input {
        if r >= '0' && r <= '9' {
            return "Error: numbers not allowed"
        }
    }
    
    // Rule 4: specific bad word
    if strings.Contains(input, "bad") {
        return "Error: contains banned word"
    }

    // Rule 5: everything passed
    return "Valid input: " + strings.ToUpper(input)
}
```
Imagine we pass the string "Only 5 people can travel the badlands." It would take 2 runs through this function to detect the number and the banned word. Additionally if any cleanup needed to happen before leaving the function, it must occur in five different places. Imagine instead the below code (once again there are multiple ways to implement this functionality.)

```go
func appendText(errMsg string, returnVar string) string {
	if "" == returnVar {
		returnVar = errMsg
	} else {
		returnVar = returnVar + " AND\n" + errMsg
	}

	return returnVar
}
func processText(input string) string {

	retValue := ""

	// Rule 1: empty string
	if "" == input {
		retValue = appendText("Error: empty input", retValue)
	}

	// Rule 2: too short
	if len(input) < 3 {
		retValue = appendText("Error: too short", retValue)
	}

	// Rule 3: contains numbers
	for _, r := range input {
		if r >= '0' && r <= '9' {
			retValue = appendText("Error: numbers not allowed", retValue)
		}
	}

	// Rule 4: specific bad word
	if strings.Contains(input, "bad") {
		retValue = appendText("Error: contains banned word", retValue)
	}

	// Rule 5: everything passed
	if "" == retValue {
		retValue = "Valid input: " + strings.ToUpper(input)
	}

	return retValue
}
```
With only one pass through the function, all the errors are reported! While one decision isn't necessarily better than the other, these types of decisions can have wide ranging impacts from the ability to read and debug code to maintability. In production environments these decisions aren't usually left to individual developers, but contained in a company Coding Standard document specifying things like spaces VS tabs, function name conventions, and return statement usage.


# Challenge
In this challenge you will write a program that uses functions. Implement the below functionality.
- Write a function that takes three integers as arguments. It adds the first number to the third number, then subtracts the second number from the sum. If the value is negative, return 0, otherwise return the value. Print the return value to the screen.
- Write another function that takes two strings as arguments. Concatenate them using the "+" operator, separating both words with one space, and return the new value. Print the return value to the screen.
    - For example, if you call the function with the words "tiny" and "fruit" the output will be "tiny fruit"
- Write another function that takes the input from the first two functions (an `int` and `string`). Find the length of the string, add it to the `int` argument, and return the total. Print the return value to the screen.
    - >__HINT__ You will use the `len` function to find the length of strings.

1. Open a new VSCode Workspace environment and open the folder "/home/hacker/".
    - If you want to organize your code into different folders, you will need to include that folder in subsequent commands.
2. Create a new file with the file extension `.go` and write your solution.
3. Open a terminal in VSCode to build and run your code with the commands `go build yourFile.go` and `./yourFile`.
4. Verify your solution by running the command `cd /challenge` and `./verify yourFile`.
    `yourFile` must be the absolute path to your built Go program, not your `.go` source code file. This will likely be "/home/hacker/yourFile" unless you organized your code differently.