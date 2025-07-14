# Input and Output Module
This module will teach you the fundamentals of how Go prints to the screen and takes input from users. 
Sending data to the screen is the fundamental way that programs interact with users. If you don't send data to the user there's no way for them to interact with your program! In Go, there are a few different choices to print data, depending on your intent.

For all these challenges you will be required to write a Go program with the `.go` extension. You can name these files whatever you want. In challenges when you see `yourFile.go`, replace that with the file you made. To run a Go program, first you must compile it with `go build yourFile.go` and then `./yourFile` to run it.

# Go file syntax
Over the course of these modules different componets of the Go source code file will be discussed. Multiple tutorials can be found at the [Go website](https://go.dev/doc/tutorial) or other places on the internet if you prefer learning it all at once.

## Package
The first line in every Go program is the `package` keyword. In short, this is used to create modules and identifies where the main function is located. For almost all of our code, we'll use `package main`.

## Import
If you've programmed in any other language, then you know about importing libraries to extend functionality. In Go the syntax will feel familiar, but it has an added twist for multiple packages.

In Go, you use the `import` command to bring in additional functionality. If you only import one package, it can be on one line `import "fmt"` to import the 'fmt' (format) package. However, if you use multiple packages the syntax changes.
```go
    import (
        "fmt"
        "math/rnd"
    )
```
When using Goland by JetBrains, the environment automatically updates the `import` field when it detects you using a new package. Other environments you must do this manually.

## Functions
We'll cover function declarations later, but one function that must be declared for your program to work is `main`. This is done by using the `func` keyword followed by the function name.
```go
func main(){
    // Your code will go here
}
```
There is more you can do with function declarations, but for now this is the minimum to get started.

## All Together
When you bring all of these components together, you have the bare bones structure necessary to create a Go program! This is what everything looks like brought together.
```go
package main

import "fmt"

func main(){
    // Your code here
}
```
If you were to compile this code *without* using the `fmt` import, the Go compiler would emit an error. Continue to the first challenge to see how to use a print function from within your Go program.