package main

import "fmt"

func main(){
	//Slices
	sli := []int{1, 2, 3, 4, 5}
	sli = append(sli, 6)
	fmt.Println(sli[1:3])

}