# If statements
If statements are a very powerful control flow structure, and comprise the core of many a programs' decision making abilities.
- If the result of the function is 'X', execute this function.
- If the user inputed 'Y', tell them it is not the correct value. Otherwise convert the value to a letter and print it to the screen.
- If the value is less than 10, execute "myFunction". Otherwise, if the value is greater than 10, execute "theirFunction." Otherwise, if the vaule is equal to 10, execute "errorFunction." Otherwise, print an error to the screen because something strange happened.

The list of examples could go on forever, with any amount of complexity you can imagine. Before we introduce the structure for ifs, we need to introduce the syntax for "greater than", "less than", and the other comparison operators we used above.

## Comparison Operators
Before going any further, we must discuss the comparison operators. You will use these in nearly any program you write, so having a strong understanding is essential. In Go, the comparison operators are:
- `==` (equal)
- `!=` (not equal)
- `<`  (less than)
- `<=` (less than or equal)
- `\>`  (greater than)
- `\>=` (greater than or equal)

Thankfully, you've probably been using these operators since grade school and have a basic intuition for what they mean. When you use a comparison operator it will generate a boolean result (true or false) based on the operation. Nearly always when you use these operators, you will be using a control flow statment to *do* something based on the result. They are very flexible and can be used to compare nearly every data type available in Go. Their syntax, unsurprisingly, is as follows:
```go
// If the two values ARE equal generate 'true', otherwise 'false'.
myFirstInt == mySecondInt

// If the two values ARE NOT equal, genearte 'true' otherwise 'false'.
myFirstInt != mySecondInt

// If the first value IS LESS THAN the second value, generate 'true', otherwise 'false'.
myFirstValue < mySecondValue

// If the first value IS LESS THAN OR EQUAL to the second value, generate 'true', otherwise 'false'.
myFirstValue <= mySecondValue

// If the first value IS GREATER THAN the second value, generate 'true', otherwise 'false'.
myFirstValue > mySecondValue

// If the first value is GREATER THAN OR EQUAL to the secnod value, generate 'true', otherwise 'false'.
myFirstValue >= mySecondValue
```

## If, else, and else if
Back to ifs! The structure of an if statement is very simple. You use the keyword "if" followed by the condition you want to check, and then use curly brackets to define what happens if the statement is true.
```go
if 1 < 2 {
    fmt.Println("This will print")
}
```

You can have as many if statements in sequence as you want: and you can even nest them if the situation calls for it!
```go
numOne := 10
numTwo := 20
numThree := 30
numFour := 40

// More functionality can go here if you need it to.
if numFour < numThree {
    // You can inlude whatever code you want here!
    if numThree < numTwo {
        // And here too.
        if numTwo < numOne {
            // Same here!
        }
        if numTwo > numOne {
            // And here!
        }
        // Even more code can go here!
    }
    // Why not here too?
    if numTwo > numThree {
        // Let's not stop now!
    }
    // You guessed it, code can go here too!
}
```
### If/Else

While the above code is valid, it is very messy and confusing. Instead of constantly using if statements, the 'else' keyword exists to simplify the flow of code. Using 'else', the above code becomes:
```go
numOne := 10
numTwo := 20
numThree := 30
numFour := 40

if numFour < numThree {
    if numThree < numTwo {
        if numTwo < numOne {
        
        } else {
            // This else statement executes whenever "numTwo < numOne" evaluates to FALSE.
        }
        // If you wish, you can write code here.
    } else {
        // This else statement executes whenever "numThree < numTwo" evaluates to FALSE.
    }
    // If you wish, you can write code here too.
}
```
Lets examine one if/else statement from above as a knowledge check. In the first code block two if statements, "numTwo < numOne" and "numTwo > numOne", are evaluated one after the other. In this scenario, there is no possible way for both if statements to be true: either numTwo is greater than numOne OR numTwo is less than numOne. This situation is exactly why else statements were made! We just add the 'else' after the closing '\}' of the if and include whatever functionality we want and we're done!

Here's another simple example. If you want to print the words "Success" if numOne was bigger than numTwo, and "Failure" otherwise then your code could look like the following
```go
numOne := 20
numTwo := 10

if numOne > numTwo {
    fmt.Println("Success")
} else {
    fmt.Println("Failure")
}
```

### Else If
If/else statements are useful for binary choices: do this thing or do another thing. However, not all choices boil down that simply. When you have multiple situations to check for, else if statements come to the rescue.
```go

if userInput >= 50 {
    // if means that the userInput is greater than or equal to 50.
    userGuess = userInput
} else if userInput >= 25 {
    // else if means that userInput less than 50, but greater than or equal to 25.
    userGuess = userInput
} else if userInput >= 10 {
    // else if means that userInput less than 25, but greater than or equal to 10.
    userGuess = userInput
} else {
    // else means that userInput is less than 10. 
    userGuess = 0
}
```
The final else is important; it is a catch all, or default, case. If all the else if statements don't execute, this final else will. It is good programming practice to include a default case to ensure you've accounted for as many situations as possible. Imagine the above is checking that the user did not provide a negative value. Without the else case setting the value to 0, there could be unexpected consequences for the rest of the program.

### Special Cases

