package main 

import (

"fmt"
// "time"
"reflect"

)



func main(){
// 	for i := 0; i <= 10; i++ {
// 	fmt.Println(i)  ....    for loop
// 	}

// *
// var  x = 20
// if x == 10 {
// 	fmt.Println("Hello")
                            //  ...... if else 
// }else {
// 	fmt.Println("hi")
// }

//  * switch case

// today := time.Now()

// switch today.Day() {

// case 1:
// 	fmt.Println("Today is 1st. Clean your house.")
// case 10:
// 	fmt.Println("Today is 10th. Buy some wine.")
// case 15:
// 	fmt.Println("Today is 15th. Visit a doctor.")
// case 25:
// 	fmt.Println("Today is 25th. Buy some food.")
// case 31:
// 	fmt.Println("Party tonight.")
// default:
// 	fmt.Println("No information available for that day.")

// }

// * slice in go 
var intSlice = new([50]int)[0:10]

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

}


