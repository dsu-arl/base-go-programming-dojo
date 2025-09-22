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

## Intermediate Use Cases
Now that we have some familiarity with pointers, lets demonstrate some of their capabilities and nuances. Since pointers just contain addresses, multiple variables can contain the same address. Below we assign the `alice` pointer to a new variable `alice_temp`.
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