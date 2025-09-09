# Arrays

The syntax to define an array is pretty straight forward:
```go
var myArray [5]int
```
The array above is called `myArray` and consists of five `int` elements. To access any element you use its index, starting with 0. To access the first element of the array you would use this syntax `myArray[0]`, and the last element of the array would be `myArray[4].` 
>__NOTE__: Counting from 0 can take some practice, but in programming its a worthwhile skill! 

Assigning a value to an array element is the same as assigning to any variable `myArray[4] = 123`. Note since you've already defined and created the array, you do not need to use `:=`.

There are a couple variations on creating arrays that are worth discussing. The below code segment will contain all of the variations, and they'll be discussed afterwards
```go
// Initializing array of size 5 with values
b := [5]int{1, 2, 3, 4, 5}

// Initializing array with values, but letting compiler define the size
c := [...]int{1, 2, 3, 4 ,5}

// Initializing array with specific index values.
d := [...]int{100, 3: 400, 500}
```
Here `b[0]` will contain the value 1, `b[1]` the value 2, and so forth. This can be convenient when you know the values of an array ahead of time. Note that the size of the array is defined by the programmer. This is different from `c` where we use `[...]` to tell the compiler to count the number of values. This can be convenient when you are unsure if you'll later add additional values to the initialization list. 

The final example is very unique. This example sets the first value to 100, then uses an index value of `3:`, which means the values between the 1st and 3rd indexes will be set to 0. Then it sets the 4th and 5th values to 400 and 500 respectively. This syntax can be slightly confusing, so its recommended only in very specific circumstances!

Now that we've stored values in our arrays, how do we print them? Easily! The example code below shows you how its done:
```go
a := [10]int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

// To print the whole array, just provide the name
fmt.Println("The whole array:", a)

// To print specific elements, list the elements!
fmt.Println("1st and 2nd index:", a[0], a[1])
```

Once you've mastered arrays, or 1-dimensional arrays, you can start to think in multiple dimensions! 2D arrays are the most common, and are simply thought of as tables. You create a 2D array like you do 1D arrays, but with an extra value.
```go
// The 1st value represents the number of rows (and row number), the 2nd the number of columns (and column number.)
/*
    Will create the array with the below values
    [0, 1, 2]
    [1, 2, 3]
*/
var twoDimensional [2][3]int
for i := range 2 {
    for j := range 3 {
        twoDimensional[i][j] = i+j
    }
}

```
An alternative way would be to initialize the array with values.
```go
twoDimensional [2][3]int{
    {0, 1, 2},
    {3, 4, 5},
}
```

Even though we introduced arrays first, they are not that common in Go. They are the underlying data structure for many other structures in Go, so it is important to understand how they function and their place in the language. Instead, Go uses a data structure called a slice.

# Slices

The below code section introduces slice syntax, in addition to other functionality available to slices.
```go
// Create an uninitialized slice. Note if you put a number between [] you've created an array not a slice!
var myIntSlice []int

// Expands the slice by 1 value; appending the value 100.
myIntSlice = append(myIntSlice, 100)

// Can still initialize a slice with values.
anotherSlice := []int{3, 6, 9, 12}

// Built-in function to create an integer slice of length 3, meaning it can hold 3 integers.
newSlice := make([]int, 3)

// This is similar to the above make command, but instead creates a slice of length 3 and capacity 10.
sliceWithCapacity := make([]int, 3, 10)

// You can assign to and print a slice just like an array.
newSlice[0] = 32
newSlice[1] = 135
newSlice[2] = 3

// This will delete all the values in the slices between index 1 inclusively to index 2 exclusively. In this case this means index 1 will be deleted.
newSlice = slices.Delete(newSlice, 1, 2)

// This will now print "32 3" because "135" was deleted.
fmt.Println(newSlice)
```

The biggest differences so far are the `make`, `append`, and `delete` commands, which may leave you asking; what's so nice about slices? Lets use an example to demonstrate their significance.

Imagine you're writing an accounting program and you need to input an unknown number of expenses. You *could* write it using arrays, and set the size to be *extremely* large (like 100,000 values) but that would be extremely wasteful, especially if you only needed 100 values. Instead you can create a slice of 100 values and `append` as needed.
>__NOTE__: For advanced programmers, there is a case to be made for using the `new` command to create dynamic arrays. This quickly becomes an idiomatic discussion that is outside the scope of these modules however.

Lets discuss some other features of slices that will be very useful. To make a copy of an array, you simply assign it to a new array. This does what is called a "value-copy", where all of the "values" of the original array are "copied" to the new array, making a completely new and independent array. Lets see what happens when we try that same technique on slices.
```go
firstSlice := []string{"Mike", "Jim", "Sally", "Sue"}
secondSlice := make([]string, len(firstSlice))
firstArray := [5]int{}
secondArray := [5]int{}

for i := 0; i <= len(firstSlice); i++ {
    firstArray[i] = rand.Int() % 200
}
fmt.Println("-----------BEFORE COPY")
fmt.Println(firstArray)
fmt.Println(secondArray)

fmt.Println("Contents, len, empty (T/F), capacity")
fmt.Println(firstSlice, len(firstSlice), firstSlice == nil, cap(firstSlice))
fmt.Println(secondSlice, len(secondSlice), secondSlice == nil, cap(secondSlice))

secondSlice = firstSlice
secondArray = firstArray
fmt.Println("-----------AFTER COPY")
fmt.Println(firstSlice, len(firstSlice), firstSlice == nil, cap(firstSlice))
fmt.Println(secondSlice, len(secondSlice), secondSlice == nil, cap(secondSlice))
fmt.Println(firstArray)
fmt.Println(secondArray)
// ----------OUTPUT--------
// -----------BEFORE COPY
// [67 192 112 5 69]
// [0 0 0 0 0]
// Contents, len, empty (T/F), capacity
// [Mike Jim Sally Sue] 4 false 4
// [   ] 4 false 4
// -----------AFTER COPY
// [Mike Jim Sally Sue] 4 false 4
// [Mike Jim Sally Sue] 4 false 4
// [67 192 112 5 69]
// [67 192 112 5 69]
```
All seems well-and-good; no surprises. But what happens if we make changes to the original array and slice? 
```go
firstArray[0] = 300
firstSlice[0] = "Bob"

fmt.Println("-----------AFTER MODIFICATION")
fmt.Println(firstSlice, len(firstSlice), firstSlice == nil, cap(firstSlice))
fmt.Println(secondSlice, len(secondSlice), secondSlice == nil, cap(secondSlice))
fmt.Println(firstArray)
fmt.Println(secondArray)
// ----------OUTPUT--------
// -----------AFTER MODIFICATION
// [Bob Jim Sally Sue] 4 false 4
// [Bob Jim Sally Sue] 4 false 4
// [300 192 112 5 69]
// [67 192 112 5 69]
```
What happened to the slice?

