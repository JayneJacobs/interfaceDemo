package people

import "fmt"

// Person strict
type Person struct{
	First string
}

// Human allows access to method speak()
type Human interface {
	Speak()
}

// Speak method for person type
func (p Person) Speak() {
	fmt.Printf("Hi I am %s and here is my concrete type %T \n", p.First, p)
}

// Secretagent struct
type Secretagent struct{
	Person
	LTK bool
}

// Speak method for secretAgent type
func (s Secretagent) Speak() {
	fmt.Printf("I am a %T and my name is %v \n", s, s.First)
}