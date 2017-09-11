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

	// Validate Environment exists based on name and type
	if model.DoesEnvironmentResourceExist(&environmentInstance.Environment) == false {

		respondWithError(w, http.StatusBadRequest, "Environment resource does not exist")
		return
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
