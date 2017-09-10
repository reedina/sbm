package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reedina/sbm/ctrl"
	"github.com/reedina/sbm/model"

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

}

//RunApplication - Start the HTTP server
func (a *App) RunApplication(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
