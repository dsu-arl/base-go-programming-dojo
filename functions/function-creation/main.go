package main

import (
    "fmt"
    "os"
)

/* ~*~*~*~*~*~*~*~* Begin your function definitions here ~*~*~*~*~*~*~*~* */

/* ~*~*~*~*~*~*~*~* End your function definitions here ~*~*~*~*~*~*~*~* */

func main() {
    if 6 == len(os.Args) {
        numberArgs := os.Args[1:4]
        stringArgs := os.Args[4:6]

        retValueOne := arithmetic(numberArgs[0], numberArgs[1], numberArgs[2])
        fmt.Println(retValueOne)

        retValueTwo := combination(stringArgs[0], stringArgs[1])
        fmt.Println(retValueTwo)

        retValueThree := finale(retValueOne, retValueTwo)
        fmt.Println(retValueThree)
    } else {
        /* ~*~*~*~*~*~*~*~* Write your code within this else statement ~*~*~*~*~*~*~*~* */
    }
}
