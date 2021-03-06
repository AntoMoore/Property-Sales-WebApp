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

type PropertyResults []struct {
	PropertyID      int     `json:"propertyId"`
	PropertyType    string  `json:"propertyType"`
	PropertyAddress string  `json:"propertyAddress"`
	PropertyValue   float64 `json:"propertyValue"`
	PropertyAgent   struct {
		AgentID         int     `json:"agentId"`
		AgentName       string  `json:"agentName"`
		AgentCommission float64 `json:"agentCommission"`
	} `json:"propertyAgent"`
}

type SaleResults []struct {
	SaleID       int    `json:"saleId"`
	SaleDate     string `json:"saleDate"`
	SaleProperty struct {
		PropertyID      int     `json:"propertyId"`
		PropertyType    string  `json:"propertyType"`
		PropertyAddress string  `json:"propertyAddress"`
		PropertyValue   float64 `json:"propertyValue"`
		PropertyAgent   struct {
			AgentID         int     `json:"agentId"`
			AgentName       string  `json:"agentName"`
			AgentCommission float64 `json:"agentCommission"`
		} `json:"propertyAgent"`
	} `json:"saleProperty"`
}

// connection
func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient}
}

// get agents from API 
func (c *Client) GetAgents(query string) (*AgentResults, error) {
	var endpoint string

	if query == "" {
		// no paramaters (get all agents)
		endpoint = "http://localhost:4567/openproperty/agents/"
	} else {
		// get Agent by given id
		endpoint = fmt.Sprintf("http://localhost:4567/openproperty/agents/?id=%s", url.QueryEscape(query))
	}
	

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
		return err
	}

	// print results
	fmt.Printf("%+v", body)

	return nil
}

func (c *Client) DeleteAgent(data string) (error) {
	var endpoint = "http://localhost:4567/openproperty/agents/remove/?id=" + data
	
	// create request
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// fetch request
	resp,err := c.http.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}



	// close body and check for errors
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// print results
	fmt.Printf("%+v", body)

	return nil

}

// get agents from API 
func (c *Client) GetProperties(query string) (*PropertyResults, error) {
	var endpoint string

	if query == "" {
		// no paramaters (get all)
		endpoint = "http://localhost:4567/openproperty/properties/"
	} else {
		// get by given id
		endpoint = fmt.Sprintf("http://localhost:4567/openproperty/properties/?id=%s", url.QueryEscape(query))
	}

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
	res := &PropertyResults{}
	return res, json.Unmarshal(body, res)
}

func (c *Client) PostProperty(details url.Values) (error) {
	var endpoint = "http://localhost:4567/openproperty/properties"
	
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

func (c *Client) GetSales(query string) (*SaleResults, error) {
	var endpoint string

	if query == "" {
		// no paramaters (get all)
		endpoint = "http://localhost:4567/openproperty/sales/"
	} else {
		// get by given id
		endpoint = fmt.Sprintf("http://localhost:4567/openproperty/sales/?id=%s", url.QueryEscape(query))
	}

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
	res := &SaleResults{}
	return res, json.Unmarshal(body, res)
}

func (c *Client) PostSale(details url.Values) (error) {
	var endpoint = "http://localhost:4567/openproperty/sales"
	
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
