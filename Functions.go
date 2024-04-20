package main

import "fmt"

func add(x, y, z int){
	ans := x * y * z
	fmt.Println("The product is", ans)

}

func main(){
	add(5, 6, 5)

}