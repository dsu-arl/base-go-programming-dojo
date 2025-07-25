# User Input
So far we've learned how to create and store values in variables and print their values, along with other text, to the screen. But how do we write programs to get information from the user? That is exactly what this challenge will teach you!

## Scan()
Much like `Println`, the functions discussed in this challenge all reside in the `fmt` package, and must be preceded with `fmt.`. The `Scan` function is probably the most straight forward one to use. If you want the user to provide a single value you provide one variable: doesn't matter if its an intereg, string, bool, or float.
```go
fname := ""

fmt.Println("Please enter your first name")
fmt.Scan(&fname)

fmt.Println("Your name is:", fname)
```
***Remember, you can also define a variable as*** `var fname string` ***and you will get the exact same result.***

You can include multiple variables in the `Scan` statement separated by spaces. The user can then enter multiple values, each separated by a single space.
```go
fname, lname := "", ""

fmt.Println("Please enter your first name")
fmt.Scan(&fname &lname)

fmt.Println("Your name is:", fname, lname)
```
Of special note, `Scan` will treat new lines (carriage returns, '\n', single press of the 'ENTER' button, etc.) as spaces, so the two inputs below are treated the same by `Scan`.
```text
INPUT ONE: Mike Sonoma
INPUT TWO: Mike
Sonoma
```
## Scanln()
The `Scanln` function is nearly identical to `Scan` except it stops scanning for input when it detects a new line. So using the example above, 'INPUT ONE' and 'INPUT TWO' ***are not*** the same, and would result in "Mike Sonoma" and "Mike" being printed to the screen, respectively. Otherwise it behaves the exact same as `Scan`!

## Scanf
`Scanf` is the most versitle and complex choice out of the three choices. If you have programmed in C before, this will be the most familiar.

### Format & Verbs
Unlike the other scan functions, `Scanf` requires specific tokens, or verbs, to collect and store the input data appropriately. In this challenge we'll focus on just the common verbs you may encounter, but a complete list of possible verbs is provided on [their website](https://pkg.go.dev/fmt).

When using verbs, the best practice is to chose the verb most aligned with your variables data type. Below is a small list of ones you'll see frequently.
- `%d` : integer (base10 value)
- `%s` : string
- `%x` : hexadecimal value

To use these, the format is `fmt.Scanf("%d", &varName)`. The `&` is important, but 

# Challenge

1. Open a new VSCode Workspace environment and open the folder "/home/hacker/".
    - If you want to organize your code into different folders, you will need to include that folder in subsequent commands.
2. Create a new file with the file extension `.go`.
3. 
4. Open a terminal in VSCode to build and run your code with the commands `go build yourFile.go` and `./yourFile`.
5. Verify your solution by running the command `cd /challenge` and `./verify yourFile`.
    `yourFile` must be the absolute path to your built Go program, not your `.go` source code file. This will likely be "/home/hacker/Documents/yourFile".