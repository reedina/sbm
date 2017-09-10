package ctrl

import (
	"encoding/json"
	"net/http"

	"github.com/reedina/buildenvironment/model"
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

	// Validate if User exists based on email address
	if model.DoesUserResourceExist(&environmentInstance.User) == false {
		model.CreateUser(&environmentInstance.User) // check for minimum values & error handling: FIXME
	}

	// Validate if Team exists based on name
	if model.DoesTeamResourceExist(&environmentInstance.Team) == false {
		model.CreateTeam(&environmentInstance.Team) // check for minimum values & error handling: FIXME
	}

	// Validate User is a member of Team
	teamuser := model.TeamUser{TeamID: environmentInstance.Team.ID, UserID: environmentInstance.User.ID}
	if model.DoesTeamUserResourceExist(&teamuser) == false {
		model.CreateTeamUser(&teamuser)
	}

	// Validate Project exists based on name
	if model.DoesProjectResourceExist(&environmentInstance.Project) == false {
		model.CreateProject(&environmentInstance.Project) // check for minimum values & error handling: FIXME
	}

	// Validate Team has Project
	teamproject := model.TeamProject{TeamID: environmentInstance.Team.ID, ProjectID: environmentInstance.Project.ID}
	if model.DoesTeamProjectResourceExist(&teamproject) == false {
		model.CreateTeamProject(&teamproject)
	}
	// Validate Environment exists based on name and type
	if model.DoesEnvironmentResourceExist(&environmentInstance.Environment) == false {
		model.CreateEnvironment(&environmentInstance.Environment) // check for minimum values & error handling: FIXME
	}
	// Validate parsing date time string to UTC epoch
	utcEpoch, err := model.ConvertStringToUtcEpoch(environmentInstance.ExpirationStringTime)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Expiration Time format: "+environmentInstance.ExpirationStringTime)
		return
	}
	//Validate Expiration time is atleast 24 hours from now
	if model.IsUtcEpoch24HoursFromNow(utcEpoch) == false {
		respondWithError(w, http.StatusBadRequest, "Expiration Time less than 24 hours from now: "+environmentInstance.ExpirationStringTime)
		return
	}
	environmentInstance.ExpirationUtcEpoch = utcEpoch

	// Resources exist to create an environment instance
	if err := model.CreateEnvironmentInstance(&environmentInstance); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, environmentInstance)
}
