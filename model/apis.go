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

// GetProject - Validate Project Object
func GetProject(projectID int) Project {

	url := "http://localhost:4040/api/project/" + strconv.Itoa(projectID)

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

	project := Project{}
	jsonErr := json.Unmarshal(body, &project)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return Project{}
	}

	return project
}
