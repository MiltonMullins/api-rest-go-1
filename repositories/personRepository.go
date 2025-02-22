package repositories

import (
	"database/sql"
	"github.com/miltonmullins/api-rest-go/entities"
)

/* var people = []entities.Person{
	{Name: "Alice", Age: 25},
	{Name: "Bob", Age: 30},
} */

type PersonRepository interface {
	GetAll() *[]entities.Person
	GetByName(name string) (*entities.Person, error)
	Post(person entities.Person) (*entities.Person, error)
	Put(name string, person entities.Person) (*entities.Person, error)
	Delete(name string) (*entities.Person, error)
}

type personRepository struct{
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) personRepository {
	return personRepository{db}
}

func (p personRepository) GetAll() *[]entities.Person {
	rows, err := p.db.Query("SELECT * FROM people")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	people := []entities.Person{}
	for rows.Next() {
		var person entities.Person
		err := rows.Scan(&person.ID, &person.Name, &person.Age)
		if err != nil {
			panic(err)
		}
		people = append(people, person)
	}
	return &people
}

func (p personRepository) GetByName(name string) (*entities.Person, error) {
	rows, err := p.db.Query("SELECT * FROM people WHERE name = $1", name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var person entities.Person
	for rows.Next() {
		err := rows.Scan(&person.ID, &person.Name, &person.Age)
		if err != nil {
			//TODO dont throw a panic and stop the service
			panic(err)
		}
	}
	//TODO: return an error if the person is not found
	return &person, nil
}

func (p personRepository) Post(person entities.Person) (*entities.Person, error) {
	_, err := p.db.Exec("INSERT INTO people (name, age) VALUES ($1, $2)", person.Name, person.Age)
	if err != nil {
		panic(err)
	}

	return &person,nil
}

func (p personRepository) Put(name string, person entities.Person) (*entities.Person, error) {
	_, err := p.db.Exec("UPDATE people SET name = $1, age = $2 WHERE name = $3", person.Name, person.Age, name)
	if err != nil {
		//TODO dont throw a panic and stop the service
		panic(err)
	}

	return &person, nil
}

func (p personRepository) Delete(name string) (*entities.Person, error) {
	_, err := p.db.Exec("DELETE FROM people WHERE name = $1", name)
	if err != nil {
		//TODO dont throw a panic and stop the service
		panic(err)
	}

	return nil, nil
}
