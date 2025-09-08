# Classes and Data Structures Module
This module is best completed after the Control Flow, Functions, and Input and Output modules. These modules provide the necessary foundation for this module.

These modules will also start to use some built-in functions and features of the Go language without giving them much introduction. As these modules progress through more and more content, there will simply be too many cool things to talk about to spend time describing everything! As such, it'll be up to you the reader to go and spend time looking up the different functions, their meaning, and how you can use them to your advantage!

## Built-in Data Structures
So far we've introduced data types like `int` and `bool` to store values. This is convenient when our data is simple, like storing a person's age or a person's name. But how would we store multiple peoples' ages? If it was only five people we *could* make five variables, but what if it's 100, or 1,000, or 10,000 people? What if we don't know the number of people we may need to record? As we start to learn more and approach problems we might see "out in the wild" (school, work, or various programming challenges) we need to expand how we think about data. This module introduces different ways to manage data, both how to store it (data structures) and access it (classes.)

Data structures can be simply defined as novel ways to store data. Some structures allow you to save multiple values of one type, while other structures allow you to save multiple values of different types. You can even create your own structures, defining exactly what types of data you want to store. But first, we'll introduce some of the built-in data structures that you'll encounter, and use, frequently.
>__NOTE__: A more technical definition would be "a specific way to organize, manage, and store data for efficient storage and access."

### Arrays
In Go, much like every other programming language, there are built-in data structures that we can use to help us solve problems. If you've programmed in other languages you'll likely be familiar with arrays; one of the foundational data structures. The easiest way to think of arrays is to visualize a table comprised of horizontal rows and verticle columns: an array is one row of the table. Each location within an array can be referred to by many names: element, value, etc.

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

### Slices
Slices will be the bread and butter of your array-equivalent usage in Go. There are many technical reasons why slices are preferred over arrays in Go, of which we'll get to in this section. 

Slices are implemented 'on top of' arrays, which means slices can be very dynamic and support operations like append and delete while still maintaining array-like functionality.

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