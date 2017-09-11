package model

import (
	"database/sql"
)

//Environment  (TYPE)
type Environment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

//Environments (TYPE)
type Environments struct {
	Environments []*Environment `json:"environments"`
}

//DoesEnvironmentResourceExist (POST)
func DoesEnvironmentResourceExist(environment *Environment) bool {

	err := db.QueryRow("SELECT id, name, description, type FROM sbm_environments WHERE name=$1 and type=$2",
		environment.Name, environment.Type).Scan(&environment.ID, &environment.Name, &environment.Description, &environment.Type)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//CreateEnvironment (POST)
func CreateEnvironment(environment *Environment) error {

	err := db.QueryRow(
		"INSERT INTO sbm_environments(name, description, type) VALUES($1, $2, $3) RETURNING id",
		environment.Name, environment.Description, environment.Type).Scan(&environment.ID)

	if err != nil {
		return err
	}

	return nil
}

//GetEnvironments (GET)
func GetEnvironments() ([]Environment, error) {
	rows, err := db.Query("SELECT id, name, description, type FROM sbm_environments")

	if err != nil {
		return nil, err
	}

	environments := []Environment{}

	for rows.Next() {
		defer rows.Close()

		var e Environment
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Type); err != nil {
			return nil, err
		}
		environments = append(environments, e)
	}

	return environments, nil
}

//GetEnvironment (GET)
func GetEnvironment(environment *Environment) error {
	return db.QueryRow("SELECT name, description, type FROM sbm_environments WHERE id=$1",
		environment.ID).Scan(&environment.Name, &environment.Description, &environment.Type)
}

//UpdateEnvironment (PUT)
func UpdateEnvironment(environment *Environment) error {
	_, err :=
		db.Exec("UPDATE sbm_environments SET name=$1, description=$2, type=$3 WHERE id=$4",
			environment.Name, environment.Description, environment.Type, environment.ID)

	return err
}

//DeleteEnvironment (DELETE)
func DeleteEnvironment(environment *Environment) error {
	_, err := db.Exec("DELETE FROM sbm_environments WHERE id=$1", environment.ID)

	return err
}
