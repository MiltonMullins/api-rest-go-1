package services

import (
	"errors"
	"github.com/miltonmullins/api-rest-go/entities"
	"github.com/miltonmullins/api-rest-go/repositories"
)

func GetAll() *[]entities.Person {
	return repositories.GetAll()
}

func Get(name string) (*entities.Person, error) {
	people := repositories.GetAll()

	for _, pp := range *people {
		if pp.Name == name {
			return &pp, nil
		}
	}

	return nil, errors.New("person not found")
}

func Post(person entities.Person) (*[]entities.Person, error) {
	return repositories.Post(person)
}

func Put(name string, person entities.Person) (*entities.Person, error){
	return repositories.Put(name,person)
}

func Delete(name string) (*entities.Person, error){
	return repositories.Delete(name)
}