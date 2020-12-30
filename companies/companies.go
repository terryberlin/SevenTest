package companies

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/quikserve/SevenTest/db"
)

type (
	//CompanyList is a structure of cities from SQL query.
	CompanyList struct {
		LocationID *string `db:"LocationID" json:"LocationID"`
		APIKey     *string `db:"APIKey" json:"APIKey"`
	}

	//ListPosts is a structure of weather posts.
	ListPosts struct {
		Result *string `db:"Result"`
	}
)

//MainCompanies is a function
func MainCompanies() {

	var key string

	sql2 := `
        select distinct api_key as LocationID, api_key as APIKey
		from quikserve.dbo.seven_shifts_locations s
		where s.active in (1)
	`

	CompanyLists := []CompanyList{}
	err := db.MyDB().Select(&CompanyLists, sql2)
	if err != nil {
		log.Println(err)
	}

	var LocationID string
	for i := range CompanyLists {
		LocationID = fmt.Sprint(*CompanyLists[i].LocationID)
		key = fmt.Sprint(*CompanyLists[i].APIKey)
		ContactAPI(LocationID, key)
	}

}

//ContactAPI is a function that contacts the weather API.
func ContactAPI(LocationID string, key string) {

	log.Println("Getting COMPANIES for Location:")

	url := fmt.Sprintf("https://api.7shifts.com/v1/companies")

	c := exec.Command("curl", "-u", key, url)
	c.Stdout = os.Stdout
	outfile, err1 := os.Create("./companies.json")
	if err1 != nil {
		fmt.Println("Error:", err1)
	}

	c.Stdout = outfile
	//c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	content, err := ioutil.ReadFile("./companies.json")
	if err != nil {
		fmt.Println(err)
	}

	text := string(content)

	PostIt(text, key, LocationID)

}

//PostIt is a function for posting to SQL
func PostIt(text string, key string, LocationID string) {
	//log.Println("95")
	listposts := []ListPosts{}
	sql := `exec crm.dbo.import_seven_shifts_companies $1`

	err := db.MyDB().Select(&listposts, sql, text)
	if err != nil {
		log.Println(err)
	}
}
