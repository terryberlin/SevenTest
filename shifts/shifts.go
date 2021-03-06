package shifts

import (
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"fmt"
	"log"

	"github.com/quikserve/SevenTest/db"
)

type (
	//UserList is a structure of cities from SQL query.
	UserList struct {
		UserID *string `db:"UserID" json:"UserID"`
		APIKey *string `db:"APIKey" json:"APIKey"`
	}

	//ListPosts is a structure of weather posts.
	ListPosts struct {
		Result *string `db:"Result"`
	}
)

func MainShifts() {

	var key string

	sql2 := `
		select distinct s.[user_id] as UserID, s.api_key as APIKey
		from quikserve.dbo.seven_shifts_users s 
		join quikserve.dbo.seven_shifts_locations l on s.qs_store_id=l.store_id
		where s.[api_key] is not null
		and l.active in (1)
		--and s.user_id in (1926968,2510397)
		order by s.[user_id]
    `

	UserLists := []UserList{}
	err := db.MyDB().Select(&UserLists, sql2)
	if err != nil {
		log.Println(err)
	}

	var UserID string

	for i := range UserLists {
		UserID = fmt.Sprint(*UserLists[i].UserID)
		key = fmt.Sprint(*UserLists[i].APIKey)
		ContactAPI(UserID, key)
	}

}

//ContactAPI is a function that contacts the weather API.
func ContactAPI(UserID string, key string) {

	start := time.Now().AddDate(0, 0, -0).Format("2006-01-02")
	log.Println("Getting SHIFTS for User:", UserID, start)

	url := fmt.Sprintf("https://api.7shifts.com/v1/shifts/?start=%s&user_id=%s", start, UserID)

	c := exec.Command("curl", "-u", key, url)
	c.Stdout = os.Stdout
	outfile, err1 := os.Create("./shifts.json")
	if err1 != nil {
		fmt.Println("Error:", err1)
	}

	c.Stdout = outfile
	//c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	content, err := ioutil.ReadFile("./shifts.json")
	if err != nil {
		fmt.Println(err)
	}

	text := string(content)

	PostIt(text, key)

}

//PostIt is a function for posting to SQL
func PostIt(text string, key string) {

	listposts := []ListPosts{}
	sql := `exec crm.dbo.import_seven_shifts_shifts $1`

	err := db.MyDB().Select(&listposts, sql, text)
	if err != nil {
		log.Println(err)
	}

}
