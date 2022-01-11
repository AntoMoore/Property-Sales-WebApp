package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"example.com/resources"
)

// reference to index.html
var tmpls = template.Must(template.ParseFiles("../templates/header.html", 
	"../templates/footer.html", "../templates/header.html", "../templates/main.html", 
	"../templates/about.html", "../templates/agents.html", "../templates/postAgent.html"))

type Page struct {
	Title string
	AgentResults    *resources.AgentResults
}

// template-handler helper function
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpls.ExecuteTemplate(w, tmpl, data)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "main", &Page{Title: "Home"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "about", &Page{Title: "About"})
}

func getAgentsHandler(w http.ResponseWriter, r *http.Request) {
	// connection to API
	httpClient := &http.Client{Timeout: 10 * time.Second}
	resourceApi := resources.NewClient(httpClient)
	searchQuery := ""

	// check results for errors
	results, err := resourceApi.GetAgents(searchQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// print results
	fmt.Printf("%+v", results)

	//tmpls.ExecuteTemplate(w, "agents", agentSearch)
	display(w, "agents", &Page{Title: "Agent", AgentResults: results})
}

func postAgentHandler(w http.ResponseWriter, r *http.Request){
	// connection to API
	httpClient := &http.Client{Timeout: 10 * time.Second}
	resourceApi := resources.NewClient(httpClient)

	// check for post method
	if r.Method != http.MethodPost {
		// tmpls.Execute(w, nil)
		display(w, "postAgent", nil)
		return
	}

	// map key-value pairs from the form data
	data := url.Values{}
	data.Set("name", r.FormValue("agentName"))
	data.Set("commission", r.FormValue("agentCommission"))

	// call resources
	err := resourceApi.PostAgent(data)

	// error handling
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// display agents
	getAgentsHandler(w,r)
}

// NB - handle search paramaters
/*
	// check URL for errors		
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get Query params
	params := u.Query()
	searchQuery := params.Get("q")
*/

func main() {
	// Server PORT
	port := "3000"

	// static styling sheets
	fs := http.FileServer(http.Dir("../assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// URL routing
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/agents", getAgentsHandler)
	mux.HandleFunc(("/postAgent"), postAgentHandler)

	//Listen on port 3000
	http.ListenAndServe(":"+port, mux)
}
