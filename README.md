Step by step to run project.
===============
# Requirement for Hyperledger Fabric 2.5.x
1. Git
2. Docker & Docker-compose
3. Nodejs >= 18
4. JQ >= 1.6
5. cURL

## Step 1:
Get install script:

```shell
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh
```
Format install command:
```text
./install-fabric.sh -h
Usage: ./install-fabric.sh [-f|--fabric-version <arg>] [-c|--ca-version <arg>] <comp-1> [<comp-2>] ... [<comp-n>] ...
        <comp>: Component to install one or more of  d[ocker]|b[inary]|s[amples]. If none specified, all will be installed
        -f, --fabric-version: FabricVersion (default: '2.5.5')
        -c, --ca-version: Fabric CA Version (default: '1.5.7')
```
Choosing which version
```shell
./install-fabric.sh --fabric-version 2.5.5 binary
```
## Step 2:
The Asset Transfer (basic) sample demonstrates how to create, update, and query assets. It involves the following two components:
```text
asset-transfer-basic/application-gateway-typescript
asset-transfer-basic/chaincode-typescript
```
Launch the blockchain network
---------------

```shell
cd networking
```
If you already have a test network running, bring it down to ensure the environment is clean.
```shell
./network.sh down
```
Launch the Fabric test network using the network.sh shell script.
```shell
./network.sh up createChannel -c mychannel -ca
```
Deploy the smart contract
-------------------------

```shell
./network.sh deployCC -ccn basic -ccp ../transfer/chaincode-go -ccl go
```

# Run Restful API with Golang
Preparation steps:
1. Access the folder /networking/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/
2. Create a cert.pem file from the *-cert.pem file
3. Access the transfer/rest-api-go directory and execute the command
```
go run main.go
```
Done!

sample api to test

API POST CREATE:
```curl
curl --location 'http://localhost:3000/invoke' \
--header 'content-type: application/json' \
--data '{
    "channelid": "mychannel",
    "chaincodeid": "basic",
    "function": "CreateTodoItem",
    "args": {
        "ID": "todo5",
        "Description": "Test Todo Item",
        "Owner": "John",
        "Status": "Pending",
        "StartDate": "2023-01-01T00:00:00Z",
        "EndDate": "2023-01-07T00:00:00Z",
        "Priority": 2
    }
}'
```
API GET ALL:
```curl
curl --location 'http://localhost:3000/query?channelid=mychannel&chaincodeid=basic&function=GetAllTodoItems'
```
API GET DETAIL:
```curl
curl --location 'http://localhost:3000/query?channelid=mychannel&chaincodeid=basic&function=ReadTodoItem&args=todo1'
```

API PUT UPDATE:
```curl
curl --location --request PUT 'http://localhost:3000/invoke' \
--header 'content-type: application/json' \
--data '{
    "channelid": "mychannel",
    "chaincodeid": "basic",
    "function": "UpdateTodoItem",
    "args": {
        "ID": "todo1",
        "Description": "Test Todo Item",
        "Owner": "John",
        "Status": "Pending",
        "StartDate": "2023-01-01T00:00:00Z",
        "EndDate": "2023-01-07T00:00:00Z",
        "Priority": 2
    }
}'
```

API PUT UPDATE:
```curl
curl --location --request DELETE 'http://localhost:3000/delete' \
--header 'content-type: application/json' \
--data '{
    "channelid": "mychannel",
    "chaincodeid": "basic",
    "function": "DeleteTodoItem",
    "args": [
        "todo1"
    ]
}'
```