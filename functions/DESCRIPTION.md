# Functions Module
This module is best completed after completing the Input and Output module, where you learn the basics of printing to the screen, getting input from users, and the general layout of a Go file. If you are already familiar with these concepts, you're ready for this module!

For all these challenges you will be required to write a Go program with the `.go` extension. You can name these files whatever you want. In challenges when you see `yourFile.go`, replace that with the file you made. To run a Go program, first you must compile it with `go build yourFile.go` and then `./yourFile` to run it.

## Functions
You've already been introduced to the most important function in a Go file, `main`. This function is where package execution begins and, in complex packages, where initialization of classes and structures occurs.

### Definition
To define a function, you must use the keyword `func` before your chosen function name. Like every programming language, there are restrictions on the symbols that can be used in function names. For specific lists and instructions on what symbols can or can't be used in function names, please examine the Golang Language Specification. In short, only letters can start a function name, and then any combination of letters or numbers can be used to complete the function name. Examples of legal function names include

```text
myAddingFunction
myRestrictedFunction
minus1and2
tHiSiShArDtOrEaD123
```
While legal, names such as `tHiSiShArDtOrEaD123` and `minus1and2` are not recommended for readability and usability reasons.

In Go, the capitalization of the first letter of a function indicates if the function is "exported" or "unexported," or in other words if the function is available outside of the package (exported) or restricted to the package (unexported.) If you went through the Input and Output module, you're already familiar with this concept. `Println` and the `Scan` family of functions are capitalized so that we can use them outside of the `fmt` package. If you were to inspect this package, you would likely find many unexported functions assisting in the package's work. The majority of modules in this dojo will utilize unexported functions, unless specifically noted in the challenge.

### Arguments
While functions don't require arguments, most often you will create functions with arguments. Golang does not restrict the number of function arguments you can have; practically, you should limit the number of arguments you use to a reasonable amount. For example, a function requiring 13 arguments would be cumbersome to use.

If you want to create a function with three arguments, two strings and an integer, there are a few options available:

```go
func myFunction(var1 string, var2 string, var3 int)
func myFunction(var1, var2 string, var3 int)
func myFunction(var1 int, var2, var3 string)
```
When multiple arguments of the same type are present, you can specify their type individually (as in example 1,) or you can combine them (as in the remaining examples.) This makes the signature more concise but does not offer any other advantages.

There are a couple other ways to pass arguments that will be introduced in-depth in their own challenges: structures and interfaces, generics. The last one we'll introduce in this introduction is variadic. Variadic functions are special in that they accept a variable number of arguments of a specific type.

```go
func variadicExample(numbers ...int)
```
`...` is how variadic functions are declared, followed immediately by the data type; notice not even a space is present. If you want to include other function arguments, they must be placed before the variadic component.

```go
func anotherExample(var1 string, numbers ...int)
func yetAnotherExample(var1 string, var2 string, numbers ...int)
```

### Return Types
If you need to return a value, or multiple values, from a function you must specify the return type(s). This is done at the end, as opposed to other languages like C that require it at the beginning. When only one return value is used, `()` are not necessary.
```go
func anotherExample(var1 string, numbers ...int) int
```
In Go you can return multiple values, and of differing types, simply by specifying the return type in the definition. You can even provide a variable name, which declares and initializes the variable for you. If you provide names for your return values, you don't even need to specify them in the return statement! Two example functions showing this difference are provided below. Note that even when only one return type is used, if you provide a name, it must be surrounded by `()`.

```go
//Function with a return name. Notice the `()` surrounding it.
func testFunction(num1 int, num2 int)(ret1 int){
    ret1 = num1 + num2

    return
}

//Function without return names, notice no `()`.
func testFunction2(num1 int, num2 int)int{
    var ret1 = num1 + num2

    return ret1
}
```

### Main
In Go, `main` is a special function. It does not take arguments nor have a return type in the way we've discussed so far, however you can still pass arguments to `main` and return values. When you execute `main`, you do so from the command line. So to pass arguments to `main`, you must do so from the terminal.
```bash
./main arg1 arg2 arg3
```
A package called `os` contains functionality to access these command line arguments. `os.Args` will return a slice (a mutable, or changable, view of data) containing the program name and the command line arguments. For example, the below program will print all of the command line arguments, plus program name, to the screen. If you build this program as `main` and run it as `./main abc def ghi` the output will be `[./main abc def ghi]`.
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	commandLine := os.Args
	fmt.Println(commandLine)
}
```
If you do not want the program name, alter the 9th line to read `commandLine := os.Args[1:]` which will return a slice containing only the command line arguments.