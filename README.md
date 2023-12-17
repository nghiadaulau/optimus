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
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go
```

# Run Restful API with Golang



