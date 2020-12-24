package keys

import (
	"log"

	"github.com/quikserve/SevenTest/db"
)

type (
	//ListPosts is a structure of weather posts.
	ListPosts struct {
		Result *string `db:"Result"`
	}
)

//MainKeys is a func
func MainKeys(status string) {
	//log.Println(status)
	//toggle keys
	listposts := []ListPosts{}
	sql1 := `exec crm.dbo.key_status $1`

	err1 := db.MyDB().Select(&listposts, sql1, status)
	if err1 != nil {
		log.Println(err1)
	}
}
