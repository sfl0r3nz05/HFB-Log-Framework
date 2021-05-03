 package main

 import (
	 "fmt"
	 "github.com/hyperledger/fabric/protos/peer"
	 "github.com/hyperledger/fabric/core/chaincode/shim"
 )

 type SimpleAsset struct {
 }

 func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
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

 func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	 // Extract the function and args from the transaction proposal
	 function, args := stub.GetFunctionAndParameters()
	 var result string
	 var err error
	 if function == "set" {
		 result, err = set(stub, args)
	 } else if function == "get" {
		 result, err = get(stub, args)
	 } 
	 if err != nil {
		return shim.Error(err.Error())
	}
	 // Return the result as success payload
	 return shim.Success([]byte(result))
 }
 
 func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
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
 func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
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
	 return string(value), nil
 }
 
 // main function starts up the chaincode in the container during instantiate
 func main() {
	 if err := shim.Start(new(SimpleAsset)); err != nil {
		 fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	 }
 }