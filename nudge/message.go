// Package nudge contains utility functions for building Slack messages from nudge notifications
package nudge

import (
	"encoding/json"
	"fmt"
)

// IssuesDevstreamActivityPayload
type IssuesDevstreamActivityPayload struct {
	Issue struct {
		Key string `json:"key"`
	}
	Activity string `json:"activity"`
	Feedback struct {
		Text     string     `json:"text"`
		Timeline [][]string `json:"timeline"`
	}
}

// IssuesDevstreamActivity ...
func IssuesDevstreamActivity(body []byte) string {
	var m = IssuesDevstreamActivityPayload{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println("error:", err)
		return "error"
	}
	issueID := m.Issue.Key
	feedback := m.Feedback
	message := issueID + " possibly has " + feedback.Text
	return message
}
