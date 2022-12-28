package main

import (
	"fmt"
)

func MissingNumber(arr []int, n int) int {
	sum := 0
	n_num_sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i] // sum variable will contain the sum fo all integers of an array.
	}
	n_num_sum = ((n + 1) * (n + 2)) / 2 // we use formula for obtaining sum of first n numbers.
	return (n_num_sum - sum)            // return the missing number.
}

func main() {
	arr := []int{1,2,3,4,6,7,8,9,10} // declaring and initialising the array having n-1(8) integers.
	size1 := len(arr)
	missing_no := MissingNumber(arr, size1)
	fmt.Printf("The missing number in the array of size %d is %d", size1+1, missing_no) // Printing the missing number
}

