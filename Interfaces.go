package main

import (
	"fmt"
	"math"

)

type Area interface{
	area() float64

}

type Circle struct{
	width float64
	height float64
	pie float64

}

type Rect struct{
	width float64
	height float64

}

func (c Circle) area() float64{
	return math.Pi * c.width * c.height

}

func (r Rect) area() float64{
	return r.width * r.height

}

func main(){
	c1 := Circle{5.5, 7.2, 3.14}
	r1 := Rect{4.5, 2.5}
	fmt.Println(r1)
	fmt.Println(c1)
	grade := []int{c1, r1}

	for _, i := range grade{
		fmt.Println(_, i)
		
	}


	



	
	
}