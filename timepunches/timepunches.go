package timepunches

import (
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	//"database/sql"
	"fmt"
	"log"

	// _ "github.com/denisenkom/go-mssqldb"
	// _ "github.com/go-sql-driver/mysql"
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

//MainTimePunches is a func
func MainTimePunches() {

	var key string

	/*sql2 := `
		select [user_id] as UserID, s.api_key as APIKey
		from quikserve.dbo.seven_shifts_users s
		join quikserve.dbo.seven_shifts_locations l on l.store_id=s.qs_store_id and l.api_key=s.api_key
		where s.api_key is not null and l.active in (1)
		--and s.[user_id]=2513633
		order by [user_id]
	`*/

	sql2 := `
		select distinct u.[user_id] as UserID, u.[api_key] as APIKey --, r.empl_id, r.emp_name
		from QsSupport.dbo.RDMTimePunches r
		join QuikServe.dbo.seven_shifts_locations l on r.StoreId=l.store_id
		join QuikServe.dbo.seven_shifts_users u on u.qs_emplID=r.empl_id
		where l.active=1
		union
		select distinct u.[user_id] as UserID, u.[api_key] as APIKey --, e.emplid, e.emp_name
		from quikserve.dbo.employee_hours e join quikserve.dbo.stores s on e.Store_No=s.STORE_NO
		join QuikServe.dbo.seven_shifts_locations l on l.store_id=s.id
		join QuikServe.dbo.seven_shifts_users u on u.qs_emplID=e.emplid
		where l.active=1
		and e.Bus_Date=convert(date,dateadd(d,-1,getdate()))
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

	start := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
	//start := time.Now().AddDate(0, 0, -3).Format("2006-01-02")
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

	err := db.MyDB().Select(&listposts, sql, text)
	if err != nil {
		log.Println(err)
	}

}

// //DB : DB is a function that connects to SQL server.
// func DB() *sqlx.DB {

// 	serv := os.Getenv("DB_SERVER")
// 	user := os.Getenv("DB_USER")
// 	pass := os.Getenv("DB_PASS")
// 	database := os.Getenv("DB_DATABASE")

// 	db, err := sqlx.Connect("mssql", fmt.Sprintf(`server=%s;user id=%s;password=%s;database=%s;log1;encrypt=disable`, serv, user, pass, database))

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return db

// }
