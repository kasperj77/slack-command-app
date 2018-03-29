package pkg

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
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

func GetOrgTickets(org string) ([]string, []string, []string) {

	url := "https://forgerock.zendesk.com/api/v2/search.json?query=" + url.QueryEscape("type:ticket organization:\""+org+"\" status:open status:pending")

	body := Init(url)

	myTicket := MyJsonData{}

	jsonErr := json.Unmarshal(body, &myTicket)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// getting the numbers for org
	orgTickets := make([]string, len(myTicket.Tickets))

	for i := range orgTickets {
		orgTickets[i] = strconv.Itoa(myTicket.Tickets[i].ID)
	}

	// getting the tickets status
	ticketPriority := make([]string, len(myTicket.Tickets))

	for i := range ticketPriority {
		ticketPriority[i] = myTicket.Tickets[i].Priority
	}

	// getting the tickets status
	ticketStatus := make([]string, len(myTicket.Tickets))

	for i := range ticketStatus {
		ticketStatus[i] = myTicket.Tickets[i].Status
	}

	//jsonErr := json.Unmarshal(body, &myTicket)
	return orgTickets, ticketPriority, ticketStatus

}
