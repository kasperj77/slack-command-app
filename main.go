package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slack-command-app/pkg"
	"strings"
)

func main() {
	http.HandleFunc("/", getTicketHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func getTicketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	text := strings.Replace(r.FormValue("text"), "\r", "", -1)

	subject, orgID, userID := pkg.GetTicketSubject(text)
	orgName := pkg.GetOrganization(orgID)
	assignee := pkg.GetAssignee(userID)

	jsonResp, _ := json.Marshal(struct {
		Type string `json:"response_type"`
		Text string `json:"text"`
	}{
		Type: "in_channel",
		Text: fmt.Sprintf(subject + orgName + assignee),
	})

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}
