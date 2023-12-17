package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Asset struct {
	AppraisedValue int    `json:"AppraisedValue"`
	Color          string `json:"Color"`
	ID             string `json:"ID"`
	Owner          string `json:"Owner"`
	Size           int    `json:"Size"`
}

type AllAsset []struct {
	AppraisedValue int    `json:"AppraisedValue"`
	Color          string `json:"Color"`
	ID             string `json:"ID"`
	Owner          string `json:"Owner"`
	Size           int    `json:"Size"`
}

// Query handles chaincode query requests.
func (setup OrgSetup) Query(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Query request")
	queryParams := r.URL.Query()
	chainCodeName := queryParams.Get("chaincodeid")
	channelID := queryParams.Get("channelid")
	function := queryParams.Get("function")
	args := r.URL.Query()["args"]
	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)
	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	evaluateResponse, err := contract.EvaluateTransaction(function, args...)
	if err != nil {
		response := APIResponse{
			Status:  "error",
			Message: fmt.Sprintf("Error: %s", err),
		}
		writeJSONResponse(w, response)
		return
	}
	if function == "GetAllAssets" {
		var responseData AllAsset
		// Replace YourStructType with the actual type of your response
		if err := json.Unmarshal(evaluateResponse, &responseData); err != nil {
			response := APIResponse{
				Status:  "error",
				Message: "Query error",
				Data:    Asset{}}
			writeJSONResponse(w, response)
			return
		}

		response := APIResponse{
			Status:  "success",
			Message: "Query successful",
			Data:    responseData,
		}
		writeJSONResponse(w, response)
	} else {

		var responseData Asset
		// Replace YourStructType with the actual type of your response
		if err := json.Unmarshal(evaluateResponse, &responseData); err != nil {
			response := APIResponse{
				Status:  "error",
				Message: "Query error",
				Data:    Asset{}}
			writeJSONResponse(w, response)
			return
		}

		response := APIResponse{
			Status:  "success",
			Message: "Query successful",
			Data:    responseData,
		}
		writeJSONResponse(w, response)
	}
}

// writeJSONResponse writes a JSON response to the http.ResponseWriter.
func writeJSONResponse(w http.ResponseWriter, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
