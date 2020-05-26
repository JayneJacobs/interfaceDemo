package main

import (
	"fmt"

	"github.com/JayneJacobs/interfaceDemo/dbs"
	"github.com/JayneJacobs/interfaceDemo/people"
)


func main() {
setpeople()
usedb()
 
}


func put(a dbs.Accessor, n int, p dbs.Person)   {
	a.Save(n,p)
}

func get(a dbs.Accessor, n int) dbs.Person {
	return  a.Retrieve(n)
}
func usedb()  {
	dbm := dbs.Mongo{}
	dbp := dbs.Postg{}
	p1 := dbs.Person{
		First: "Jayne",
	}
	p2 := dbs.Person{
		First: "James",
	}
	put(dbm, 1, p1)
	put(dbm, 2, p2)

	fmt.Println((get(dbm, 1)))
	fmt.Println((get(dbm, 2)))
	put(dbp, 1, p1)
	put(dbp, 2, p2)

	fmt.Println((get(dbp, 1)))
	fmt.Println((get(dbp, 2)))

}


func setpeople()  {
	Person1 := people.Person{
		First: "Jayne",
	}
   sa1 := people.Secretagent {
	   Person: people.Person{
		   First: "James",
	   },
   
	   LTK: true,
   }
	Person1.Speak()
   
	var h,sa people.Human
	sa = sa1
	h = Person1
	 sa.Speak()
	h.Speak()
}