While generally straightforward, comparison operators *can* be used in very complex or non-obvious ways. For example, the following Go snippet is completely valid and will generate "Yes" as the response
```go
if "test" > "TeSt" {
    fmt.Println("Yes")
} else {
    fmt.Println("No")
}
```
More of these situations will be described in a different module.

## Logical Operators
If you have programmed in C you have likely seen, and used, logical operators. These operators are more specialized than comparison operators, but no less useful. These operate on boolean values and returns booleans, which as the name suggests is ideal for logical operations. 

An example of logical operators with just boolean values is provided below.
```go
trueStatement := true
falseStatement := false

if !falseStatement {
	fmt.Println("Negating a false statement makes it TRUE; this will ALWAYS print")
}

if trueStatement && (!falseStatement) {
	fmt.Println("true AND not false is TRUE; this will ALWAYS print")
}

if trueStatement && falseStatement {
	fmt.Println("true AND false is FALSE; this will NEVER print")
}

if trueStatement || falseStatement {
	fmt.Println("true OR false is TRUE; this will ALWAYS print")
}
```

An important distinction from other programming languages (like C) logical operators *cannot* be used with non-boolean types (like integers.) However, there are ways to still use logical operators when dealing with non-boolean types.

> **NOTE**: In C, a boolean value of "true" is represented by anything non-zero, and "false" must be 0. This is not the case in Go.

Below is an example of logical *and* comparison operators with non-boolean values. 

```go
valueOne := 10
valueTwo := 3
valueThree := 24

if valueOne > valueThree || valueTwo <= valueThree {
    fmt.Println(valueOne, ">", valueThree, "=", valueOne > valueThree)
    resOne := valueOne > valueThree
    fmt.Println(valueTwo, "<=", valueThree, valueTwo <= valueThree)
    resTwo := valueTwo <= valueThree
    fmt.Println(resOne, "||", resTwo, "=", resOne || resTwo)
}
```
> **NOTE**: Print statements can also evaluate expressions, but the results will not be stored. This is useful when quickly debugging problems or practicing techniques like we are doing!

Examine the code and see if you can figure out the results before continuing; was it what you expecetd?

Since 10 is greater than 3, the first comparison evaluates to 'false'. This is stored in 'resOne.' Since 3 is less than or equal to 24, the second comparison evaluates to 'true' and is stored in 'resTwo.' Finally 'false' or'd with 'true' results in 'true', so the final result is 'true.' The output of the above code is provided below
```text
10 > 24 = false
3 <= 24 true
false || true = true
```

## Conclusion
A quick recap of everything covered by this content:
- Comparison opertators can operate on a variety of values to generate a boolean value indicating if a statement is true or false.
- Logical operators use boolean values to determine if a statement is true or false.
- If statements, along with else and else if, are primarily used to guide program flow based on the result of comparison or logical operators.

What is covered here is just the beginning! More complex and challenging use-cases exist than we can possibly hope to address. Its up to you to explore and learn how to implement them in code!

# Challenge
You will implement a program that asks the user for a number. If the number is greater than 100, print to the screen "The number is too big, run the program again." Otherwise if the number is zero print to the screen "You cannot choose zero, run the program again." Include a default case that prints to the screen "Unexpected case, run the program again."

In addition, you will implement the below functionality using if, else, and else if statements. 
- If the number is even print to the screen "Your number is even." If it isn't even, print "Your number is odd" to the screen instead.
- If the number is greater than 50, print to the screen "Your number is greater than 50". If it isn't, print "Your number is not greater than 50" instead.
- If the number is GREATER than 50, is it greater than 75? If it isn't, print to the screen "The number is between 50 and 75." 
- If the number is also GREATER than 75, is it greater than 90? If it isn't, print to the screen "The number is between 75 and 90."
- If the number is not greater than 50, is it greater than 25? If it is, print to the screen "The number is between 25 and 50."
- If the number is also LESS than 25, is it greater than 10? if it is, print to the screen "The number is between 10 and 25."

A template has been provided; use the provided functions and add to the existing code where indicated.

## BONUS
Once you have completed the challenge as described above, you may attempt this bonus challenge.

The descriptions above do not consistently account for the case where numbers are equal. For example, examine the number 50. According to the descriptions, it will say "Your number is not greater than 50" then say "Your number is between 25 and 50" when in fact the number is exactly 50. These situations are often called "edge cases" by developers.

For this bonus, you will update your code to account for the edge cases where the user enters a value that matches a conditional check. For example if the user enters the number 50, instead of printing "greater than" or "less than" you will print "Your number is equal to 50." Other checks must be updated to either print "greater than or equal to" or "less than or equal to" as appropriate.

> **HINT** There are multiple ways to solve this problem!

1. Open a new VSCode Workspace environment and open the folder "/home/hacker/".
    - If you want to organize your code into different folders, you will need to include that folder in subsequent commands.
2. Create a new file with the file extension `.go` and write your solution.
3. Open a terminal in VSCode to build and run your code with the commands `go build yourFile.go` and `./yourFile`.
4. Verify your solution by running the command `cd /challenge` and `./verify yourFile`.
    `yourFile` must be the absolute path to your built Go program, not your `.go` source code file. This will likely be "/home/hacker/yourFile" unless you organized your code differently.