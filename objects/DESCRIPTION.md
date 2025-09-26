# Classes and Data Structures Module
This module is best completed after the Control Flow, Functions, and Input and Output modules. These modules provide the necessary foundation for this module.

These modules will also start to use some built-in functions and features of the Go language without giving them much introduction. As these modules progress through more and more content, there will simply be too many cool things to talk about to spend time describing everything! As such, it'll be up to you the reader to go and spend time looking up the different functions, their meaning, and how you can use them to your advantage!

Additionally, the challenges will become more difficult. Not everything you need to solve the challenge will be given to you; again there's just too much to cover! Instead, where applicable, specific packages that contain useful functionality will be mentioned and it'll be left to you to find the functions you need to solve the problem. You'll never be left out in the dark, but you *will* be challenged to think through the problems.

## Built-in Data Structures
So far we've introduced data types like `int` and `bool` to store values. This is convenient when our data is simple, like storing a person's age or a person's name. But how would we store multiple peoples' ages? If it was only five people we *could* make five variables, but what if it's 100, or 1,000, or 10,000 people? What if we don't know the number of people we may need to record? As we start to learn more and approach problems we might see "out in the wild" (school, work, or various programming challenges) we need to expand how we think about data. This module introduces different ways to manage data, both how to store it (data structures) and access it (classes.)

Data structures can be simply defined as novel ways to store data. Some structures allow you to save multiple values of one type, while other structures allow you to save multiple values of different types. You can even create your own structures, defining exactly what types of data you want to store. But first, we'll introduce some of the built-in data structures that you'll encounter, and use, frequently.
>__NOTE__: A more technical definition would be "a specific way to organize, manage, and store data for efficient storage and access."

While not a complete list of topics covered in this module, we'll briefly introduce arrays, slices, runes, and const here and then dive into further detail in their own challenges.

### Arrays
In Go, much like every other programming language, there are built-in data structures that we can use to help us solve problems. If you've programmed in other languages you'll likely be familiar with arrays; one of the foundational data structures. The easiest way to think of arrays is to visualize a table comprised of horizontal rows and verticle columns: an array is one row of the table. Each location within an array can be referred to by many names: element, value, etc.

However once an array is created you cannot alter its size. For simple programs this isn't a big deal, but for advanced applications (and many of the Go internals) this is not ideal. Instead Go uses something called slices.

### Slices
Slices will be the bread and butter of your array-equivalent usage in Go. There are many technical reasons why slices are preferred over arrays in Go.

Slices are implemented 'on top of' arrays, which means slices can be very dynamic and support operations like append and delete while still maintaining array-like functionality. Slices achieve this because they are referred to as a "reference type", meaning they actually store a pointer to data rather than the data itself.

### Runes
Simply put, runes are Go's way of representing UTF-8 "code points", except in Go these values are stored, and represented, as an `int32` type. Don't worry if your head hurts after reading that sentence; it'll all be made clear in the challenge!

### Const
Less a data structure and more a special way to interact with data types , `const` is a keyword in Go and has special behavior compared to its counterpart in other languages.