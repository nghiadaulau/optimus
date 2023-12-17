package chaincode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a TodoList
type SmartContract struct {
	contractapi.Contract
}

// TodoItem describes basic details of what makes up a Todo item
type TodoItem struct {
	ID          string    `json:"ID"`
	Description string    `json:"Description"`
	Owner       string    `json:"Owner"`
	Status      string    `json:"Status"`
	StartDate   time.Time `json:"StartDate"`
	EndDate     time.Time `json:"EndDate"`
	Priority    int       `json:"Priority"`
}

// InitLedger adds a base set of todo items to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	todoItems := []TodoItem{
		{ID: "todo1", Description: "Buy groceries", Owner: "Tomoko", Status: "Pending", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 7), Priority: 1},
		{ID: "todo2", Description: "Finish report", Owner: "Brad", Status: "InProgress", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 14), Priority: 2},
		{ID: "todo3", Description: "Exercise", Owner: "Jin Soo", Status: "Pending", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 5), Priority: 3},
		{ID: "todo4", Description: "Read a book", Owner: "Max", Status: "Completed", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 10), Priority: 1},
	}

	for _, item := range todoItems {
		itemJSON, err := json.Marshal(item)
		if err != nil {
			return fmt.Errorf("failed to marshal todo item: %v", err)
		}

		err = ctx.GetStub().PutState(item.ID, itemJSON)
		if err != nil {
			return fmt.Errorf("failed to put todo item to world state: %v", err)
		}
	}

	return nil
}

// CreateTodoItem issues a new todo item to the world state with given details.
func (s *SmartContract) CreateTodoItem(ctx contractapi.TransactionContextInterface, id string, description string, owner string, status string, startDate time.Time, endDate time.Time, priority int) error {
	exists, err := s.TodoItemExists(ctx, id)
	if err != nil {
		return fmt.Errorf("error checking if todo item exists: %v", err)
	}
	if exists {
		return fmt.Errorf("the todo item %s already exists", id)
	}

	item := TodoItem{
		ID:          id,
		Description: description,
		Owner:       owner,
		Status:      status,
		StartDate:   startDate,
		EndDate:     endDate,
		Priority:    priority,
	}
	itemJSON, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal todo item: %v", err)
	}

	err = ctx.GetStub().PutState(id, itemJSON)
	if err != nil {
		return fmt.Errorf("failed to put todo item to world state: %v", err)
	}

	return nil
}

// ReadTodoItem returns the todo item stored in the world state with given id.
func (s *SmartContract) ReadTodoItem(ctx contractapi.TransactionContextInterface, id string) (*TodoItem, error) {
	itemJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if itemJSON == nil {
		return nil, fmt.Errorf("the todo item %s does not exist", id)
	}

	var item TodoItem
	err = json.Unmarshal(itemJSON, &item)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal todo item: %v", err)
	}

	return &item, nil
}

// UpdateTodoItem updates an existing todo item in the world state with provided parameters.
func (s *SmartContract) UpdateTodoItem(ctx contractapi.TransactionContextInterface, id string, description string, owner string, status string, startDate time.Time, endDate time.Time, priority int) error {
	exists, err := s.TodoItemExists(ctx, id)
	if err != nil {
		return fmt.Errorf("error checking if todo item exists: %v", err)
	}
	if !exists {
		return fmt.Errorf("the todo item %s does not exist", id)
	}

	// Overwriting original todo item with new todo item
	item := TodoItem{
		ID:          id,
		Description: description,
		Owner:       owner,
		Status:      status,
		StartDate:   startDate,
		EndDate:     endDate,
		Priority:    priority,
	}
	itemJSON, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal todo item: %v", err)
	}

	err = ctx.GetStub().PutState(id, itemJSON)
	if err != nil {
		return fmt.Errorf("failed to put todo item to world state: %v", err)
	}

	return nil
}

// DeleteTodoItem deletes a given todo item from the world state.
func (s *SmartContract) DeleteTodoItem(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.TodoItemExists(ctx, id)
	if err != nil {
		return fmt.Errorf("error checking if todo item exists: %v", err)
	}
	if !exists {
		return fmt.Errorf("the todo item %s does not exist", id)
	}

	err = ctx.GetStub().DelState(id)
	if err != nil {
		return fmt.Errorf("failed to delete todo item from world state: %v", err)
	}

	return nil
}

// TodoItemExists returns true when todo item with given ID exists in world state
func (s *SmartContract) TodoItemExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	itemJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return itemJSON != nil, nil
}

// GetAllTodoItems returns all todo items found in world state
func (s *SmartContract) GetAllTodoItems(ctx contractapi.TransactionContextInterface) ([]*TodoItem, error) {
	// Range query with empty string for startKey and endKey does an
	// open-ended query of all todo items in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get state by range: %v", err)
	}
	defer resultsIterator.Close()

	var items []*TodoItem
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next query response: %v", err)
		}

		var item TodoItem
		err = json.Unmarshal(queryResponse.Value, &item)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal todo item: %v", err)
		}
		items = append(items, &item)
	}

	return items, nil
}
