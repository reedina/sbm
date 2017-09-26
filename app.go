package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reedina/sbm/ctrl"
	"github.com/reedina/sbm/model"
	"github.com/rs/cors"

	//Initialize pq driver
	_ "github.com/lib/pq"
)

//App  (TYPE)
type App struct {
	Router *mux.Router
}

//InitializeApplication - Init router, db connection and restful routes
func (a *App) InitializeApplication(user, password, url, dbname string) {

	model.ConnectDB(user, password, url, dbname)
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

//InitializeRoutes - Declare all application routes
func (a *App) InitializeRoutes() {

	//model.Environment struct
	a.Router.HandleFunc("/api/environment", ctrl.CreateEnvironment).Methods("POST")
	a.Router.HandleFunc("/api/environments", ctrl.GetEnvironments).Methods("GET")
	a.Router.HandleFunc("/api/environment/{id:[0-9]+}", ctrl.GetEnvironment).Methods("GET")
	a.Router.HandleFunc("/api/environment/{id:[0-9]+}", ctrl.UpdateEnvironment).Methods("PUT")
	a.Router.HandleFunc("/api/environment/{id:[0-9]+}", ctrl.DeleteEnvironment).Methods("DELETE")

	//model.Environment struct
	a.Router.HandleFunc("/api/environment/instance", ctrl.CreateEnvironmentInstance).Methods("POST")
}

//RunApplication - Start the HTTP server
func (a *App) RunApplication(addr string) {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	log.Fatal(http.ListenAndServe(addr, c.Handler(a.Router)))
}
