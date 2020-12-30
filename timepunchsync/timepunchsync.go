package timepunchsync

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/quikserve/SevenTest/db"
)

type (
	//TimePunchList is a structure of cities from SQL query.
	TimePunchList struct {
		PunchID         *string `db:"PunchID" json:"PunchID"`
		PunchJSON       *string `db:"PunchJSON" json:"PunchJSON"`
		APIKey          *string `db:"APIKey" json:"APIKey"`
		EmployeeHoursID *string `db:"EmployeeHoursID" json:"EmployeeHoursID"`
		Curl            *string `db:"Curl" json:"Curl"`
		TimePunchID     *string `db:"TimePunchID" json:"TimePunchID"`
		ClockedIn       *string `db:"ClockedIn" json:"ClockedIn"`
		ClockedOut      *string `db:"ClockedOut" json:"ClockedOut"`
	}

	//ListPosts is a structure of weather posts.
	ListPosts struct {
		Result *string `db:"Result"`
	}
)

//MainTimePunchSync is a func
func MainTimePunchSync() {

	var PunchID string
	var PunchJSON string
	var key string
	var hoursID string
	var curl string
	var TimePunchID string
	var ClockedIn string
	var ClockedOut string

	//cache punches
	TimePunchLists := []TimePunchList{}
	sql2 := `exec crm.dbo.seven_shifts_cache_time_punches`

	err2 := db.MyDB().Select(&TimePunchLists, sql2)
	if err2 != nil {
		log.Println(err2)
	}

	for i := range TimePunchLists {
		PunchID = fmt.Sprint(*TimePunchLists[i].PunchID)
		PunchJSON = fmt.Sprint(*TimePunchLists[i].PunchJSON)
		key = fmt.Sprint(*TimePunchLists[i].APIKey)
		hoursID = fmt.Sprint(*TimePunchLists[i].EmployeeHoursID)
		curl = fmt.Sprint(*TimePunchLists[i].Curl)
		TimePunchID = fmt.Sprint(*TimePunchLists[i].TimePunchID)
		ClockedIn = fmt.Sprint(*TimePunchLists[i].ClockedIn)
		ClockedOut = fmt.Sprint(*TimePunchLists[i].ClockedOut)

		//post results to 7Shifts
		log.Println("pushing to 7shifts API", PunchID, hoursID, curl)
		ContactAPI(PunchJSON, key, hoursID, curl, TimePunchID, ClockedIn, ClockedOut)

		//clear from SQL cache table
		//log.Println("deleting from SQL cache table", PunchID)
	}

}

//ContactAPI is a function that contacts the weather API.
func ContactAPI(PunchJSON string, key string, hoursID string, curl string, TimePunchID string, ClockedIn string, ClockedOut string) {

	url := fmt.Sprintf("https://api.7shifts.com/v1/time_punches")

	if curl == "PUT" {
		url = fmt.Sprintf("https://api.7shifts.com/v1/time_punches/%s", TimePunchID)
	}

	c := exec.Command("curl", "-X", curl, "-u", key, "-d", PunchJSON, url)
	c.Stdout = os.Stdout
	outfile, err1 := os.Create("./punch.json")
	if err1 != nil {
		fmt.Println("Error:", err1)
	}

	c.Stdout = outfile
	//c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	content, err := ioutil.ReadFile("./punch.json")
	if err != nil {
		fmt.Println(err)
	}

	text := string(content)

	PostToSQL(text, key, hoursID, ClockedIn, ClockedOut)

}

//PostToSQL is a function for posting to SQL
func PostToSQL(text string, key string, hoursID string, ClockedIn string, ClockedOut string) {
	//log.Println("Posting", text, hoursID, ClockedIn, ClockedOut)

	listposts := []ListPosts{}
	sqlQ := `exec crm.dbo.import_seven_shifts_punch_map $1, $2, $3, $4, $5`

	errQ := db.MyDB().Select(&listposts, sqlQ, text, key, hoursID, ClockedIn, ClockedOut)
	if errQ != nil {
		log.Println(errQ)
	} else {
		log.Println(`exec crm.dbo.import_seven_shifts_punch_map '` + text + `', '5', ` + hoursID + `, '` + ClockedIn + `', '` + ClockedOut + `'`)
	}
}
