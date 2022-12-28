// Create a function that takes two numbers as arguments (num, length) and returns a list of multiples of num until the list length reaches length.

//Examples
// list_of_multiples(7, 5) ➞ [7, 14, 21, 28, 35]
// list_of_multiples(12, 10) ➞ [12, 24, 36, 48, 60, 72, 84, 96, 108, 120]
// list_of_multiples(17, 6) ➞ [17, 34, 51, 68, 85, 102]


package main

import "fmt"

func main(){
  c := multiples(7,5)
  fmt.Println(c)
  d := multiples(12,10)
  fmt.Println(d)
  e := multiples(17,6)
  fmt.Println(e)
}

func multiples(num int, length int) []int {
	var res []int

	for count:=1; count<=length; count++ {
         res = append(res, num*count)
	}
  return res
}