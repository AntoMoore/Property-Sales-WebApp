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
	"../templates/footer.html", "../templates/main.html", 
	"../templates/agents.html", "../templates/createAgent.html", 
	"../templates/properties.html"))

type Page struct {
	Title string
	AgentResults    *resources.AgentResults
	PropertyResults *resources.PropertyResults
}

// template-handler helper function
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpls.ExecuteTemplate(w, tmpl, data)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "main", &Page{Title: "Home"})
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
	display(w, "agents", &Page{Title: "Agents", AgentResults: results})
}

func postAgentHandler(w http.ResponseWriter, r *http.Request){
	// connection to API
	httpClient := &http.Client{Timeout: 10 * time.Second}
	resourceApi := resources.NewClient(httpClient)

	// check for post method
	if r.Method != http.MethodPost {
		// tmpls.Execute(w, nil)
		display(w, "createAgent", nil)
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

func getPropertiesHandler(w http.ResponseWriter, r *http.Request) {
	// connection to API
	httpClient := &http.Client{Timeout: 10 * time.Second}
	resourceApi := resources.NewClient(httpClient)
	searchQuery := ""

	// check results for errors
	results, err := resourceApi.GetProperties(searchQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// print results
	fmt.Printf("%+v", results)

	//tmpls.ExecuteTemplate(w, "agents", agentSearch)
	display(w, "properties", &Page{Title: "Properties", PropertyResults: results})
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
	mux.HandleFunc("/agents", getAgentsHandler)
	mux.HandleFunc(("/createAgent"), postAgentHandler)
	mux.HandleFunc(("/properties"), getPropertiesHandler)

	//Listen on port 3000
	http.ListenAndServe(":"+port, mux)
}
