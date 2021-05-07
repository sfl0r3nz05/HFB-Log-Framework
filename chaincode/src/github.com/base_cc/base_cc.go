package main

import (
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Chaincode struct {
}

func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
   fmt.Println("ex02 Init")
   _, args := stub.GetFunctionAndParameters()
   if len(args) != 2 {
	   return shim.Error("Incorrect arguments. Expecting a key and a value")
   }

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}
	return shim.Success(nil)
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// Extract the function and args from the transaction proposal
	function, args := stub.GetFunctionAndParameters()
	var result string
	var err error

	if function == "set" {
		result, err = cc.set(stub, args)
	} else if function == "get" {
		result, err = cc.get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	// Return the result as success payload
	return shim.Success([]byte(result))
}

func (cc *Chaincode) set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
   fmt.Println("ex02 set")

	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}
	
	fmt.Println(args[0])
	fmt.Println(args[1])
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
}

// Get returns the value of the specified asset key
func (cc *Chaincode) get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
   fmt.Println("ex02 get")

	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	fmt.Println(string(value))
	return string(value), nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}