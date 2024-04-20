package main

import "fmt"

func main(){
	//AND OR NOT expressions
	num1 := 40
	num2 := 50
	n3 := 45
	n4 := 60
	n5 := 70
	n6 := 25
	stat := n5 < n6 && num1 < num2
    val := num1 < num2 || n3 > n4
	fmt.Println("This is", val)
	fmt.Println("And this is", stat)
	 
}