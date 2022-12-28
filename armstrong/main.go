// Golang Program to check Armstrong Number

package main
import "fmt"

func main() {
	var number,tempNumber,remainder int
	var result int =0
	fmt.Print("Enter any three digit number :")
	fmt.Scan(&number)

	tempNumber = number

	// Use of For Loop as While Loop
	for {
		remainder = tempNumber % 10
		result += remainder*remainder*remainder		
		tempNumber /=10
		
		if(tempNumber==0) {
			break  // Break Statement used to stop the loop
		}
	}
	
   if(result==number){
	   fmt.Printf("%d is an Armstrong number.",number)
   }else{
	fmt.Printf("%d is not an Armstrong number.",number)
   }
}


