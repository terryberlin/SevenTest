package main

import (
	"log"

	"github.com/quikserve/SevenTest/locations"
	"github.com/quikserve/SevenTest/shifts"
	"github.com/quikserve/SevenTest/users"
)

func main() {
	doLocations := false
	doUsers := false
	doShifts := false

	if doLocations {
		log.Println("Call locations package")
		locations.MainLocations()
	}

	if doUsers {
		log.Println("Call users package")
		users.MainUsers()
	}

	if doShifts {
		log.Println("Call shifts package")
		shifts.MainShifts()
	}
}
