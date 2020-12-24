package userssync

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/quikserve/SevenTest/db"
)

type (
	//LocationList is a structure of cities from SQL query.
	LocationList struct {
		ID         *string `db:"ID" json:"ID"`
		EmployeeID *string `db:"EmployeeID" json:"EmployeeID"`
		UserJSON   *string `db:"UserJSON" json:"UserJSON"`
		APIKey     *string `db:"APIKey" json:"APIKey"`
		Curl       *string `db:"Curl" json:"Curl"`
		UserID     *string `db:"UserID" json:"UserID"`
	}

	//ListPosts is a structure of weather posts.
	ListPosts struct {
		Result *string `db:"Result"`
	}
)

//MainUsersSync is a func.
func MainUsersSync() {

	var ID string
	var UserJSON string
	var key string
	var empID string
	var curl string
	var userID string

	//var store string
	//store = "9367"

	//cache punches
	LocationLists := []LocationList{}
	//sql2 := `exec crm.dbo.seven_shifts_cache_users $1`
	sql2 := `exec crm.dbo.seven_shifts_cache_users`

	err2 := db.MyDB().Select(&LocationLists, sql2)
	if err2 != nil {
		log.Println(err2)
	}

	for i := range LocationLists {
		ID = fmt.Sprint(*LocationLists[i].ID)
		UserJSON = fmt.Sprint(*LocationLists[i].UserJSON)
		key = fmt.Sprint(*LocationLists[i].APIKey)
		empID = fmt.Sprint(*LocationLists[i].EmployeeID)
		curl = fmt.Sprint(*LocationLists[i].Curl)
		userID = fmt.Sprint(*LocationLists[i].UserID)

		//post results to 7Shifts
		log.Println("pushing to 7shifts API", ID, empID, curl)
		ContactAPI(UserJSON, key, empID, curl, userID)

		//clear from SQL cache table
		//log.Println("deleting from SQL cache table", ID)
	}

}

//ContactAPI is a function that contacts the weather API.
func ContactAPI(UserJSON string, key string, empID string, curl string, userID string) {

	url := fmt.Sprintf("https://api.7shifts.com/v1/users")

	if curl == "PUT" {
		url = fmt.Sprintf("https://api.7shifts.com/v1/users/%s", userID)
	}

	c := exec.Command("curl", "-X", curl, "-u", key, "-d", UserJSON, url)
	c.Stdout = os.Stdout
	outfile, err1 := os.Create("./user.json")
	if err1 != nil {
		fmt.Println("Error:", err1)
	}

	c.Stdout = outfile
	//c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	content, err := ioutil.ReadFile("./user.json")
	if err != nil {
		fmt.Println(err)
	}

	text := string(content)

	PostToSQL(text, key, empID, userID)

}

//PostToSQL is a function for posting to SQL
func PostToSQL(text string, key string, empID string, userID string) {
	//log.Println("Posting", text, empID, ClockedIn, ClockedOut)

	listposts := []ListPosts{}
	sqlQ := `exec crm.dbo.import_seven_shifts_user_map $1, $2, $3, $4`

	errQ := db.MyDB().Select(&listposts, sqlQ, text, key, empID, userID)
	if errQ != nil {
		log.Println(errQ)
	}
}
