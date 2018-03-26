package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic am9yZGFuLmthc3BlckBmb3JnZXJvY2suY29tL3Rva2VuOmtja0JTREx6YWs2V2NSWEZmQkt6eldCZjBNZ1pnWHJEQWFCbk1nRGc=")

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	myOrg := OrgGenerated{}

	jsonErr := json.Unmarshal(body, &myOrg)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	defer res.Body.Close()

	response := myOrg.Organization.Name

	return response
}
