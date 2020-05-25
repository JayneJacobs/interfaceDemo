package main

import "fmt"

type person struct{
	first string
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Printf("Hi I am %s and here is my type %T \n", p.first, p)
}

func main() {
 person1 := person{
	 first: "Jayne",
 }

 person1.speak()

 var h human

 h = person1

 h.speak()
}