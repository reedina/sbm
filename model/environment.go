package model

import (
	"database/sql"
	"fmt"
)

//Environment  (TYPE)
type Environment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

//Environments (TYPE)
type Environments struct {
	Environments []*Environment `json:"environments"`
}

//DoesEnvironmentIDExist (POST)
func DoesEnvironmentIDExist(environment *Environment) bool {

	err := db.QueryRow("SELECT id  FROM sbm_environments WHERE id=?", environment.ID).Scan(&environment.ID)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesEnvironmentResourceExist (POST)
func DoesEnvironmentResourceExist(environment *Environment) bool {

	err := db.QueryRow("SELECT id, name, type, url FROM sbm_environments WHERE name=? and type=?",
		environment.Name, environment.Type).Scan(&environment.ID, &environment.Name, &environment.Type, &environment.URL)

	fmt.Println(err)
	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//CreateEnvironment (POST)
func CreateEnvironment(environment *Environment) error {
	/*
		err := db.QueryRow(
			"INSERT INTO sbm_environments(name, type, url) VALUES($1, $2, $3) RETURNING id",
			environment.Name, environment.Type, environment.URL).Scan(&environment.ID)
	*/
	res, err := db.Exec("INSERT INTO sbm_environments(name, type, url) VALUES(?, ?, ?)",
		environment.Name, environment.Type, environment.URL)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	environment.ID = int(id)

	return nil
}

//GetEnvironments (GET)
func GetEnvironments() ([]Environment, error) {
	rows, err := db.Query("SELECT id, name, type, url FROM sbm_environments")

	if err != nil {
		return nil, err
	}

	environments := []Environment{}

	for rows.Next() {
		defer rows.Close()

		var e Environment
		if err := rows.Scan(&e.ID, &e.Name, &e.Type, &e.URL); err != nil {
			return nil, err
		}
		environments = append(environments, e)
	}

	return environments, nil
}

//GetEnvironment (GET)
func GetEnvironment(environment *Environment) error {
	return db.QueryRow("SELECT name, type, URL FROM sbm_environments WHERE id=?",
		environment.ID).Scan(&environment.Name, &environment.Type, &environment.URL)
}

//UpdateEnvironment (PUT)
func UpdateEnvironment(environment *Environment) error {
	_, err :=
		db.Exec("UPDATE sbm_environments SET name=?, type=?, url=? WHERE id=?",
			environment.Name, environment.Type, environment.URL, environment.ID)

	return err
}

//DeleteEnvironment (DELETE)
func DeleteEnvironment(environment *Environment) error {
	_, err := db.Exec("DELETE FROM sbm_environments WHERE id=?", environment.ID)

	return err
}
