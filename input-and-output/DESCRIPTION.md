# Input and Output Module
This module will teach you the fundamentals of how Go prints to the screen and takes input from users. 
Sending data to the screen is the fundamental way that programs interact with users. If you don't send data to the user there's no way for them to interact with your program! In Go, there are a few different choices to print data, depending on your intent. The first choice we'll focus on will be the primary one we use: `fmt.Println`.

## Println
`Println` is the easiest choice, and most similar to Python's `print`, for just printing data to the screen; its syntax is very straight forward.
```go
fmt.Println("Hello world!")
```
The `fmt` preceeding the function is the package in which `Println` resides. This programming idiom will be common while programming in Go, and is also common in other programming languages, like Python. \(Whenever a new package is used it must be included in the `import` section, commonly found towards the top of the program, but we won't dwell on that detail here.\)

`Println` can print basically anything you'll need in the course of writing a program. Below are other examples. Don't worry too much about the syntax for variables, it'll be introduced in subsequent modules.
```Go
temp := 0

fmt.Println(temp)
fmt.Println("My variable's value is:", temp)
fmt.Println(temp,"is my variabel's value.")
fmt.Println("temp =",temp,"which is my variables value")
```

This will result in the following output
```
0
My variable's value is: 0
0 is my variable's value.
var = 0 which is my variables value

```