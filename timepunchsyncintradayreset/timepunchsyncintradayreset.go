package timepunchsyncintradayreset

import (
	"log"
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

//MainTimePunchSyncIntradayReset is a func
func MainTimePunchSyncIntradayReset() {
	log.Println("match intraday with new eod and update hour_id in _punch_map table")
	log.Println("for remaining records spin through and delete from 7Shifts")
	log.Println("then delete from _punch_map table")
	log.Println("rerun eod sync")
}
