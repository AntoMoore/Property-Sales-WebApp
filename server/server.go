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
	"../templates/properties.html", "../templates/createProperty.html", 
	"../templates/sales.html", "../templates/createSale.html"))

type Page struct {
	Title string
	AgentResults    *resources.AgentResults
	PropertyResults *resources.PropertyResults
	SaleResults 	*resources.SaleResults
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

func deleteAgentHandler(w http.ResponseWriter, r *http.Request){
	// map key-value pairs from the form data
	//data := url.Values{}
	//data.Set("id", r.FormValue("agentID"))

	// connection to API
	httpClient := &http.Client{Timeout: 10 * time.Second}
	resourceApi := resources.NewClient(httpClient)

	data := r.FormValue("agentID")
	
	fmt.Println("Delete Agent: ", r.FormValue("agentID"))

	// call resources
	err := resourceApi.DeleteAgent(data)

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

func postPropertyHandler(w http.ResponseWriter, r *http.Request){
	// connection to API
	httpClient := &http.Client{Timeout: 10 * time.Second}
	resourceApi := resources.NewClient(httpClient)

	// check for post method
	if r.Method != http.MethodPost {
		// tmpls.Execute(w, nil)
		display(w, "createProperty", nil)
		return
	}

	// map key-value pairs from the form data
	data := url.Values{}
	data.Set("type", r.FormValue("propertyType"))
	data.Set("address", r.FormValue("propertyAddress"))
	data.Set("value", r.FormValue("propertyValue"))
	data.Set("agentId", r.FormValue("propertyAgentID"))


	// call resources
	err := resourceApi.PostProperty(data)

	// error handling
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// display properties
	getPropertiesHandler(w,r)
}

func getSalesHandler(w http.ResponseWriter, r *http.Request) {
	// connection to API
	httpClient := &http.Client{Timeout: 10 * time.Second}
	resourceApi := resources.NewClient(httpClient)

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
	results, err := resourceApi.GetSales(searchQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// print results
	fmt.Printf("%+v", results)

	//tmpls.ExecuteTemplate(w, "agents", agentSearch)
	display(w, "sales", &Page{Title: "Sales", SaleResults: results})
}

func postSaleHandler(w http.ResponseWriter, r *http.Request){
	// connection to API
	httpClient := &http.Client{Timeout: 10 * time.Second}
	resourceApi := resources.NewClient(httpClient)

	// check for post method
	if r.Method != http.MethodPost {
		// tmpls.Execute(w, nil)
		display(w, "createSale", nil)
		return
	}

	// map key-value pairs from the form data
	data := url.Values{}
	data.Set("propertyId", r.FormValue("propertyID"))

	// call resources
	err := resourceApi.PostSale(data)

	// error handling
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// display sales
	getSalesHandler(w,r)
}

// NB - handle search paramaters
// Needed to implement search functionality
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
	mux.HandleFunc(("/deleteAgent"), deleteAgentHandler)
	mux.HandleFunc(("/properties"), getPropertiesHandler)
	mux.HandleFunc(("/createProperty"), postPropertyHandler)
	mux.HandleFunc(("/sales"), getSalesHandler)
	mux.HandleFunc(("/createSale"), postSaleHandler)

	//Listen on port 3000
	http.ListenAndServe(":"+port, mux)
}
