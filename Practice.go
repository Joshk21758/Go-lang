package main

import "fmt"

func main(){
	//looping through arrays
	arr1 := [6]int{2, 4, 6, 8, 10, 12}
	arr1[1] = 5
	//for loop
	for i, element := range arr1{
		fmt.Println(i, element)

	}
	
	
}