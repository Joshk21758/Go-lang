package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"

)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("The year you were born: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d by the end of the year", 2024 - input)

}