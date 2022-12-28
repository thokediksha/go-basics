package main

import "fmt"

func mergesort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	s1 := mergesort(items[:len(items)/2])
	fmt.Println(s1)
	s2 := mergesort(items[len(items)/2:])
	fmt.Println(s2)
	return merge(s1, s2)
}

func merge(a []int, b []int) []int {
	s3 := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			s3 = append(s3, a[i])
			i++
		} else {
			s3 = append(s3, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		s3 = append(s3, a[i])
	}
	for ; j < len(b); j++ {
		s3 = append(s3, b[j])
	}
	return s3
}

func main() {
	unsorted := []int{100, 60, 20, 10, 50, 80, 30, 50, 40, 70, 90}
	sorted := mergesort(unsorted)
	fmt.Println(sorted)
}

