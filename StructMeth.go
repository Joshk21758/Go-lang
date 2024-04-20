package main

import "fmt"

type Student struct{
	name string
	grades []int
	age int

}

/*func (s *Student) setAge(age int){
	s.age = age

}*/

func (s *Student) getAverageGrade() int{
	sum := 0
	for _, i := range s.grades{
		sum += i


	}
	return sum / len(s.grades)

}

func main(){
	s1 := Student{"Mwansa", []int{20, 40, 50, 60, 85}, 21}
	average := s1.getAverageGrade()
	fmt.Println(average)

	
}