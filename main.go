package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq" // postgres driver

	"github.com/miltonmullins/api-rest-go/controllers"
	"github.com/miltonmullins/api-rest-go/repositories"
	"github.com/miltonmullins/api-rest-go/services"
)

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create the table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS people (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		age INT NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
	}

	router := initializeRoutes(db) // configure routes

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening on port:8080...")
	server.ListenAndServe() // Run the http server
}

func initializeRoutes(db *sql.DB) *http.ServeMux {



	controller := controllers.NewControllerPerson(
		services.NewServicePerson(
			repositories.NewPersonRepository(db)))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /people", controller.GetAll)
	mux.HandleFunc("GET /person/{name}", controller.GetByName)
	mux.HandleFunc("POST /person", controller.Post)
	mux.HandleFunc("PUT /person/{name}", controller.Put)
	mux.HandleFunc("DELETE /person/{name}", controller.Delete)
	return mux
}
