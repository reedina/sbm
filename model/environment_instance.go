package model

import (
	"database/sql"
)

//EnvironmentInstance  (TYPE)
type EnvironmentInstance struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	DelectionTime string      `json:"deletion_time"` // "2017-12-01 22:43:22" - Local time
	Environment   Environment `json:"environment"`
	Team          Team        `json:"team"`
}

//Team  (TYPE)
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//EnvironmentInstances (TYPE)
type EnvironmentInstances struct {
	EnvironmentInstances []*EnvironmentInstance `json:"environment_instances"`
}

//DoesEnvironmentInstanceNameExist (POST)
func DoesEnvironmentInstanceNameExist(environmentInstance *EnvironmentInstance) bool {

	err := db.QueryRow("SELECT id, name FROM sbm_environment_instances WHERE name=$1",
		environmentInstance.Name).Scan(&environmentInstance.ID, &environmentInstance.Name)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//CreateEnvironmentInstance - Store in database
func CreateEnvironmentInstance(environmentInstance *EnvironmentInstance) error {
	err := db.QueryRow(
		"INSERT INTO sbm_environment_instances(name, environment_id, deletion_time, team_id) VALUES($1, $2, $3, $4) RETURNING id",
		environmentInstance.Name, environmentInstance.Environment.ID, environmentInstance.DelectionTime, environmentInstance.Team.ID).Scan(&environmentInstance.ID)

	if err != nil {
		return err
	}

	teamObject := GetTeam(environmentInstance.Team.ID)
	environmentInstance.Team.Name = teamObject.Name

	return nil
}

//GetEnvironmentInstances (GET)
func GetEnvironmentInstances() ([]EnvironmentInstance, error) {
	rows, err := db.Query("SELECT sbm_environment_instances.id, sbm_environment_instances.name, sbm_environments.name, " +
		"sbm_environment_instances.deletion_time, sbm_environments.type, sbm_environments.ID, sbm_environment_instances.team_id FROM sbm_environment_instances " +
		"inner join sbm_environments on environment_id = sbm_environments.id")

	if err != nil {
		return nil, err
	}

	environmentInstances := []EnvironmentInstance{}

	for rows.Next() {
		defer rows.Close()

		var e EnvironmentInstance
		if err := rows.Scan(&e.ID, &e.Name, &e.Environment.Name, &e.DelectionTime, &e.Environment.Type, &e.Environment.ID, &e.Team.ID); err != nil {
			return nil, err
		}
		teamObject := GetTeam(e.Team.ID)
		e.Team.Name = teamObject.Name
		environmentInstances = append(environmentInstances, e)
	}

	return environmentInstances, nil
}
