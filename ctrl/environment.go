package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/sbm/model"
)

//CreateEnvironment (POST)
func CreateEnvironment(w http.ResponseWriter, r *http.Request) {
	var environment model.Environment

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&environment); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Environment Attribute Resource Exist ?
	if model.DoesEnvironmentResourceExist(&environment) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateEnvironment(&environment); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, environment)
}

//GetEnvironments  (GET)
func GetEnvironments(w http.ResponseWriter, r *http.Request) {

	environments, err := model.GetEnvironments()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, environments)
}

//GetEnvironment (GET)
func GetEnvironment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Environment ID")
		return
	}

	environment := model.Environment{ID: id}
	if err := model.GetEnvironment(&environment); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Environment not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, environment)
}

//UpdateEnvironment (PUT)
func UpdateEnvironment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Environment ID")
		return
	}

	var environment model.Environment

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&environment); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	environment.ID = id

	if err := model.UpdateEnvironment(&environment); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, environment)
}

//DeleteEnvironment (DELETE)
func DeleteEnvironment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Environment ID")
		return
	}
	environment := model.Environment{ID: id}

	if err := model.DeleteEnvironment(&environment); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
