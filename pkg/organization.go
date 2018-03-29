package pkg

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type OrgGenerated struct {
	Organization struct {
		URL                string        `json:"url"`
		ID                 int           `json:"id"`
		Name               string        `json:"name"`
		SharedTickets      bool          `json:"shared_tickets"`
		SharedComments     bool          `json:"shared_comments"`
		ExternalID         interface{}   `json:"external_id"`
		CreatedAt          time.Time     `json:"created_at"`
		UpdatedAt          time.Time     `json:"updated_at"`
		DomainNames        []interface{} `json:"domain_names"`
		Details            string        `json:"details"`
		Notes              string        `json:"notes"`
		GroupID            interface{}   `json:"group_id"`
		Tags               []string      `json:"tags"`
		OrganizationFields struct {
			AmerCsm            interface{} `json:"amer_csm"`
			ApacCsm            interface{} `json:"apac_csm"`
			Carr               float64     `json:"carr"`
			EmeaCsm            interface{} `json:"emea_csm"`
			ExtSvc             interface{} `json:"ext_svc"`
			OrganisationCcList interface{} `json:"organisation_cc_list"`
		} `json:"organization_fields"`
	} `json:"organization"`
}

func GetOrganization(id int) string {

	url := "https://forgerock.zendesk.com/api/v2/organizations/" + strconv.Itoa(id) + ".json"

	body := Init(url)

	myOrg := OrgGenerated{}

	jsonErr := json.Unmarshal(body, &myOrg)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	response := myOrg.Organization.Name

	return response
}
