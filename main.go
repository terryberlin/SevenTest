package main

import (
	"log"

	_ "github.com/quikserve/SevenTest/locations"
)

//"github.com/quikserve/7ShiftTest/locations"

func main() {
	log.Println("Call locations package")
	locations.mainLocations()
}
