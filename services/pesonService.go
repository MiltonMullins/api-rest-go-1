package services

import (
	"errors"
	"github.com/miltonmullins/api-rest-go/entities"
	"github.com/miltonmullins/api-rest-go/repositories"
)

type ServicePerson interface {
	GetAll() *[]entities.Person
	Get(name string) (*entities.Person, error)
	Post(person entities.Person) (*[]entities.Person, error)
	Put(name string, person entities.Person) (*entities.Person, error)
	Delete(name string) (*entities.Person, error)
}

type servicePerson struct {
	personRepository repositories.PersonRepository
}

func NewServicePerson(personRepository repositories.PersonRepository) servicePerson {
	return servicePerson{personRepository}
}

func (s servicePerson) GetAll() *[]entities.Person {
	return s.personRepository.GetAll()
}

func (s servicePerson) Get(name string) (*entities.Person, error) {
	people := s.personRepository.GetAll()

	for _, pp := range *people {
		if pp.Name == name {
			return &pp, nil
		}
	}

	return nil, errors.New("person not found")
}

func (s servicePerson) Post(person entities.Person) (*[]entities.Person, error) {
	return s.personRepository.Post(person)
}

func (s servicePerson) Put(name string, person entities.Person) (*entities.Person, error) {
	return s.personRepository.Put(name, person)
}

func (s servicePerson) Delete(name string) (*entities.Person, error) {
	return s.personRepository.Delete(name)
}
