package main

import (
	"log"

	"github.com/quikserve/SevenTest/companies"
	"github.com/quikserve/SevenTest/departments"
	"github.com/quikserve/SevenTest/locations"
	"github.com/quikserve/SevenTest/roles"
	"github.com/quikserve/SevenTest/shifts"
	"github.com/quikserve/SevenTest/timepunches"
	"github.com/quikserve/SevenTest/timepunchsync"
	"github.com/quikserve/SevenTest/timepunchsyncintraday"
	"github.com/quikserve/SevenTest/timepunchsyncintradayreset"
	"github.com/quikserve/SevenTest/users"
	"github.com/quikserve/SevenTest/userssync"
)

func main() {

	doCompanies := false
	doLocations := false
	doDepartments := false
	doRoles := false
	doUsers := true
	doShifts := false
	doTimePunches := true

	doUsersSync := false
	doTimePunchSync := false
	doTimePunchSyncIntraday := false

	doTimePunchSyncIntradayReset := false

	if doCompanies {
		log.Println("Call COMPANIES package")
		companies.MainCompanies()
	}

	if doLocations {
		log.Println("Call LOCATIONS package")
		locations.MainLocations()
	}

	if doDepartments {
		log.Println("Call DEPARTMENTS package")
		departments.MainDepartments()
	}

	if doRoles {
		log.Println("Call ROLES package")
		roles.MainRoles()
	}

	if doUsers {
		log.Println("Call USERS package")
		users.MainUsers()
	}

	if doShifts {
		log.Println("Call SHIFTS package")
		shifts.MainShifts()
	}

	if doTimePunches {
		log.Println("Call TIMEPUNCHES  package")
		timepunches.MainTimePunches()
	}

	if doTimePunchSync {
		log.Println("Call TIME PUNCH SYNC  package")
		timepunchsync.MainTimePunchSync()
	}

	if doTimePunchSyncIntraday {
		log.Println("Call TIME PUNCH SYNC INTRADAY  package")
		timepunchsyncintraday.MainTimePunchSyncIntraday()
	}

	if doTimePunchSyncIntradayReset {
		log.Println("Call TIME PUNCH SYNC INTRADAY RESET  package")
		timepunchsyncintradayreset.MainTimePunchSyncIntradayReset()
	}

	if doUsersSync {
		log.Println("Call USERS SYNC  package")
		userssync.MainUsersSync()
	}
}
