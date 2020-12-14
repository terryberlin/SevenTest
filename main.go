package main

import (
	"log"

	"github.com/quikserve/SevenTest/locations"
	"github.com/quikserve/SevenTest/users"
)

func main() {
	log.Println("Call locations package")
	locations.MainLocations()

	log.Println("Call usres package")
	users.MainUsers()
}
