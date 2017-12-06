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
	Project       Project     `json:"project"`
}

//Team  (TYPE)
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Project  (TYPE)
type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Team Team   `json:"team"`
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
		"INSERT INTO sbm_environment_instances(name, environment_id, deletion_time, team_id, team_name, project_id, project_name) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		environmentInstance.Name, environmentInstance.Environment.ID, environmentInstance.DelectionTime, environmentInstance.Team.ID, environmentInstance.Team.Name,
		environmentInstance.Project.ID, environmentInstance.Project.Name).Scan(&environmentInstance.ID)

	if err != nil {
		return err
	}

	return nil
}

//GetEnvironmentInstances (GET)
func GetEnvironmentInstances() ([]EnvironmentInstance, error) {
	rows, err := db.Query("SELECT sbm_environment_instances.id, sbm_environment_instances.name, sbm_environments.name, " +
		"sbm_environment_instances.deletion_time, sbm_environments.type, sbm_environments.ID, sbm_environment_instances.team_id, " +
		"sbm_environment_instances.team_name, sbm_environment_instances.project_id, sbm_environment_instances.project_name " +
		"FROM sbm_environment_instances " +
		"inner join sbm_environments on environment_id = sbm_environments.id")

	if err != nil {
		return nil, err
	}

	environmentInstances := []EnvironmentInstance{}

	for rows.Next() {
		defer rows.Close()

		var e EnvironmentInstance
		if err := rows.Scan(&e.ID, &e.Name, &e.Environment.Name, &e.DelectionTime, &e.Environment.Type, &e.Environment.ID, &e.Team.ID,
			&e.Team.Name, &e.Project.ID, &e.Project.Name); err != nil {
			return nil, err
		}

		e.Project.Team.ID = e.Team.ID
		e.Project.Team.Name = e.Team.Name

		environmentInstances = append(environmentInstances, e)
	}

	return environmentInstances, nil
}
