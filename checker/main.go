package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Checker service starting")

	for {
		fmt.Println("Checking endpoint status")
		// collect urls from DB
		endpoints := dbSelectAllEndpoints()

		// check urls are available and update the DB data
		for _, e := range endpoints {
			e.checkStatus()
			e.dbUpdateEndpointStatus()
		}

		// sleep
		fmt.Println("Endpoint check complete")
		time.Sleep(time.Minute)
	}
}
