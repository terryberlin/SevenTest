package main

import (
	"log"

	"github.com/quikserve/SevenTest/locations"
	"github.com/quikserve/SevenTest/shifts"
	"github.com/quikserve/SevenTest/timepunchsync"
	"github.com/quikserve/SevenTest/users"
)

func main() {
	doLocations := false
	doUsers := false
	doShifts := false
	doTimePunchSync := true

	if doLocations {
		log.Println("Call LOCATIONS package")
		locations.MainLocations()
	}

	if doUsers {
		log.Println("Call USERS package")
		users.MainUsers()
	}

	if doShifts {
		log.Println("Call SHIFTS package")
		shifts.MainShifts()
	}

	if doTimePunchSync {
		log.Println("Call TIME PUNCH SYNC  package")
		timepunchsync.MainTimePunchSync()
	}
}
