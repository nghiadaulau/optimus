package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/hyperledger/fabric-gateway/pkg/client"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type TodoItem struct {
	ID          string    `json:"ID"`
	Description string    `json:"Description"`
	Owner       string    `json:"Owner"`
	Status      string    `json:"Status"`
	StartDate   time.Time `json:"StartDate"`
	EndDate     time.Time `json:"EndDate"`
	Priority    int       `json:"Priority"`
}

type AllTodoItems []struct {
	ID          string    `json:"ID"`
	Description string    `json:"Description"`
	Owner       string    `json:"Owner"`
	Status      string    `json:"Status"`
	StartDate   time.Time `json:"StartDate"`
	EndDate     time.Time `json:"EndDate"`
	Priority    int       `json:"Priority"`
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
	if function == "GetAllTodoItems" {
		var responseData AllTodoItems
		// Replace YourStructType with the actual type of your response
		if err := json.Unmarshal(evaluateResponse, &responseData); err != nil {
			response := APIResponse{
				Status:  "error",
				Message: "Query error",
				Data:    TodoItem{}}
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

		var responseData TodoItem
		// Replace YourStructType with the actual type of your response
		if err := json.Unmarshal(evaluateResponse, &responseData); err != nil {
			response := APIResponse{
				Status:  "error",
				Message: "Query error",
				Data:    TodoItem{}}
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
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
