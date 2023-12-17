package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// SearchResponse represents the structure of the search API response.
type SearchList []struct {
	ID          string    `json:"ID"`
	Description string    `json:"Description"`
	Owner       string    `json:"Owner"`
	Status      string    `json:"Status"`
	StartDate   time.Time `json:"StartDate"`
	EndDate     time.Time `json:"EndDate"`
	Priority    int       `json:"Priority"`
}

// Search handles chaincode search requests.
func (setup *OrgSetup) Search(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Search request")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeJSONError(w, fmt.Sprintf("Error reading request body: %s", err))
		return
	}

	// Define a struct to unmarshal the JSON request body
	var request struct {
		ChainCodeID string `json:"chaincodeid"`
		ChannelID   string `json:"channelid"`
		Status      string `json:"status"`
	}

	// Unmarshal the JSON into the request struct
	err = json.Unmarshal(body, &request)
	if err != nil {
		writeJSONError(w, fmt.Sprintf("Error unmarshalling JSON: %s", err))
		return
	}

	// Extract values from the request struct
	status := request.Status
	// For simplicity, assuming you have a single OrgSetup instance.
	network := setup.Gateway.GetNetwork(request.ChannelID) // Update with your channel name
	contract := network.GetContract(request.ChainCodeID)   // Update with your chaincode name
	results, err := contract.EvaluateTransaction("SearchTodoItems", status)
	if err != nil {
		writeJSONError(w, fmt.Sprintf("Error creating txn proposal: %s", err))
		return
	}

	var searchResponse SearchList
	// Replace YourStructType with the actual type of your response
	if err := json.Unmarshal(results, &searchResponse); err != nil {
		response := APIResponse{
			Status:  "error",
			Message: fmt.Sprintf("Error unmarshalling search response: %s", err),
			Data:    searchResponse,
		}
		writeJSONResponse(w, response)
		return
	}

	response := APIResponse{
		Status:  "success",
		Message: "Search successful",
		Data:    searchResponse,
	}

	// Respond with the search results
	writeJSONResponse(w, response)
}
