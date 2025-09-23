# Pointers & Structs
Pointers are a *reasonably* advanced topic that can generate many-a-problem for even experienced developers. However, the capabilities they offer enable some very dynamic and powerful features! But before getting into all that, lets introduce them.

## What is a Pointer?
When we start writing a program and need to save information, we need to know what *type* of information we need to store: integer, string, float, etc. That's because when we assign a variable a value, that variable *stores* that *value.* We don't store floats in integer variables because we'd encounter problems with our program. To really drive the point home, when we declare a variable `var myInt int` and assign it a value `myInt = 10` we can essentially substitute `10` everywhere we see `myInt` (until the value changes of course.)

Pointers are different. Instead of storing a *value* pointers store an *address* that **points** to a value. Lets explain with an analogy.

Lets say you want to talk to your roommate Alice. So you walk out into the living room and find your other roommate Bob. Bob tells you that Alice is upstairs digging through a closet, so upstairs you go where you find Alice exactly where Bob indicated. In this scenario, Bob is the pointer who is holding the location to Alice. Bob *isn't* Alice, nor is Bob Alice's location, rather Bob just holds the knowledge of Alice's location that can be shared with other people. This is basically how pointers work. Pointers don't hold the value you're looking for but *the address* of the value. 

Lets talk some syntax and then get to an example. To declare a pointer you have a few choices, but they all involve the '\*' symbol. Below list a few examples.
```go
var myInt *int
var myInt = new(int)
func printPointer(myPointer *string)
func returnPointer() *int
```

The `new` function is extremely important if you want to make and use a pointer right away, and should look familiar to more seasoned programmers. `new` allocates memory and provides it for us to use; a fancy way of saying `new` creates a "new" variable for us dynamically. Without it we basically wouldn't be able to use pointers.

Once you've defined your pointer, you must use '\*' to assign a value to it. If you don't, you'll change the *address* stored in the pointer. Remember a pointer doesn't store the *value* but a reference *to* (the address containing) your value. Below is a program that will demonstrate this idea with some tangible output.
```go
package main

import "fmt"

func main() {
	var myInt = new(int)

	fmt.Println(myInt)
	fmt.Println(*myInt)

	*myInt = 10
	fmt.Println(myInt)
	fmt.Println(*myInt)
	myInt = nil

	fmt.Println(myInt)
}
/*
    OUTPUT
    0xc0000100f0
    0
    0xc0000100f0
    10
    <nil>
*/

```
Note what happens when you don't use the `\*` to assign a value; your address goes away! If you were to add a print statement to display your value, you'd actually crash you program. You'd receive an error similar to the one shown below.
```text
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x491c07]
```

## Using Pointers
So what's the "point" of pointers you may be asking? To be honest, for the beginner programmer there isn't much to see here. But for experienced programmer pointers can offer performance boosts, better memory efficiency, and more dynamic functionality. Pointers are utilized everywhere in the Go source code. Slices use pointers heavily to reference their underlying array, which is how they're able to offer their dynamic capabilities. In order to demonstrate this, we need to introduce our other topic: structs.

# What is a Struct?
We've discussed common data types like `int`s and `string`s, but what happens when you need to represent something more complex, lets say a person? Of course you could create many variables representing a person, but that becomes clumsy and difficult to track fast. Imagine a separate variable tracking age, height, weight, driver's license number, insurance information, home address, work address, cell phone, e-mail address...the list goes on! Before you know it you'd have to remember dozens of variable names for just one person; how would you track a whole city, or even state?! Structs to the rescue.

Structs, in essence, allow you to create your own custom data type by combining other data types.
```go
type person struct {
    first_name string
    last_name string
    age int
}
```
Once defined you can assign values in multiple ways.
```go
	bob := person{
		"Bob",
		"Smith",
		30}

	alice := person{
		first_name: "Alice",
		last_name:  "Jane",
		age:        31}
```

To print specific struct elements you use "dot notation". To access `first_name` you'd type `alice.first_name` and to access `age` you'd type `alice_age`. 

## Structs and Pointers
One way you can create structs more dynamically is by using functions that return pointers to objects.
```go
/*1*/func newPerson(fname, lname string, age int) *person{
/*2*/   p := person{
        fname,
        lname,
        age
    }
/*3*/return &p
}
```
Lets go through this line by line. The function definition on line 1 for `newPerson` takes three arguments. `fname` and `lname` are both strings and `age` is an integer. The return type is a pointer to a `person` data type (which is the same definition as above.) Line 2 uses the same method to create a new `person` as we've seen before. Line 3 however is something very different as it uses the '&' or "address of" operator. Line 3 says to return the "address of" the struct `p`. And since a pointer only holds "a reference to", or an address of a variable, we've now returned a pointer to a `person` type from the function `newPerson`. We haven't given the '\*' a name yet but we do now: it is the "dereference" operator. When used to create a pointer, such as `var myInt *int` we simply say "creating a pointer myInt" but when using it to retrieve a value, such as `fmt.Println(*myInt)` we say "dereferencing the pointer myInt." Below is the whole program.

