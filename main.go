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
	http.HandleFunc("/tickets", getTicketHandler)
	http.HandleFunc("/orgs", getOrgTicketHandler)
	http.HandleFunc("/user", getUserTicketsHandler)
	http.HandleFunc("/urgent", getUrgentTicketHandler)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func getTicketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	text := strings.Replace(r.FormValue("text"), "\r", "", -1)

	subject, orgID, userID, groupID, status := pkg.GetTicketSubject(text)
	orgName := pkg.GetOrganization(orgID)
	assignee := pkg.GetAssignee(userID)
	group := pkg.GetGroup(groupID)

	jsonResp, _ := json.Marshal(struct {
		Type string `json:"response_type"`
		Text string `json:"text"`
	}{
		Type: "in_channel",
		Text: fmt.Sprintf("*Subject* : %v \n *Organization* : %v \n *Assignee* : %v \n *Group* : %v \n *Status* : %v", subject, orgName, assignee, group, status),
	})

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}

func getOrgTicketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	text := strings.Replace(r.FormValue("text"), "\r", "", -1)

	orgTicket, orgPriority, orgStatus := pkg.GetOrgTickets(text)

	organizationTicketList := make([]string, len(orgTicket))

	for i := range organizationTicketList {
		organizationTicketList[i] = orgTicket[i] + " " + orgPriority[i] + " " + orgStatus[i] + " \n"
	}

	jsonResp, _ := json.Marshal(struct {
		Type string `json:"response_type"`
		Text string `json:"text"`
	}{
		Type: "in_channel",
		Text: fmt.Sprintf("%v", organizationTicketList),
	})

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}

func getUserTicketsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	text := strings.Replace(r.FormValue("text"), "\r", "", -1)

	userTicket, ticketPrority, ticketStatus := pkg.GetUserTickets(text)

	userTicketList := make([]string, len(userTicket))

	for i := range userTicketList {
		userTicketList[i] = userTicket[i] + " " + ticketPrority[i] + " " + ticketStatus[i] + " \n"
	}

	jsonResp, _ := json.Marshal(struct {
		Type string `json:"response_type"`
		Text string `json:"text"`
	}{
		Type: "in_channel",
		Text: fmt.Sprintf("%v", userTicketList),
	})

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}

func getUrgentTicketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	urgentTicket, urgentStatus, ugentAssignee := pkg.GetUrgentTickets()

	urgentTicketList := make([]string, len(urgentTicket))

	for i := range urgentTicketList {
		urgentTicketList[i] = urgentTicket[i] + " " + urgentStatus[i] + " " + ugentAssignee[i] + " \n"
	}

	jsonResp, _ := json.Marshal(struct {
		Type string `json:"response_type"`
		Text string `json:"text"`
	}{
		Type: "in_channel",
		Text: fmt.Sprintf("%v", urgentTicketList),
	})

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}
