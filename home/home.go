package main

import (
	//"bytes"
	"html/template"
	"net/http"
	//"net/url"
	//"time"

	"example.com/resources"
)

// reference to index.html
var templates = template.Must(template.ParseFiles("../templates/header.html", "../templates/footer.html", "../templates/header.html", "../templates/main.html", "../templates/about.html"))

type Page struct {
	Title string
}

type AgentSearch struct {
	Query      		string
	AgentResults    *resources.AgentResults
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

//The handlers.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "main", &Page{Title: "Home"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "about", &Page{Title: "About"})
}


// template execution handler
/*
func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := templates.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}
*/

// handle search paramaters
/*
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
		err = templates.Execute(buf, agentSearch)

		// check for IO errors
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	}
}
*/

func main() {
	// Server PORT
	port := "3000"

	// Resource Client
	//myClient := &http.Client{Timeout: 10 * time.Second}
	//resourceApi := resources.NewClient(myClient)

	// static styling sheets
	fs := http.FileServer(http.Dir("../assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//mux.HandleFunc("/", indexHandler)
	//mux.HandleFunc("/search", searchHandler(resourceApi))
	//http.ListenAndServe(":"+port, mux)

	// URL routing
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/about", aboutHandler)

	//Listen on port 3000
	http.ListenAndServe(":"+port, mux)


	
}
