package main

import (
	"net/http"
	"encoding/json"
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

	http.ListenAndServe(":8080", r)
}