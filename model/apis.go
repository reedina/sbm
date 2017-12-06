package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetTeam - Validate Team Object
func GetTeam(teamID int) Team {

	url := "http://localhost:4040/api/team/" + strconv.Itoa(teamID)

	teamClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := teamClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	team := Team{}
	jsonErr := json.Unmarshal(body, &team)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return Team{}
	}

	return team
}
