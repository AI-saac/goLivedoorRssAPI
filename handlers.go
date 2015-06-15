package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewsIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	url := fmt.Sprintf("http://news.livedoor.com/topics/rss/%s.xml", category)
	news := Items{}.get(url)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(news); err != nil {
		panic(err)
	}
}
