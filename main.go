package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"journeysoftware/jaya/github"
	"journeysoftware/jaya/nudge"
	"log"
	"net/http"
	"os"
	"strings"
)

const secret = "itsthejourney"

type HookDelivery struct {
	Signature string
	Event     string
	ID        string
	Body      []byte
}

func signBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}

func verifySignature(secret []byte, signature string, body []byte) bool {

	const signaturePrefix = "sha1="
	const signatureLength = 45 // len(SignaturePrefix) + len(hex(sha1))

	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))

	return hmac.Equal(signBody(secret, body), actual)
}

func ParseHook(secret []byte, req *http.Request) (*HookDelivery, error) {

	payload := HookDelivery{}
	payload.Signature = req.Header.Get("X-Hub-Signature")
	if len(payload.Signature) == 0 {
		return nil, errors.New("no signature")
	}
	payload.Event = req.Header.Get("X-GitHub-Event")
	if len(payload.Event) == 0 {
		return nil, errors.New("no event")
	}

	payload.ID = req.Header.Get("X-GitHub-Delivery")
	if len(payload.ID) == 0 {
		return nil, errors.New("no delivery ID")
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	if !verifySignature(secret, payload.Signature, body) {
		return nil, errors.New("invalid signature")
	}

	payload.Body = body

	return &payload, nil
}

type NudgeDelivery struct {
	ActivityType string
	Body         []byte
}

func ParseNudge(req *http.Request) (*NudgeDelivery, error) {
	payload := NudgeDelivery{}
	payload.ActivityType = req.Header.Get("Activity-Type")
	if len(payload.ActivityType) == 0 {
		return nil, errors.New("missing activity type")
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	payload.Body = body

	return &payload, nil
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	hc, err := ParseHook([]byte(secret), r)
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed processing hook! ('%s')", err)
		io.WriteString(w, "{}")
		return
	}
	message := github.Issues(hc.Body)
	postToSlack(message)
	w.WriteHeader(http.StatusOK)
	return
}

func nudgeHandler(w http.ResponseWriter, r *http.Request) {
	hc, err := ParseNudge(r)
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed processing nudge! ('%s')", err)
		io.WriteString(w, "{}")
		return
	}
	f := map[string]func(body []byte) string{
		"IssuesDevstreamActivity": nudge.IssuesDevstreamActivity}

	activityType := r.Header.Get("Activity-Type")
	message := f[activityType](hc.Body)
	postToSlack(message)
	w.WriteHeader(http.StatusOK)
}

func postToSlack(message string) {
	url := "https://hooks.slack.com/services/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"
	m := make(map[string]string)
	m["text"] = message
	body, marshalErr := json.Marshal(m)
	if marshalErr != nil {
		fmt.Printf("Error: ", marshalErr)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Request error", err)
		return
	}
	fmt.Printf("post to slack success", resp)
	resp.Body.Close()
	return
}

func main() {
	http.HandleFunc("/hooks", hookHandler)
	http.HandleFunc("/nudges", nudgeHandler)
	port := os.Getenv("PORT")
	// port := "8080"
	if port == "" {
		log.Printf("$PORT not set")
		return
	}
	log.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