Slices are fundamentally different than arrays, and as such require different methods to handle them. Slices are referred to as "reference types" in Go, or in more general programming terms, slices contain pointers to their underlying data. When we use the "=" to "copy" arrays, Go is moving the values from one array to the other. Since the value stored by a slice is a pointer, we assign the pointer to the new slice; not the values we wanted! If you're not aware of this distinction, it can have very insidious consequences! So instead of "=" to make a copy of a slice, use the `copy` command instead; this will create a new and distinct copy of the slice as intended.
```go
sourceSlice := []int{2, 4, 6}
destSlice := make([]int, len(sourceSlice))
copy(destSlice, sourceSlice)
sourceSlice[0] = 5
fmt.Println(sourceSlice)
fmt.Println(destSlice)
// ----------OUTPUT--------
// [5 4 6]
// [2 4 6]
```
>__NOTE__: Don't worry if you didn't understand the discussion about pointers! They'll be discussed at-length in their own challenge later in this module. The discussion here is to help define the differences between arrays and slices and why they must be treated differently.

Worth a brief discussion is the `make` command. `make`, as the name suggests, creates a slice with the provided attributes. The arguments to the function are the type of slice to make, the length of the slice, and the capacity of the slice. The difference between the *length* and the *capacity* is significant; the **length** is how many elements the slice *currently* contains, and the **capacity** is the *maximum amount* of elements the slice can contain (without appending.) Without delving too deeply into the weeds, this *capacity* feature provides space for the slice to grow past what the developer might have envisioned, without requiring modification to the underlying array.

# Challenge
This challenge will be different from the challenges you've done so far (if you're going "in order" at least.) Instead of simple, strict requirements this challenge will provide a problem to solve. It'll be up to you to choose the best solution. The only requirement is that the topic presented **must** be used to solve the challenge.

You will be provided a file containing student grades. While the format of the file is shown below, helper functions will be provided to retrieve data from the file for you. Your task is to determine the class's average grade, lowest grade, highest grade, and standard deviations based on the scores contained in the document. To determine a student's grade you will use a weighted grading scale. The first 7 scores contained in the document are assignment scores, and are weighted at 60% of the total grade. The remaining 3 scores are test grades and are weighted at 40%. Each student is guaranteed to have 10 scores, and each score is guaranteed to have a value between 0 and 99.

You will be provided two helper functions.
- `returnStudentCount` takes a file object as an argument and returns the number of student records in the file (this is represented by the first line in the file.)
- `returnStudentScore` takes a file object and returns a 2D slice containing all of the student grades.

Below is an example file, with sample output of the score results. As way of example, also included below are the correct scores for each request, including the break down of the students' weighted scores.
```text
5
96 75 67 77 54 57 0 94 80 71 
94 80 73 90 72 30 54 94 79 44 
85 89 95 98 47 28 43 98 75 82 
95 88 90 74 60 68 71 90 75 88 
85 95 80 96 55 29 10 90 74 98 
```
```text
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Student 1 weighted assignment, test, and overall grade: 37%, 33%, 69%
Student 2 weighted assignment, test, and overall grade: 42%, 29%, 71%
Student 3 weighted assignment, test, and overall grade: 42%, 34%, 76%
Student 4 weighted assignment, test, and overall grade: 47%, 34%, 81%
Student 5 weighted assignment, test, and overall grade: 39%, 35%, 74%

Class average grade: 74%
Class lowest grade : 69%
Class highest grade: 81%
Standard deviation :
```


A template has been provided; use the provided functions and add to the existing code where indicated. The file is located in "/challenges." Issue the following command to move the file to your local directory. ***IT WILL DELETE ANY OTHER FILE NAMED*** `main.go` ***IN THE DESTINATION. BE CAREFUL!***
- `cp /challenge/main.go /home/hacker/`
- If you want to organize your code into folders, instead use the command `cp /challenge/main.go /home/hacker/yourFolder` where "yourFolder" is the name of the folder you want to move the file to.

1. Open a new VSCode Workspace environment and open the folder "/home/hacker/".
    - If you want to organize your code into different folders, you will need to include that folder in subsequent commands.
2. Modify the provided template to complete the challenge
3. Open a terminal in VSCode to build and run your code with the commands `go build main.go` and `./main`.
4. Verify your solution by running the command `cd /challenge` and `./verify main`.
    `main` must be the absolute path to your built Go program, not your `.go` source code file. This will likely be "/home/hacker/main" unless you organized your code differently.