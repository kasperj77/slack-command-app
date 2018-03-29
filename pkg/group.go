package pkg

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type GroupGenerated struct {
	Group struct {
		URL       string    `json:"url"`
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Deleted   bool      `json:"deleted"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"group"`
}

func GetGroup(id int) string {

	url := "https://forgerock.zendesk.com/api/v2/groups/" + strconv.Itoa(id) + ".json"

	body := Init(url)

	myGroup := GroupGenerated{}

	jsonErr := json.Unmarshal(body, &myGroup)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	response := myGroup.Group.Name

	return response
}
