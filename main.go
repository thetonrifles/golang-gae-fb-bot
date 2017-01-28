package api

import (
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

// initialize routing
func init() {
	r := mux.NewRouter()
	r.HandleFunc("/api/webhook", GetWebHookHandler).Methods("GET")
	r.HandleFunc("/api/webhook", PostWebHookHandler).Methods("POST")
	http.Handle("/", r)
}

// get handler for webhook validation
func GetWebHookHandler(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	mode := r.URL.Query()["hub.mode"][0]
	token := r.URL.Query()["hub.verify_token"][0]

	if mode == "subscribe" && token == FB_WEBHOOK_VERIFY_TOKEN {
		log.Debugf(context, "Validating webhook.")
		challenge := r.URL.Query()["hub.challenge"][0]
		fmt.Fprintf(w, challenge)
	} else {
		log.Errorf(context, "Failed validation. Make sure the validation tokens match.")
		http.Error(w, "Failed validation. Make sure the validation tokens match.", http.StatusForbidden)
	}
}

// post handler for replying to messages
func PostWebHookHandler(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	body, err := ioutil.ReadAll(r.Body);
	if err != nil {
		log.Errorf(context, "Error reading request body: %v", err)
		http.Error(w, "An error occurred. Try again.", http.StatusInternalServerError)
	}

	var data PostRequestData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Errorf(context, "Request decoding failed: %v", err)
		http.Error(w, "An error occurred. Try again.", http.StatusInternalServerError)
	}

	log.Debugf(context, "Request body: %v", data)

	sendMessage(r, data)
}

func sendMessage(r *http.Request, data PostRequestData) {
	context := appengine.NewContext(r)

	response := PostResponseData{}
	response.AccessToken = FB_PAGE_TOKEN
	response.Recipient.ID = data.Entry[0].Messaging[0].Sender.ID
	response.Message.Text = data.Entry[0].Messaging[0].Message.Text

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(response)

	client := urlfetch.Client(context)
	resp, err := client.Post("https://graph.facebook.com/v2.6/me/messages", "application/json", payload)

	if err != nil {
		log.Errorf(context, "%v", err)
	}

	log.Debugf(context, "%v", resp)
	defer resp.Body.Close()
}
