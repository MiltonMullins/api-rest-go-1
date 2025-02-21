package main

import (
	"log"
	"net/http"

	"github.com/miltonmullins/api-rest-go/controllers"
	"github.com/miltonmullins/api-rest-go/repositories"
	"github.com/miltonmullins/api-rest-go/services"
)

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
	 
	controller := controllers.NewControllerPerson(
		services.NewServicePerson(
			repositories.NewPersonRepository()))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /people", controller.GetAll)
	mux.HandleFunc("GET /person/{name}", controller.Get)
	mux.HandleFunc("POST /person", controller.Post)
	mux.HandleFunc("PUT /person/{name}", controller.Put)
	mux.HandleFunc("DELETE /person/{name}", controller.Delete)
	return mux
}
