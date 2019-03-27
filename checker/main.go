package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(time.Second) // wait for API to come up
	fmt.Println("Checker service starting")

	for {
		fmt.Println("Checking endpoint status")
		// collect urls from API
		endpoints := apiSelectAllEndpoints()

		// check urls are available and update the DB data
		for _, e := range endpoints {
			e.update()
		}

		// sleep
		fmt.Println("Endpoint check complete")
		time.Sleep(time.Minute)
	}
}
