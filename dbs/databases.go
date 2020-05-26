package dbs

type Person	struct {
	First string
}

type Postg map [int]Person 
type Mongo map [int]Person

func (m Mongo) Save(n int, p Person) {
	m[n] = p
}

func (m Mongo) Retrieve(n int)Person {
	return m[n]

}


func (pg Postg) Save(n int, p Person) {
	pg[n] = p
}

func (pg Postg) Retrieve(n int)Person {
	return pg[n]

}

// Accessor interface
type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}



func put(a Accessor, n int, p Person)   {
	a.Save(n,p)
}

func get(a Accessor, n int) Person {
	return  a.Retrieve(n)
}