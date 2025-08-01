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

If you need to use multiple integers in your program, you can define multiple in one declaration. Note however you cannot assign multiple values using this format; that is discussed next!
```go
var myFirstVar, mySecondVar int
```

## Second technique
The next technique is comparatively simple: `myVariable := 0`. This technique is equivalent to the first, in that it stores a value in a variable. But it offers a little more flexibility in its usage. With this technique you can assign values to multiple variables, and of different types, all on the same line!
```go
myInt, myString = 0, "Test String"
```

## Data types
Go supports all the usual data types you'd expect. As other data types are used they'll be described in the same training. An exhaustive list won't be provided, but the basic types include:

Name and Keyword
- strings - `string`
- signed integers - `int`
- unsigned integers - `uint`
- floats - `float32` or `float64`
- boolean - `bool`

### Strings
Strings in Go are effectively the same as strings in other languages: letters, special characters, and/or numbers all enclosed in double quotes. Single characters however *are* different in Go, and will be discussed separately. Below is an example of a string variable using both techniques
```go
var myString string := "This is my string!"
mySecondString := "Here is another string!"
```

### Signed and Unsigned Integers
Go provides access to both signed integers (int) and unsigned integers (uint). Signed integers are called that because they can hold either positive or negative values, whereas unsigned integers can only hold positive values. To do this, the highest order bit (most significant bit) of signed integers is used to reflect positive (0) or negative (1) values. Becaues of this, the maximum and minimum value that a signed 32-bit integer can be is 2^31, or +/-2,147,483,648. Whereas its unsigned counterpart's maximum value is 2^32, or +4,294,967,296.

In Go you can be very specific about the type of integer you want to create, similar to C. The following are the types of integers that Go supports.
- `int` or `uint` (equivalent to either int32/int64 or uint32/uint64 depending on the system.)
- `int8` or `uint8`
- `int16` or `uint16`
- `int32` or `uint32`
- `int64` or `uint64`


# Challenge
In this challenge, you will utilize the provided `main.go` template to print three different variables, with three different data types, to the screen. You will print a **string**, **signed integer**, and **boolean** value to the screen each on their own line. Do not use any special characters in the string! Below is an example of possible output:
```
Hello world
31615
true
```
Remember that each `Println` ends with a newline, so three `Println` function calls will emit three lines of text.

1. Open a new VSCode Workspace environment and open the folder "/home/hacker/".
    - If you want to organize your code into different folders, you will need to include that folder in subsequent commands.
2. Create a new file with the file extension `.go` and write your solution.
3. Open a terminal in VSCode to build and run your code with the commands `go build yourFile.go` and `./yourFile`.
4. Verify your solution by running the command `cd /challenge` and `./verify yourFile`.
    `yourFile` must be the absolute path to your built Go program, not your `.go` source code file. This will likely be "/home/hacker/yourFile" unless you organized your code differently.