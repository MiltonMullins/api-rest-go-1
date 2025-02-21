package repositories

import (
	"errors"

	"github.com/miltonmullins/api-rest-go/entities"
)

var people = []entities.Person{
	{Name: "Alice", Age: 25},
	{Name: "Bob", Age: 30},
}

func GetAll() *[]entities.Person {
	return &people
}

func Post(person entities.Person) (*[]entities.Person, error) {
	people = append(people, person)
	//TODO error check on save person on DB
	return &people, nil
}

func Put(name string, person entities.Person) (*entities.Person, error) {
	for i, p := range people {
		if p.Name == name {
			people[i] = person
			return &people[i], nil
		}
	}
	return nil, errors.New("person not found")
}

func Delete(name string) (*entities.Person, error){
	for i, p := range people {
		if p.Name == name {
			people = append(people[:i], people[i+1:]...)
			return &p, nil
		}
	}
	return nil, errors.New("person not found")
}
