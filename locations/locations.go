package locations

import (
	"io/ioutil"
	"os"
	"os/exec"

	"fmt"
	"log"

	"github.com/quikserve/SevenTest/db"
)

type (
	//LocationList is a structure of cities from SQL query.
	LocationList struct {
		LocationID *string `db:"LocationID" json:"LocationID"`
		APIKey     *string `db:"APIKey" json:"APIKey"`
	}

	//ListPosts is a structure of weather posts.
	ListPosts struct {
		Result *string `db:"Result"`
	}
)

//MainLocations is a function to import locations
func MainLocations() {

	var key string
	// var status string

	// status = "reveal"

	// listposts := []ListPosts{}
	// sql1 := `exec crm.dbo.key_status $1`

	// err1 := db.MyDB().Select(&listposts, sql1, status)
	// if err1 != nil {
	// 	log.Println(err1)
	// }

	sql2 := `
        select location_id as LocationID, api_key as APIKey
		from quikserve.dbo.seven_shifts_locations s
		where s.active in (1)
	`

	LocationLists := []LocationList{}
	err2 := db.MyDB().Select(&LocationLists, sql2)
	if err2 != nil {
		log.Println(err2)
	}

	var LocationID string
	for i := range LocationLists {
		LocationID = fmt.Sprint(*LocationLists[i].LocationID)
		key = fmt.Sprint(*LocationLists[i].APIKey)
		ContactAPI(LocationID, key)
	}

	// status = "disguise"

	// listposts = []ListPosts{}
	// sql3 := `exec crm.dbo.key_status $1`

	// err3 := db.MyDB().Select(&listposts, sql3, status)
	// if err3 != nil {
	// 	log.Println(err3)
	// }
}

//ContactAPI is a function that contacts the weather API.
func ContactAPI(LocationID string, key string) {

	log.Println("Getting details for LOCATION:", LocationID)

	url := "https://api.7shifts.com/v1/locations"

	c := exec.Command("curl", "-u", key, url)
	c.Stdout = os.Stdout
	outfile, err1 := os.Create("./locations.json")
	if err1 != nil {
		fmt.Println("Error:", err1)
	}

	c.Stdout = outfile
	//c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	content, err := ioutil.ReadFile("./locations.json")
	if err != nil {
		fmt.Println(err)
	}

	text := string(content)

	PostIt(text, key)

}

//PostIt is a function for posting to SQL
func PostIt(text string, key string) {

	listposts := []ListPosts{}
	sql := `exec crm.dbo.import_seven_shifts_locations $1, $2`

	err := db.MyDB().Select(&listposts, sql, text, key)
	if err != nil {
		log.Println(err)
	}
}