__NOTE__: In C/C++ returning a local variable like this would result in *massive* problems due to the way those languages handle memory. Go, like some other newer languages, uses a garbage collection mechanism to handle memory, which is what enables us to return and use local variables like this. Similarly there is no 'delete' functionality. To 'assign' memory for deletion make sure there are no references to it, and the garbage collector will handle it automatically. If you want a more technical explanation search for "go garbage collection" in your search engine of choice.

```go
package main

import "fmt"

type person struct {
    first_name string
    last_name  string
    age        int
}

func newPerson(fname, lname string, age int) *person {
    p := person{fname, lname, age}
    return &p
}

func main() {
    bob := person{
        "Bob",
        "Smith",
        30}

    alice := newPerson("Alice", "Jane", 31)

    fmt.Println(bob)
    fmt.Println(alice)

}
/*
    OUTPUT:
    {Bob Smith 30}
    &{Alice Jane 31}
*/
```
Two things to note. First `fmt.Println` did not need special instructions to print the structure, which is very cool! Second, the '&' before the "Alice" struct indicates the variable is a pointer. This is extremely useful contextual information to keep in mind, especially for debugging purposes.

Not only can you return pointers, but you can pass pointers as well. For very large objects this can be a huge performance booster; but we won't delve into that here. We saw that you can print out the details of a struct using `fmt.Println` but what if you wanted it formatted in a special way? Modify the code above by adding a new function `printPerson` and calling it in `main`.
```go
func printPerson(myPerson *person) {
	fmt.Println("First name:", myPerson.first_name)
	fmt.Println("Last name:", myPerson.last_name)
	fmt.Println("Age:", myPerson.age)
	return
}
...
printPerson(&bob)
printPerson(alice)
```
Notice that because `bob` *is not* a pointer, we have to pass the address of the struct ('&' address of operator!), whereas `alice` **is** a pointer so we can just use it!

We aren't able to describe *every* use case for new topics; there are just too many! So when we introduce something new, think about how it can be used with previous things you've learned in these modules. For example, you can create slices using structs!
```go
town := make([]person, 100)
```
And now you can iterate over `town` and access each "person's" information!
```go
//range returns two values, an 'iterator' value indicating what element you're on and a 'value' at that location.
for iter, value := range town {
    fmt.Println("Person", iter, "has name", value.first_name)
}
```

## Intermediate Use Cases
Now that we have some familiarity with pointers, lets demonstrate some of their capabilities and nuances. 

