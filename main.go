package main

import "fmt"

type person struct{
	first string
}
func (p person) speak() {
	fmt.Printf("Hi I am %s and here is my type %T \n", p.first, p)
}

func main() {
 person1 := person{
	 first: "Jayne",
 }

 person1.speak()
}