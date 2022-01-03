package main

import (
	"bytes"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"example.com/resources"
)

// reference to index.html
var tpl = template.Must(template.ParseFiles("../template/index.html"))

type AgentSearch struct {
	Query      		string
	AgentResults    *resources.AgentResults
}

// template execution handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

// handle search paramaters
func searchHandler(resourceApi *resources.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check URL for errors
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// get Query params
		params := u.Query()
		searchQuery := params.Get("q")

		// check results for errors
		results, err := resourceApi.GetAgents(searchQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// agent search params
		agentSearch := &AgentSearch{
			Query:      	searchQuery,
			AgentResults:   results,
		}

		// parse AgentSearch
		buf := &bytes.Buffer{}
		err = tpl.Execute(buf, agentSearch)

		// check for IO errors
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	}
}

func main() {
	// Server PORT
	port := "3000"

	// Resource Client
	myClient := &http.Client{Timeout: 10 * time.Second}
	resourceApi := resources.NewClient(myClient)

	// static styling sheets
	fs := http.FileServer(http.Dir("../assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// URL routing
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/search", searchHandler(resourceApi))
	http.ListenAndServe(":"+port, mux)
}
