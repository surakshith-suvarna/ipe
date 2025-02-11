package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/pusher/pusher-http-go/v5"
)

var client pusher.Client

func init() {
	client = pusher.Client{
		AppID:  "1",
		Key:    "278d525bdf162c739803",
		Secret: "7ad3753142a6693b25b9",
		Host:   ":8080",
	}
}

func pusherPresenceAuth(res http.ResponseWriter, req *http.Request) {
	log.Println("Presence Request")
	presenceData := pusher.MemberData{
		UserID:   "1",
		UserInfo: map[string]string{},
	}

	params, _ := io.ReadAll(req.Body)
	response, err := client.AuthorizePresenceChannel(params, presenceData)

	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprint(res, string(response))
}

func pusherPrivateAuth(res http.ResponseWriter, req *http.Request) {
	params, _ := io.ReadAll(req.Body)
	response, err := client.AuthorizePrivateChannel(params)

	log.Printf("Private Request %s", params)
	log.Printf("Auth %s", response)

	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprint(res, string(response))
}

func triggerMessage(res http.ResponseWriter, _ *http.Request) {
	err := client.Trigger("private-messages", "messages", "The message from server")
	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprint(res, "OK")
}

func hookcallback(res http.ResponseWriter, r *http.Request) {
	bytes, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	event := struct {
		Events []struct {
			Name string `json:"name"`
		} `json:"events"`
	}{}

	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		panic(err)
	}

	err = client.Trigger("private-webhook", event.Events[0].Name, "The Webhoook from server")
	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprint(res, "OK")
}

func main() {
	http.HandleFunc("/pusher/presence/auth", pusherPresenceAuth)
	http.HandleFunc("/pusher/private/auth", pusherPrivateAuth)
	http.HandleFunc("/trigger", triggerMessage)
	http.HandleFunc("/hook", hookcallback)
	http.Handle("/", http.FileServer(http.Dir("./")))
	_ = http.ListenAndServe(":5000", nil)
}