Taking the example of slices above, making a slice that is compatible with our `makePerson` function is extremely easy.
```go
//Now we have a slice of pointers to person structs!
town := make([]*person, 100)
``

Since pointers just contain addresses, multiple variables can contain the same address. Below we assign the `alice` pointer to a new variable `alice_temp`.
```go
alice_temp := alice
```
`alice_temp` can access `first_name`, `last_name`, and `age` just like it were a `person` struct: because it is! ***But*** it isn't a *new* `person` struct. Observe what happens when you change the value of `alice_temp.age`.
```go
alice_temp.age = 50
fmt.Println(alice, alice_temp)
/*
    OUTPUT
    &{Alice Jane 50} &{Alice Jane 50}
*/
```
First they're both pointers, as indicated by '&'. Second, and more importantly, `alice`'s age changed even though we only updated `alice_temp`'s age. This is an important lesson. Because pointers contain addresses to data, and not the data itself, whenever we change the underlying data all references to that data reflect that change. While a challenging concept to grasp when first working with pointers, this is why pointers are so powerful and useful, but also complex and challenging.

To further complicate this, lets examine two alternative ways to work with pointers, going back to `bob`. Remember, we initialized `bob` at the top of the program, so `bob` **is not** a pointer. We can assign `bob` to `bob_temp` and change the age of `bob_temp` without changing the age of `bob`, just like any normal variable interaction.
```go
bob_temp := bob
bob_temp.age = 101
fmt.Println(bob, bob_temp)
/*
    OUTPUT
    {Bob Smith 30} {Bob Smith 101}
*/
```
But we can easily create pointers **to** `bob` by using the '&' (address of) operator.
```go
bob_temp := &bob
bob_temp.age = 101
fmt.Println(bob, bob_temp)
/*
    OUTPUT
    {Bob Smith 101} &{Bob Smith 101}
*/
```
Now we can create a **NEW** variable containing the old information, but separate from the original!
```go
bob_temp := &bob                        // Get the address of bob and store it in bob_temp
bob_temp_two := *bob_temp               // Get the address bob_temp is storing and store it in bob_temp_two
bob_temp.age = 101
fmt.Println(bob, bob_temp, bob_temp_two)
/*
    OUTPUT
    {Bob Smith 101} &{Bob Smith 101} {Bob Smith 30}
*/
```
If things weren't complicated enough, instead of dereferencing `bob_temp`, imagine if you instead just assigned it to `bob_temp_two`!
```go
bob_temp := &bob                        // Get the address of bob and store it in bob_temp
bob_temp_two := bob_temp               // Get the address OF bob_temp and store it in bob_temp_two
bob_temp.age = 101
fmt.Println(bob, bob_temp, bob_temp_two)
/*
    OUTPUT
    {Bob Smith 101} &{Bob Smith 101} &{Bob Smith 101}
*/
```
This is why pointers, '\*' (deference) operator, and '&' (address of) operator cause so many troubles even to experienced developers. A simple flip, or omission, of '\*' and '&' can cause serious damage, AND be extremely difficult to find and debug.

# Challenge
This challenge will be different from the challenges you've done so far (if you're going "in order" at least.) Instead of simple, strict requirements this challenge will provide a problem to solve. It'll be up to you to choose the best solution. The only requirement is that the topic(s) presented **must** be used to solve the challenge.

## Description
This challenge expands on the work done in "Arrays and Slices." If you've already completed that challenge you'll be able to reuse some of that code. If you haven't then you get to start fresh here!

You'll once again be given a file containing student scores, with the same helper functions as before. Your task this time however is to create a structure that stores all of the student's scores and also stores their lowest score, highest score, and calculates the weighted final grade for the course. Therefore the structure **must at minimum** consist of a slice storing the float values, two floats storing the highest and lowest score, and a float storing the final weighted grade. Additional elements may be added for your convenience, but are not required to successfully complete the challenge.

Once again the first 7 scores are assignment scores weighted at 60% and last 3 are test scores weighted at 40%. The lowest and highest score are chosen out of all the scores. Each student is guaranteed to have 10 scores, and each score is guaranteed to have a value between 0 and 99.

Once completed you'll again determine the class's highest, lowest, and average(mean) grade and calculate the standard deviation. Now for the additional twist! Once the standard deviation is calculated, you will determine the number of students that land within one, two, and three standard deviations of the mean. To determine where a student's score lands, you add and subtract the standard deviation from the median score to get a range. For additional standard deviations, you multiply the value accordingly: 2x for the 2nd standard deviation, 3x for the 3rd, etc.

For example, if the mean grade is 84% and the standard deviation is 5%, then you'll be counting the number of students that fall into the following bands:
1st standard deviation: 79% - 89%
2nd standard deviation: 74% - 94%
3rd standard deviation: 69% - 99%

The formula for standard deviation (SD) is provided below, as well as calculating weighted grades (WG).
$$
SD = \sqrt{\frac{\sum(x-\mu)^2}{N}}
$$
Where $x$ is an element from the set of final grades, $\mu$ is the mean of the set, $N$ is the total number of final grades.

To calculate the weighted grades, you sum each grade in the category, divide by the total points of the category, then multiply by the weight. These two formulas are provided below as Weighted Assignment (WA) and Weighted Test (WT).

$$WA = \frac{\sum(y)}{700} * 60\%$$
$$WT = \frac{\sum(z)}{300} * 40\% $$

### Hints
If you find yourself struggling with populating structs, give this article a read on [Medium.com](https://medium.com/@caring_smitten_gerbil_914/why-your-go-range-loop-isnt-updating-slice-values-and-what-to-do-instead-4428ae2b369e)

## Required Packages
You will need the `math` library to solve this challenge. **And `int` will not be the only data type required by this challenge!**

## Directions
A template has been provided; use the provided functions and add to the existing code where indicated. The file is located in "/challenges." Issue the following command to move the file to your local directory. ***IT WILL DELETE ANY OTHER FILE NAMED*** `main.go` ***IN THE DESTINATION. BE CAREFUL!***
- `cp /challenge/main.go /home/hacker/`
- If you want to organize your code into folders, instead use the command `cp /challenge/main.go /home/hacker/yourFolder` where "yourFolder" is the name of the folder you want to move the file to.

1. Open a new VSCode Workspace environment and open the folder "/home/hacker/".
    - If you want to organize your code into different folders, you will need to include that folder in subsequent commands.
2. Modify the provided template to complete the challenge
3. Open a terminal in VSCode to build and run your code with the commands `go build main.go` and `./main`.
4. Verify your solution by running the command `cd /challenge` and `./verify main`.
    `main` must be the absolute path to your built Go program, not your `.go` source code file. This will likely be "/home/hacker/main" unless you organized your code differently.