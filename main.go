package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		resp := struct{
			SlackUsername string `json:"slackUsername"`
			Backend bool `json:"backend"`
			Age int `json:"age"`
			Bio string `json:"bio"`
		}{
			SlackUsername: "o.j. lewis",
			Backend: true,
			Age: 22,
			Bio: "tech ethusiast",
		}

		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "Get")
		json.NewEncoder(res).Encode(resp)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}