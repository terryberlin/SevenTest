package timepunches

import (
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	//"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

//MainTimePunches is a func
func MainTimePunches() {
	UserLists()
}

//UserLists is a function that returns a list of cities
func UserLists() {

	var status string
	var key string

	status = "reveal"

	listposts := []ListPosts{}
	sql1 := `exec crm.dbo.key_status $1`

	err1 := DB().Select(&listposts, sql1, status)
	if err1 != nil {
		log.Println(err1)
	}

	sql2 := `
		select user_id as UserID, api_key as APIKey
		from quikserve.dbo.seven_shifts_users s
		where api_key is not null
		--and s.user_id in (1926968,2510397)
		order by user_id
    `

	UserLists := []UserList{}
	err := DB().Select(&UserLists, sql2)
	if err != nil {
		log.Println(err)
	}

	var UserID string

	for i := range UserLists {
		UserID = fmt.Sprint(*UserLists[i].UserID)
		key = fmt.Sprint(*UserLists[i].APIKey)
		ContactAPI(UserID, key)
	}

	status = "disguise"

	listposts = []ListPosts{}
	sql3 := `exec crm.dbo.key_status $1`

	err3 := DB().Select(&listposts, sql3, status)
	if err3 != nil {
		log.Println(err3)
	}
}

//ContactAPI is a function that contacts the weather API.
func ContactAPI(UserID string, key string) {
	//log.Println("66")
	start := time.Now().AddDate(0, 0, -3).Format("2006-01-02")
	log.Println("Getting TIMEPUNCHES for User:", UserID, start)

	url := fmt.Sprintf("https://api.7shifts.com/v1/time_punches/?clocked_in=%s&user_id=%s", start, UserID)

	c := exec.Command("curl", "-u", key, url)
	c.Stdout = os.Stdout
	outfile, err1 := os.Create("./timepunches.json")
	if err1 != nil {
		fmt.Println("Error:", err1)
	}

	c.Stdout = outfile
	//c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	content, err := ioutil.ReadFile("./timepunches.json")
	if err != nil {
		fmt.Println(err)
	}

	text := string(content)

	PostIt(text, key)

}

//PostIt is a function for posting to SQL
func PostIt(text string, key string) {

	listposts := []ListPosts{}
	sql := `exec crm.dbo.import_seven_shifts_timepunches $1`

	err := DB().Select(&listposts, sql, text)
	if err != nil {
		log.Println(err)
	}

}

//DB : DB is a function that connects to SQL server.
func DB() *sqlx.DB {

	serv := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	database := os.Getenv("DB_DATABASE")

	db, err := sqlx.Connect("mssql", fmt.Sprintf(`server=%s;user id=%s;password=%s;database=%s;log1;encrypt=disable`, serv, user, pass, database))

	if err != nil {
		log.Println(err)
	}

	return db

}