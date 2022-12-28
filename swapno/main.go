//  swap two numbers without using temporary variables

package main

import (
	"fmt"
)

func swap()[]int {
	a, b,  c:= 10, 12, 13
	c, b, a = a, b, c
	return []int{a,b,c}
}

func main() {
// fmt.Print("Enter first number : ")
// var first int
// fmt.Scanln(&first)
// fmt.Print("Enter second number : ")
// var second int
// fmt.Scanln(&second)
// first = first-second
// second = first+second
// first = second-first
// fmt.Println("First number :",first)
// fmt.Println("Second number :",second
fmt.Println(swap())
}

