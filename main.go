package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		resp := map[string]any{
			"slackUsername": "o.j. lewis",
			"backend": true,
			"age": 22,
			"bio": "tech ethusiast",
		}

		json.NewEncoder(res).Encode(resp)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}