package resources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"log"

)

type Client struct {
	http *http.Client
}

type AgentResults[] struct {
	AgentID         int     `json:"agentId"`
	AgentName       string  `json:"agentName"`
	AgentCommission float64 `json:"agentCommission"`
}

// connection
func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient}
}

// get agents from API 
func (c *Client) GetAgents(query string) (*AgentResults, error) {
	var endpoint = "http://localhost:4567/openproperty/agents/"

	/*
	if query == "" {
		// no paramaters (get all agents)
		endpoint = "http://localhost:4567/openproperty/agents/"
	} else {
		// get Agent by given id
		endpoint = fmt.Sprintf("http://localhost:4567/openproperty/agents/?id=%s", url.QueryEscape(query))
	}
	*/

	// response errors
	resp, err := c.http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	// close body and check for errors
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// error if status not 200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	// parse json and return
	res := &AgentResults{}
	return res, json.Unmarshal(body, res)
}


func (c *Client) PostAgent(details url.Values) (error) {
	var endpoint = "http://localhost:4567/openproperty/agents"
	
	// response error
	resp, err := c.http.PostForm(endpoint, details)
	if err != nil {
		log.Fatal(err)
	}

	// close body and check for errors
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print results
	fmt.Printf("%+v", body)

	return nil
}
