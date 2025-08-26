# Looping
Unlike other programming languages, Go only has 1 'looping' structure: `for`. Go uses the `for` control structure as both the `for` loop and the while loop structures. This can be slightly confusing if you've programmed in other languages before, however it starts to become more natural as you develop in Go.

## Traditional for loops
The `for` loop syntax is timeless; why change perfection? Below is an example `for` loop that will print "Hello world" to the screen 10 times.
```go
for x := 0; x < 10; x++ {
    fmt.Println("Hello world")
}
```
> **_NOTE:_** The syntax `x++` means increment the value stored in `x` by 1. Go only supports this syntax, so if you're used to options other languages provides you're out of luck!

In this example the variable 'x' is instantiated and set to 0. Then the comparison operator less than is used to ensure the value stored in 'x' is less than 10. If it is, the body of the `for` loop is exeucuted, which in this case prints "Hello world" to the screen. When completed 'x' is incremented by 1, and then checked to see if it is less than 10 again. Note the `x := 0` part only happens when the `for` loop is run for the first time. This process continues until 'x' is incremented to 10. When `x < 10` is executed, since 10 is not less than 10, the for loop quits and execution continues. When writing `for` loops, there are two ';' that must be included: one after the assignment and one after the conditional.

The variations of `for` loops are near-endless. Below are just two examples:
```go
// Example 1; will print 10 - 100 in increments of 5
for x := 10; x <= 100; x += 5 {
    fmt.Println(x) 
}

// Example 2; will add the value of 'i' to 'x' for values between 2 and 39.
x := 0
for i := 2; i < 40; i++ {
    // x = x + i is syntactically the same thing, but longer.
    x += i
}
```
## Alternatives
There are many other formats that the humble `for` loop can take in Go. One format is called a 'while' loop in other programming languages. Where `for` loops can be read "for X iterations, do this thing." while loops are read "do this thing until a condition is met." Below is an example of a while loop in Go implementing the same "Hello world" example above.
```go
x := 0
for x < 10 {
    fmt.Println("Hello World")
    x += 1
}
```
Notice that all the same components are present in a while loop that exist in a traditional `for` loop, they're just in slightly different places:
- initial value
- conditional check
- body of loop

Here is another example using a function call, which will print the values 0 - 4 to the screen on new lines.
```go
for i := range 5 {
    fmt.Println("range ",i)
}
```

You can even write while loops that are non-deterministic
```go
var userInput string
// This check could also be written checkLetterCombination(userLetter) == false; it is the same outcome
for checkLetterCombination(userLetter) != true {
    fmt.Println("Wrong letter combination! Try again.")
    fmt.Scanln(&userInput)
}
```
> **_NOTE:_** Deterministic is a computer science term that basically means "the outcome can be known." The conditional `x < 10` will conclude when 'x' becomes 10, which will be in 10 iterations following our examples. Non-deterministic then basically means "the outcome cannot be known." This applies in the above example because depending on the value returned from the function call, it could take one call or many before the loop quits: there is no way to tell.

This while loop will continue executing until the user enters the correct combination of letters; who knows how long that will take!

Many other variations on this form exist, but we'll discuss only two others. The first one is an alternative form of the while loop, where the check is either moved to the body of the loop, or entirely removed!
```go
// Option 1, the check is moved to the body.
for {
    fmt.Println("This is interesting...")
    // shouldQuit is not a built-in Go function.
    if shouldQuit() == true {
        break
    }
}

// Option 2, the check doesn't exist!
for {
    fmt.Println("Infinite loop!!!!")
}
```
Practically speaking, the second option is not one you will ever see in practice; but it is valid Go code! This is called an infinite loop, namely because there is nothing causing the loop to quit, so it will continue printing until the program is terminated. This is contrasted with the first option, where the loop *could* continue indefinitely if the function call `shouldQuit` always returns `false`, but a condition does exist to exit.

`break` is a special keyword that is used to terminate, or exit, from a control flow structure. Generally this term is seen within control flow structures like while loops when you need to explicitly define an exit. In the example above we exit, or break, the control flow when `shouldQuit` returns `true`. You can use `break` in other control flow statements: for loops, if statements, etc.

However, using `break` can sometimes have unintended consequences to the state of the overall program, so be sure you know **exactly** what you're doing when you use `break`! Below is an example of how `break` could cause problems.
```go

var guess int
for i := 0; i < 10; i++ {
    if i == 4 {
        fmt.Println("Too many guesses, try again later!")
        break
    }
    fmt.Println("Enter your guess!")
    
    fmt.Scanln(&guess)
    // checkGuess is not a built-in Go function.
    if checkGuess(guess) == true {
        fmt.Println("Congrats you've guess the number!")
    }

    if i == 5 {
        fmt.Println("You've tried enough, here's the secret number")
        // showSecretNumber is not a built-in Go function.
        fmt.Println(showSecretNumber())
        break
    }
}
```
While a trivial example, this shows how `break` could change a fun game into a frustrating experience. The `if` statement at the bottom will never trigger because of the `break` occuring when `i == 4` despite the for loop indicating it should iterate through 10. The cause of the problem here is easy to identify, however real world scenarios can be much more complex and difficult to spot especially when multiple layers of functionality are in play.

# Challenge
In this challenge, you will write a program that asks the user to guess a random number. A function has been provided that takes as input the random number, the user's guess, and a hint level, and will print hints to the screen that get progressively more revealing. Below is an example of the most common way you'll call this function.
```go
randomNumber := 15
userGuess := 1
hintLevel := 2

hintGenerator(randomNumber, userGuess, hintLevel)
```
Your program must implement the following functionality using the things you've learned in this dojo.
- If the user guesses the number correctly on the first try, print the message "On your first try?!? Congratulations!!!" to the screen and quit the program.
- If the user does not guess the number correctly on the first guess, prompt the user for a hint by printing the message "You did not guess the number, would you like a hint? (y/n)" to the screen.
- For subsequent guesses, print the message "You still didn't guess it, would you like another hint? (y/n)" to the screen.
- After the first guess, prompt the user for a guess by printing the message "What's your next guess?" to the screen.
- If the user asks more than three times for a hint, simply call the `hintGenerator` function in such a way that the default case will trigger (HINT: any value that isn't 0, 1, or 2 will trigger the default case.) Note that counting starts at 0!
- When the user correctly guesses the number, print the message "Congratulations, you've guessed the number!" to the screen.
- The program should loop until the user guesses the correct number.

1. Open a new VSCode Workspace environment and open the folder "/home/hacker/".
    - If you want to organize your code into different folders, you will need to include that folder in subsequent commands.
2. Create a new file with the file extension `.go` and write your solution.
3. Open a terminal in VSCode to build and run your code with the commands `go build yourFile.go` and `./yourFile`.
4. Verify your solution by running the command `cd /challenge` and `./verify yourFile`.
    `yourFile` must be the absolute path to your built Go program, not your `.go` source code file. This will likely be "/home/hacker/yourFile" unless you organized your code differently.