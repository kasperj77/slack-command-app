package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic am9yZGFuLmthc3BlckBmb3JnZXJvY2suY29tL3Rva2VuOmtja0JTREx6YWs2V2NSWEZmQkt6eldCZjBNZ1pnWHJEQWFCbk1nRGc=")

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	myGroup := GroupGenerated{}

	jsonErr := json.Unmarshal(body, &myGroup)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	defer res.Body.Close()

	response := myGroup.Group.Name

	return response
}
