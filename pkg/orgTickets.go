package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type MyJsonData struct {
	Tickets []struct {
		URL                     string        `json:"url"`
		ID                      int           `json:"id"`
		ExternalID              interface{}   `json:"external_id"`
		CreatedAt               time.Time     `json:"created_at"`
		UpdatedAt               time.Time     `json:"updated_at"`
		Type                    string        `json:"type"`
		Subject                 string        `json:"subject"`
		RawSubject              string        `json:"raw_subject"`
		Description             string        `json:"description"`
		Priority                string        `json:"priority"`
		Status                  string        `json:"status"`
		Recipient               interface{}   `json:"recipient"`
		RequesterID             int           `json:"requester_id"`
		SubmitterID             int           `json:"submitter_id"`
		AssigneeID              int           `json:"assignee_id"`
		OrganizationID          int           `json:"organization_id"`
		GroupID                 int           `json:"group_id"`
		CollaboratorIds         []interface{} `json:"collaborator_ids"`
		FollowerIds             []interface{} `json:"follower_ids"`
		EmailCcIds              []interface{} `json:"email_cc_ids"`
		ForumTopicID            interface{}   `json:"forum_topic_id"`
		ProblemID               interface{}   `json:"problem_id"`
		HasIncidents            bool          `json:"has_incidents"`
		IsPublic                bool          `json:"is_public"`
		DueAt                   interface{}   `json:"due_at"`
		Tags                    []string      `json:"tags"`
		FollowupIds             []interface{} `json:"followup_ids"`
		TicketFormID            int           `json:"ticket_form_id"`
		BrandID                 int           `json:"brand_id"`
		SatisfactionProbability float64       `json:"satisfaction_probability"`
		AllowChannelback        bool          `json:"allow_channelback"`
		ResultType              string        `json:"result_type"`
	} `json:"results"`
}

func GetOrgTickets() int {

	url := "https://forgerock.zendesk.com/api/v2/search.json?query=type%3Aticket%20organization%3A%22Kansas%20City%20Plant%20(NNSA)%22%20status%3Aopen%20status%3Apending"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic am9yZGFuLmthc3BlckBmb3JnZXJvY2suY29tL3Rva2VuOmtja0JTREx6YWs2V2NSWEZmQkt6eldCZjBNZ1pnWHJEQWFCbk1nRGc=")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	myTicket := MyJsonData{}

	jsonErr := json.Unmarshal(body, &myTicket)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	defer res.Body.Close()

	//jsonErr := json.Unmarshal(body, &myTicket)

	return myTicket.Tickets[0].AssigneeID

}
