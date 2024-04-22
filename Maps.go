package main

import "fmt"

func main(){
	//Maps
	var mp map[string]int = map[string]int {
		"Dog": 4,
		"Cat": 7,
		"Cow": 52,
		
	}
	mp["Chick"] = 21
	mp["Egg"] = 20
	mp["Marble"] = 50
	delete(mp, "Cat")
	fmt.Println("This is a map", mp)

} 


