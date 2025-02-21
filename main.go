package main

import (
	"log"
	"net/http"

	"github.com/miltonmullins/api-rest-go/controllers"
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
	mux := http.NewServeMux()
	mux.HandleFunc("GET /people", controllers.GetAll)
	mux.HandleFunc("GET /person/{name}", controllers.Get)
	mux.HandleFunc("POST /person", controllers.Post)
	mux.HandleFunc("PUT /person/{name}", controllers.Put)
	mux.HandleFunc("DELETE /person/{name}", controllers.Delete)
	return mux
}
