package main

import "fmt"

type Author struct {
	AuthorId int32
	Person   Person
	PenName  string
}

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) fullName() string {
	return p.FirstName + " " + p.LastName
}

func (a *Author) updateName(firstName, lastName string) {
	a.Person.FirstName = firstName
	a.Person.LastName = lastName
}

func main() {
	anAuthor := Author{
		Person: Person{
			FirstName: "Bartt",
			LastName:  "Simpson",
		},
		AuthorId: 1,
	}

	fmt.Println(anAuthor.Person.fullName())

	anAuthor.updateName("Bart", "Simpson")

	fmt.Println(anAuthor.Person.fullName())

	anAuthor.PenName = "Eat My Shorts"

}
