package main

import (
	"database/sql"
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

	_, err = db.Exec("CREATE TABLE endpoints (name VARCHAR(50), url VARCHAR(255), status INTEGER)")
	if err != nil {
		log.Println("Error creating endpoints table:", err)
	}
}

func dbSelectAllEndpoints() []endpoint {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM endpoints")
	if err != nil {
		log.Println("Error creating endpoints table:", err)
	}
	defer rows.Close()

	var endpoints []endpoint
	for rows.Next() {
		var e endpoint
		err = rows.Scan(&e.Name, &e.URL, &e.Status)
		if err != nil {
			log.Println("Error querying endpoint rows:", err)
		}
		endpoints = append(endpoints, e)
	}
	return endpoints
}

func (e *endpoint) dbUpdateEndpointStatus() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Cannot connect to DB:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE endpoints SET status = $1 WHERE name = $2", e.Status, e.Name)
	if err != nil {
		log.Println("Error updating endpoints table:", err)
	}
}
