package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/miltonmullins/api-rest-go/entities"
	"github.com/miltonmullins/api-rest-go/services"
)

type ControllerPerson interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type controllerPerson struct{
	servicePerson services.ServicePerson
}

func NewControllerPerson(servicePerson services.ServicePerson) controllerPerson{
	return controllerPerson{servicePerson}
}

func (c controllerPerson) GetAll(w http.ResponseWriter, r *http.Request) {
	// Get all
	log.Println("Get all")
	jsonPeople, err := json.Marshal(c.servicePerson.GetAll())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonPeople)
}

func (c controllerPerson) Get(w http.ResponseWriter, r *http.Request) {
	// Get
	log.Println("Get")
	name := r.PathValue("name")

	person, err := c.servicePerson.Get(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonPerson, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write(jsonPerson)
	fmt.Fprint(w,string(jsonPerson))
}

func (c controllerPerson) Post(w http.ResponseWriter, r *http.Request) {
	// Post
	log.Println("Post")
	var person entities.Person
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = c.servicePerson.Post(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Person Created")
}

func (c controllerPerson) Put(w http.ResponseWriter, r *http.Request) {
	// Put
	log.Println("Put")

	name := r.PathValue("name")
	var person entities.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//people[i] = person
	_, err = c.servicePerson.Put(name, person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v Updated", name)
}

func (c controllerPerson) Delete(w http.ResponseWriter, r *http.Request) {
	// Delete
	log.Println("Delete")
	name := r.PathValue("name")

	_, err := c.servicePerson.Delete(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Person with name %v was Deleted", name)
}
