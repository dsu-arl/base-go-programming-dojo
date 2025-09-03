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