package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type MyJsonName struct {
	Ticket struct {
		AssigneeID              int           `json:"assignee_id"`
		BrandID                 int           `json:"brand_id"`
		CollaboratorIds         []interface{} `json:"collaborator_ids"`
		CreatedAt               string        `json:"created_at"`
		Description             string        `json:"description"`
		DueAt                   interface{}   `json:"due_at"`
		EmailCcIds              []interface{} `json:"email_cc_ids"`
		ExternalID              interface{}   `json:"external_id"`
		FollowerIds             []interface{} `json:"follower_ids"`
		FollowupIds             []interface{} `json:"followup_ids"`
		ForumTopicID            interface{}   `json:"forum_topic_id"`
		GroupID                 int           `json:"group_id"`
		ID                      int           `json:"id"`
		OrganizationID          int           `json:"organization_id"`
		Priority                string        `json:"priority"`
		ProblemID               interface{}   `json:"problem_id"`
		RawSubject              string        `json:"raw_subject"`
		Recipient               interface{}   `json:"recipient"`
		RequesterID             int           `json:"requester_id"`
		SatisfactionProbability float64       `json:"satisfaction_probability"`
		SharingAgreementIds     []interface{} `json:"sharing_agreement_ids"`
		Status                  string        `json:"status"`
		Subject                 string        `json:"subject"`
		SubmitterID             int           `json:"submitter_id"`
		Tags                    []string      `json:"tags"`
		TicketFormID            int           `json:"ticket_form_id"`
		Type                    string        `json:"type"`
		UpdatedAt               string        `json:"updated_at"`
		URL                     string        `json:"url"`
	} `json:"ticket"`
}

func GetTicketSubject(text string) (string, int, int) {

	url := "https://forgerock.zendesk.com/api/v2/tickets/" + text + ".json"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic am9yZGFuLmthc3BlckBmb3JnZXJvY2suY29tL3Rva2VuOmtja0JTREx6YWs2V2NSWEZmQkt6eldCZjBNZ1pnWHJEQWFCbk1nRGc=")

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	myTicket := MyJsonName{}

	jsonErr := json.Unmarshal(body, &myTicket)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	defer res.Body.Close()

	ticketSubject := myTicket.Ticket.Subject
	orgID := myTicket.Ticket.OrganizationID
	userID := myTicket.Ticket.AssigneeID

	return ticketSubject, orgID, userID
}
