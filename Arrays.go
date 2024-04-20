package main

import "fmt"

func main(){
	//Arrays
	arr := [5]int{4, 5, 6, 7, 8}
	arr[3] = 0
	fmt.Println("This is an Array", arr)
	fmt.Println("It has", len(arr), "elements")

	//2D Arrays
	arr2D := [3][4]float64{{0.5, 2.1, 1.5},{1.1, 1.3, 2.1, 2.5}}
	arr2D[0][0] = 0.6
	fmt.Println("This is a 2D array", arr2D)
	fmt.Println("It has", len(arr2D[1]), "elements")
	

}