package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your age: ")
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	//Decision making
	if age < 19 {
		fmt.Println("You are too young!")

	} else if age == 16 {
		fmt.Println("You are a teen")

	} else {
		fmt.Println("You are eligible...")

	}
}