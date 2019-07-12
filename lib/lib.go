package lib

import (
	"bytes"
	"net/http"

	"github.com/golang/glog"
)

// GetUserFromSlack gets users from Slack
func GetUserFromSlack(token string, id string) {
	urlAddress := "https://api.slack.com/scim/v1/Users/" + id
	req, err := http.NewRequest("GET", urlAddress, nil)
	if err != nil {
		glog.Fatal(err)
	}

	req.Header.Add("authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Fatal(err)
	}
	if resp.StatusCode >= 300 {
		glog.Fatalf("Failed Slack CALL - GET /User/[%s] - status[%s]", id, resp.Status)
	} else {
		glog.V(1).Infof("Slack CALL - GET /User/[%s] - status[%s]", id, resp.Status)
	}
}

// ActivateUserInSlack activates user in Slack
func ActivateUserInSlack(token string, active bool, id string) {
	urlAddress := "https://api.slack.com/scim/v1/Users/" + id

	var jsonStr = []byte(`{"active":false}`)
	if active {
		jsonStr = []byte(`{"active":true}`)
	}
	req, err := http.NewRequest("PATCH", urlAddress, bytes.NewBuffer(jsonStr))
	if err != nil {
		glog.Fatal(err)
	}

	req.Header.Add("authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Fatal(err)
	}
	if resp.StatusCode >= 300 {
		glog.Fatalf("Failed Slack CALL - PATCH /User/[%s] - status[%s]", id, resp.Status)
	} else {
		glog.V(1).Infof("Slack CALL - PATCH /User/[%s] - status[%s]", id, resp.Status)
	}
}

// RemoveTitlesAndPhoneInSlack removes titles and phone numbers in Slack
func RemoveTitlesAndPhoneInSlack(token string, id string) {
	urlAddress := "https://api.slack.com/scim/v1/Users/" + id

	var jsonStr = []byte(`{"title": "", "phoneNumbers": [ { "value": null, "type": "mobile" } ], "urn:scim:schemas:extension:enterprise:1.0": { "department": "" }}`)
	req, err := http.NewRequest("PATCH", urlAddress, bytes.NewBuffer(jsonStr))
	if err != nil {
		glog.Fatal(err)
	}

	req.Header.Add("authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Fatal(err)
	}
	if resp.StatusCode >= 300 {
		glog.Fatalf("Failed Slack CALL - PATCH /User/[%s] - status[%s]", id, resp.Status)
	} else {
		glog.V(1).Infof("Slack CALL - PATCH /User/[%s] - status[%s]", id, resp.Status)
	}
}
