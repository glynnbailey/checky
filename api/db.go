package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "postgres://checky:checky@postgres/checky?sslmode=disable"

func dbSelectAllEndpoints() []endpoint {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM endpoints")
	if err != nil {
		log.Println("Error selecting all rows endpoints table:", err)
	}
	defer rows.Close()

	var endpoints []endpoint
	for rows.Next() {
		var e endpoint
		err = rows.Scan(&e.Name, &e.URL, &e.Status)
		if err != nil {
			log.Println("Error scanning endpoint rows:", err)
		}
		endpoints = append(endpoints, e)
	}
	return endpoints
}

func dbSelectSingleEndpoints(name string) endpoint {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM endpoints WHERE name = $1", name)
	if err != nil {
		log.Println("Error selecting single row from endpoints table:", err)
	}

	var e endpoint
	row.Scan(&e.Name, &e.URL, &e.Status)
	return e
}
