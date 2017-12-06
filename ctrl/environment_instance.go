package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/reedina/sbm/model"
)

//CreateEnvironmentInstance (POST)
func CreateEnvironmentInstance(w http.ResponseWriter, r *http.Request) {
	var environmentInstance model.EnvironmentInstance

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&environmentInstance); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate Team exists based on ID
	team := model.GetTeam(environmentInstance.Team.ID)

	if team.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Team ID does not exist")
		return
	}

	// Set Team Name
	environmentInstance.Team.Name = team.Name

	// Validate Project exists based on ID
	project := model.GetProject(environmentInstance.Project.ID)

	if project.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Project ID does not exist")
		return
	}

	// Set Project Name
	environmentInstance.Project.Name = project.Name

	// Validate this Project is assigned to THIS team
	currentTeamID := team.ID
	currentProjectTeamID := project.Team.ID

	if currentTeamID != currentProjectTeamID {

		respondWithError(w, http.StatusBadRequest, "Project ID does not exist for Team ID")
		return
	}

	// Validate Environment exists based on ID
	if model.DoesEnvironmentIDExist(&environmentInstance.Environment) == false {

		respondWithError(w, http.StatusBadRequest, "Environment ID does not exist")
		return
	}

	// Validate Environment Instance name already exists
	if model.DoesEnvironmentInstanceNameExist(&environmentInstance) == true {

		respondWithError(w, http.StatusBadRequest, "Environment Instance name already exists")
		return
	}
	// Resources exist to create an environment instance
	if err := model.CreateEnvironmentInstance(&environmentInstance); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, environmentInstance)
}

//GetEnvironmentInstances  (GET)
func GetEnvironmentInstances(w http.ResponseWriter, r *http.Request) {

	environments, err := model.GetEnvironmentInstances()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, environments)
}
