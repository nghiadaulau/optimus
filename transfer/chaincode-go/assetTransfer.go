/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"
	"optimus/transfer/chaincode-go/chaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating transfer chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting transfer chaincode: %v", err)
	}
}
