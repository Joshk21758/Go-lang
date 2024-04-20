package main

import "fmt"

func main(){
	//Pointers
	x := 10
	y := &x
	fmt.Println(x, y)
	*y = 12
	fmt.Println(x, y)
	

}