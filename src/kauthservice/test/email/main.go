package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type request struct {
	Sender         string `json:"sender,omitempty"`
	Recipient      string `json:"recipient,omitempty"`
	SenderEmail    string `json:"senderEmail,omitempty"`
	RecipientEmail string `json:"recipientEmail,omitempty"`
	Subject        string `json:"subject,omitempty"`
	Content        string `json:"content,omitempty"`
	IdempotencyKey string `json:"idempotencyKey,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var req request

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.RecipientEmail == "failemail@test.com" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	panic(http.ListenAndServe(":8081", http.HandlerFunc(handler)))
}
