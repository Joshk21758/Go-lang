package main

import "fmt"

type Point struct{
	name string
	sch string
	age int

}

type Point2 struct{
	bdy int
	*Point

}

func main(){
	//p1 := Point{"Joshua", "Arakan", 21}
	p2 := Point2{18, &Point{"Joshua", "Arakan", 21}}
	fmt.Println(p2.sch)

}