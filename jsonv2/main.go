package main

import (
	"encoding/json/v2"
	"log"
	"time"
)

type TodoEntry struct {
	UserId   int             `json:"user_id,string"`
	Priority int64           `json:"priority,omitzero"`
	Created  time.Time       `json:"created,format:RFC850"`
	Details  TaskInformation `json:",inline"`
}

type TaskInformation struct {
	Deadline  time.Time `json:"deadline,format:RFC850"`
	CreatedBy int       `json:"created_by"`
}

func main() {
	te := TodoEntry{
		Created:  time.Now(),
		UserId:   100,
		Priority: 0,
		Details: TaskInformation{
			CreatedBy: 10,
			Deadline:  time.Now().Add(24 * time.Hour),
		},
	}

	b, err := json.Marshal(te)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(b))
	log.Println("parsing again...")

	nte := TodoEntry{}

	err = json.Unmarshal(b, &nte)

	log.Printf("%+v\n", nte)
}
