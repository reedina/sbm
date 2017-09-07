package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reedina/buildenvironment/ctrl"
	"github.com/reedina/buildenvironment/model"

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
	//model.User struct
	a.Router.HandleFunc("/api/user", ctrl.CreateUser).Methods("POST")
	a.Router.HandleFunc("/api/users", ctrl.GetUsers).Methods("GET")
	a.Router.HandleFunc("/api/user/{id:[0-9]+}", ctrl.GetUser).Methods("GET")
	a.Router.HandleFunc("/api/user/profile/{email}", ctrl.GetUserByEmail).Methods("GET")
	a.Router.HandleFunc("/api/user/{id:[0-9]+}", ctrl.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/api/user/{id:[0-9]+}", ctrl.DeleteUser).Methods("DELETE")

	//model.Team struct
	a.Router.HandleFunc("/api/team", ctrl.CreateTeam).Methods("POST")
	a.Router.HandleFunc("/api/teams", ctrl.GetTeams).Methods("GET")
	a.Router.HandleFunc("/api/team/{id:[0-9]+}", ctrl.GetTeam).Methods("GET")
	a.Router.HandleFunc("/api/team/{id:[0-9]+}", ctrl.UpdateTeam).Methods("PUT")
	a.Router.HandleFunc("/api/team/{id:[0-9]+}", ctrl.DeleteTeam).Methods("DELETE")

	//model.Project struct
	a.Router.HandleFunc("/api/project", ctrl.CreateProject).Methods("POST")
	a.Router.HandleFunc("/api/projects", ctrl.GetProjects).Methods("GET")
	a.Router.HandleFunc("/api/project/{id:[0-9]+}", ctrl.GetProject).Methods("GET")
	a.Router.HandleFunc("/api/project/{id:[0-9]+}", ctrl.UpdateProject).Methods("PUT")
	a.Router.HandleFunc("/api/project/{id:[0-9]+}", ctrl.DeleteProject).Methods("DELETE")

	//model.Environment struct
	a.Router.HandleFunc("/api/environment", ctrl.CreateEnvironment).Methods("POST")
	a.Router.HandleFunc("/api/environments", ctrl.GetEnvironments).Methods("GET")
	a.Router.HandleFunc("/api/environment/{id:[0-9]+}", ctrl.GetEnvironment).Methods("GET")
	a.Router.HandleFunc("/api/environment/{id:[0-9]+}", ctrl.UpdateEnvironment).Methods("PUT")
	a.Router.HandleFunc("/api/environment/{id:[0-9]+}", ctrl.DeleteEnvironment).Methods("DELETE")

	//model.TeamUser struct
	a.Router.HandleFunc("/api/team/user", ctrl.CreateTeamUser).Methods("POST")
	a.Router.HandleFunc("/api/team/{teamid:[0-9]+}/users", ctrl.GetTeamUsers).Methods("GET")
	a.Router.HandleFunc("/api/user/{userid:[0-9]+}/teams", ctrl.GetUserTeams).Methods("GET")
	a.Router.HandleFunc("/api/team/user/{id:[0-9]+}", ctrl.UpdateTeamUser).Methods("PUT")
	a.Router.HandleFunc("/api/team/user/{id:[0-9]+}", ctrl.DeleteTeamUser).Methods("DELETE")

	//model.TeamProject struct
	a.Router.HandleFunc("/api/team/project", ctrl.CreateTeamProject).Methods("POST")
	a.Router.HandleFunc("/api/team/{teamid:[0-9]+}/projects", ctrl.GetTeamProjects).Methods("GET")
	a.Router.HandleFunc("/api/project/{projectid:[0-9]+}/teams", ctrl.GetProjectTeams).Methods("GET")
	a.Router.HandleFunc("/api/team/project/{id:[0-9]+}", ctrl.UpdateTeamProject).Methods("PUT")
	a.Router.HandleFunc("/api/team/project/{id:[0-9]+}", ctrl.DeleteTeamProject).Methods("DELETE")

	//model.EnvironmentInstance stuct
	a.Router.HandleFunc("/api/environment/instance", ctrl.CreateEnvironmentInstance).Methods("POST")
}

//RunApplication - Start the HTTP server
func (a *App) RunApplication(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
