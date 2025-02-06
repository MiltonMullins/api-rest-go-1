package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people = []person{
	{Name: "Alice", Age: 25},
	{Name: "Bob", Age: 30},
}

func main() {
	router := initializeRoutes() // configure routes

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening on port:8080...")
	server.ListenAndServe() // Run the http server
}

func initializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /people", GetAll)
	mux.HandleFunc("GET /person/{name}", Get)
	mux.HandleFunc("POST /person", Post)
	mux.HandleFunc("PUT /person/{name}", Put)
	mux.HandleFunc("DELETE /person/{name}", Delete)
	return mux
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	// Get all
	log.Println("Get all")
	jsonPeople, err := json.Marshal(people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonPeople)
}

func Get(w http.ResponseWriter, r *http.Request) {
	// Get
	log.Println("Get")
	name := r.PathValue("name")

	for _, p := range people {
		if p.Name == name {
			jsonPerson, err := json.Marshal(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonPerson)
			return
		}
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	// Post
	log.Println("Post")
	var person person
	// Decode the incoming note json
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	people = append(people, person)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Person Created"))
}

func Put(w http.ResponseWriter, r *http.Request) {
	// Put
	log.Println("Put")

	name := r.PathValue("name")

	for i, p := range people {
		if p.Name == name {
			var person person
			// Decode the incoming note json
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			people[i] = person
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Person with name Updated"))
			return
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// Delete
	log.Println("Delete")
	name := r.PathValue("name")

	for i, p := range people {
		if p.Name == name {
			people = append(people[:i], people[i+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Person Deleted"))
			return
		}
	}
}
