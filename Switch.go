package main

import (
	"fmt"
	"bufio"
	"os"

)

func main(){
	//Switch statements
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your Name: ")
	scanner.Scan()
	name := scanner.Text()
	//Cases
	switch {
	case name == "Joshua":
		fmt.Println("Welcome", name)

	case name == "Michael":
		fmt.Println("Welcome", name)

	case name == "Kunda":
		fmt.Println("Welcome", name)

	default:
		fmt.Println("Who are you?")

	}

}