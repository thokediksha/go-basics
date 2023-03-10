package main

import "fmt"

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}



func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	done := make(chan bool)
	mid := len(data) / 2
	var left []int

	go func() {
		left = MergeSort(data[:mid])
		done <- true
	}()
	right := MergeSort(data[mid:])
	<-done
	return Merge(left, right)

}

func MergeSortWithoutGoroutine(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	return Merge(left, right)


	mid := len(data)  /2

}

func main() {
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Printf("%v\n", data)
	//fmt.Printf("%v\n", MergeSortWithoutGoroutine(data))
	fmt.Printf("%v\n", MergeSort(data))
	/*
	   [1 2 3 4 5 6 7 8 9 10]
	*/

}


// type tank interface {
// 	Volume() float64
// }

// type myValue struct {
//    radius float64
//    height float64
// }

// func (v myValue) Volume() float64 {
// 	return 3.14 * v.radius * v.radius * v.height
// }

// func main(){
// 	var t tank
// 	t = myValue{10.5, 4}
// 	fmt.Println("volume:", t.Volume())
// }