package departments

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/quikserve/SevenTest/db"
)

type (
	//DepartmentList is a structure of cities from SQL query.
	DepartmentList struct {
		LocationID *string `db:"LocationID" json:"LocationID"`
		APIKey     *string `db:"APIKey" json:"APIKey"`
	}

	//ListPosts is a structure of weather posts.
	ListPosts struct {
		Result *string `db:"Result"`
	}
)

//MainDepartments is a function
func MainDepartments() {

	var key string

	sql2 := `
        select location_id as LocationID, api_key as APIKey
		from quikserve.dbo.seven_shifts_locations s
		where s.active in (1)
	`

	DepartmentLists := []DepartmentList{}
	err := db.MyDB().Select(&DepartmentLists, sql2)
	if err != nil {
		log.Println(err)
	}

	var LocationID string
	for i := range DepartmentLists {
		LocationID = fmt.Sprint(*DepartmentLists[i].LocationID)
		key = fmt.Sprint(*DepartmentLists[i].APIKey)
		ContactAPI(LocationID, key)
	}

}

//ContactAPI is a function that contacts the weather API.
func ContactAPI(LocationID string, key string) {

	log.Println("Getting DEPARTMENTS for Location:", LocationID)

	url := fmt.Sprintf("https://api.7shifts.com/v1/departments/?location_id=%s", LocationID)

	c := exec.Command("curl", "-u", key, url)
	c.Stdout = os.Stdout
	outfile, err1 := os.Create("./departments.json")
	if err1 != nil {
		fmt.Println("Error:", err1)
	}

	c.Stdout = outfile
	//c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	content, err := ioutil.ReadFile("./departments.json")
	if err != nil {
		fmt.Println(err)
	}

	text := string(content)

	PostIt(text, key, LocationID)

}

//PostIt is a function for posting to SQL
func PostIt(text string, key string, LocationID string) {

	listposts := []ListPosts{}
	sql := `exec crm.dbo.import_seven_shifts_departments $1`

	err := db.MyDB().Select(&listposts, sql, text)
	if err != nil {
		log.Println(err)
	}
}
