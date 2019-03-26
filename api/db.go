package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "postgres://checky:checky@postgres/checky?sslmode=disable"

func dbInit() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE endpoints (id SERIAL, name VARCHAR(50), url VARCHAR(255), status INTEGER, responsetime INTEGER)")
	if err != nil {
		log.Println("Error creating endpoints table:", err)
	}
	fmt.Println("DB initialized")
}

func dbSelectAllEndpoints() []byte {
	// connect to DB
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Cannot connect to DB:", err)
		return nil
	}
	defer db.Close()

	// run query
	rows, err := db.Query("SELECT * FROM endpoints")
	if err != nil {
		log.Println("Error selecting all rows endpoints table:", err)
		return nil
	}
	defer rows.Close()

	// parse rows
	var endpoints []endpoint
	for rows.Next() {
		var e endpoint
		err = rows.Scan(&e.Name, &e.URL, &e.Status)
		if err != nil {
			log.Println("Error scanning endpoint rows:", err)
			return nil
		}
		endpoints = append(endpoints, e)
	}

	// marshal JSON
	endpointsJSON, err := json.Marshal(endpoints)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return nil
	}

	return endpointsJSON
}

func dbSelectSingleEndpoint(name string) []byte {
	// Connect to DB
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Cannot connect to DB:", err)
		return nil
	}
	defer db.Close()

	// run query
	row := db.QueryRow("SELECT * FROM endpoints WHERE id = $1", name)
	if err != nil {
		log.Println("Error selecting single row from endpoints table:", err)
		return nil
	}

	// parse SQL row
	var e endpoint
	row.Scan(&e.Name, &e.URL, &e.Status)

	// marshal JSON
	eJSON, err := json.Marshal(e)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return nil
	}

	return eJSON
}
