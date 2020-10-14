package main

import (
	"log"
	"time"

	"github.com/smilextay/az_command_go/src/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd Execute err:%v", err)
	}
}

type RUNOOBTBL struct {
	// runoob_author
	RUNOOBAUTHOR string `json:runoob_author`
	// runoob_id
	RUNOOBID int32 `json:runoob_id`
	// runoob_title
	RUNOOBTITLE string `json:runoob_title`
	// submission_date
	SUBMISSIONDATE time.Time `json:submission_date`
}

func (model RUNOOBTBL) TableName() string {
	return "runoob_tbl"
}
