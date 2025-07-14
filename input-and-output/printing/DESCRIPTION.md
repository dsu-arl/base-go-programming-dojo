# Println
The first choice we'll focus on will be the primary one we use: `fmt.Println`. `Println` is the easiest choice, and most similar to Python's `print`, for just printing data to the screen; its syntax is very straight forward.
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
```text
0
My variable's value is: 0
0 is my variable's value.
var = 0 which is my variables value

```

# Challenge
Use Go to print the text 'Hello World!' to the screen.
1. Open a new VSCode Workspace environment and open the folder "/home/hacker/Documents".
2. Create a new file with the file extension `.go`.
3. Write the Go code to print the text 'Hello World!'.
4. Open a terminal in VSCode to build and run your code with the commands `go build yourFile.go` and `./yourFile`.
5. Verify your solution by running the command `cd /challenge` and `./verify yourFile`.
    - `yourFile` must be the built Go program, not your `.go` source code file.