# Variables
Variables are one of the most useful features of any programming language. They let you pass arguments to functions, store user data, and otherwise simplify the tasks of programming. Go has a few different ways to deal with variables, one of which was already introduced in Printing challenge.

## First technique
The first, and most verbose, technique we'll discuss is shown below.
```go
var myVariable int := 0
```
In this example, the statement is composed of five parts which will be described.

`var` is a Go keyword which declares that a variable is being created. `myVariable` is the name of the variable. `int` is the type of variable being created. `:=` is special in Go, and it *declares and initializes* a variable to a value: we'll discuss assignment later. Finally `0` is the value that will be stored in `myVariable`. This is the most explicit way to create a variable in Go, and is not as common as the other techniques. There are two alternatives to this style:
```go
var myVariable int
```
or
```go
var myVariable := 0
```
In the first case the value is not provided, which simply creates the variable `myVariable`. The second case stores `0` in the variable `myVariable`. Here the type is **inferred** by the Go compiler later on. You can learn more about type inferrence [here](https://go.dev/blog/type-inference).

## Second technique
The next technique is comparatively simple: `myVariable := 0`. This technique is equivalent to the first, in that it stores a value in a variable. But it offers a little more flexibility in its usage. With this technique you can assign multiple variables, and of different types, on the same line!
```go
myInt, myString = 0, "Test String"
```

## Data types
Go supports all the usual data types you'd expect. As other data types are used they'll be described in the same training. An exhaustive list won't be provided, but the basic types include:

|Name|Keyword|
|-|-|
|strings|`string`|
|signed integers|`int`|
|unsigned integers|`uint`|
|floats|`float32` or `float64`|
|boolean|`bool`|

### Strings
Strings in Go are effectively the same as strings in other languages: letters, special characters, and/or numbers all enclosed in double quotes. Single characters however *are* different in Go, and will be discussed separately. Below is an example of a string variable using both techniques
```go
var myString string := "This is my string!"
mySecondString := "Here is another string!"
```

### Integers
In Go you can be very specific about the type of integer you want to create, similar to C. You can create `int8` all the way to `int64`, and the same goes for the `uint` data type.

# Challenge
In this challenge, you will utilize the provided `main.go` template to print three different variables, with three different data types, to the screen. You will print a **string**, **signed integer**, and **boolean** value to the screen each on their own line.

1. Open a new VSCode Workspace environment and open the folder "/home/hacker/Documents".
2. Create a new file with the file extension `.go`.
3. Write the Go code to print the text 'Hello World!'.
4. Open a terminal in VSCode to build and run your code with the commands `go build yourFile.go` and `./yourFile`.
5. Verify your solution by running the command `cd /challenge` and `./verify yourFile`.
    - `yourFile` must be the built Go program, not your `.go` source code file.