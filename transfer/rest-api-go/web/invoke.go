package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// InvokeResponse represents the structure of the invoke API response.
type InvokeResponse struct {
	Status         string `json:"status"`
	Message        string `json:"message"`
	TransactionID  string `json:"transaction_id,omitempty"`
	ResponseResult string `json:"response_result,omitempty"`
}

// writeJSONResponse writes a JSON response to the http.ResponseWriter.
func writeInvokeJSONResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// writeJSONError writes a JSON error response to the http.ResponseWriter.
func writeJSONError(w http.ResponseWriter, errorMessage string) {
	response := InvokeResponse{
		Status:  "error",
		Message: errorMessage,
	}
	writeInvokeJSONResponse(w, response)
}

// Invoke handles chaincode invoke requests.
func (setup *OrgSetup) Invoke(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Invoke request")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeJSONError(w, fmt.Sprintf("Error reading request body: %s", err))
		return
	}

	// Define a struct to unmarshal the JSON request body
	var request struct {
		ChainCodeID string   `json:"chaincodeid"`
		ChannelID   string   `json:"channelid"`
		Function    string   `json:"function"`
		Args        []string `json:"args"`
	}

	// Unmarshal the JSON into the request struct
	err = json.Unmarshal(body, &request)
	if err != nil {
		writeJSONError(w, fmt.Sprintf("Error unmarshalling JSON: %s", err))
		return
	}

	// Extract values from the request struct
	chainCodeName := request.ChainCodeID
	channelID := request.ChannelID
	function := request.Function
	args := request.Args

	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)

	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	txnProposal, err := contract.NewProposal(function, client.WithArguments(args...))
	if err != nil {
		writeJSONError(w, fmt.Sprintf("Error creating txn proposal: %s", err))
		return
	}
	txnEndorsed, err := txnProposal.Endorse()
	if err != nil {
		writeJSONError(w, fmt.Sprintf("Error endorsing txn: %s", err))
		return
	}
	txnCommitted, err := txnEndorsed.Submit()
	if err != nil {
		writeJSONError(w, fmt.Sprintf("Error submitting transaction: %s", err))
		return
	}

	response := InvokeResponse{
		Status:         "success",
		Message:        "Invoke successful",
		TransactionID:  txnCommitted.TransactionID(),
		ResponseResult: string(txnEndorsed.Result()),
	}

	writeInvokeJSONResponse(w, response)
}
