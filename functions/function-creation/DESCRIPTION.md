# Functions
Within your Go file, you can define as many functions that implement whatever functionality you want. Eventually we'll discuss packages, but for now all of your functions will live in just one file.

One reason you build functions is to extract repeated code and place it in one location. There are many reasons to do this:
- Develop once and reuse
- One location to conduct code reviews
- One debug location
- Simplifies codebase and readability

For example, lets say you wanted to do the following steps after you prompted a user for an integer value:
```go
if( userValue >= 10 ){
    fmt.Println("Number must be smaller than 10")
} else if( userValue < 0 ){
    fmt.Println("Number must be greater or equal to 0")
} else {
    fmt.Println("Your number is", userValue)
}