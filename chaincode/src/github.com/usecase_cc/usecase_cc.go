package main

import (
	"log"
	"errors"
	"strconv"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Chaincode example simple Chaincode implementation
type Chaincode struct {
}

var CHANNEL_ENV string
var outC string
var TxID string
var sum string
var out string

func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	CHANNEL_ENV = stub.GetChannelID()

	_, args := stub.GetFunctionAndParameters()
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		return shim.Error(err.Error())
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error(err.Error())
	}

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	var result string
	function, args := stub.GetFunctionAndParameters()

	if function == "get" {
		// Make payment of X units from A to B
		cc.get(stub, args)
	} else if function == "set" {
		// Query an entity from its state
		cc.set(stub, args)
	}
	return shim.Success([]byte(result))
}

// Transaction makes payment of X units from A to B
func (cc *Chaincode) set(stub shim.ChaincodeStubInterface, args []string) (string, error){
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value

	log.SetFlags(0)
	uuid := uuidgen()
	TxID = stub.GetTxID()
	timestamp := timeNow()
	log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][set] ex02 set")
	re := captureOutput(func(){
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][set] ex02 set")
	})
	invokeArgs := prepareToInvoke(uuid, re)
	stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

	if len(args) != 3 {
		log.SetFlags(0)
		uuid := uuidgen()
		TxID = stub.GetTxID()
		timestamp := timeNow()
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][valueIssuer] Incorrect number of arguments. Expecting 3")
		re = captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][valueIssuer] Incorrect number of arguments. Expecting 3")
		})
		invokeArgs := prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)

	if err != nil {
		log.SetFlags(0)
		uuid := uuidgen()
		TxID = stub.GetTxID()
		timestamp := timeNow()
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][stateIssuer] Failed to get state")
		re = captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][stateIssuer] Failed to get state")
		})
		invokeArgs := prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

		return "" , errors.New(ERRORGetState)
	}

	if Avalbytes == nil {
		log.SetFlags(0)
		uuid = uuidgen()
		TxID = stub.GetTxID()
		timestamp = timeNow()
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][idIssuer] Entity not found")
		re = captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][idIssuer] Entity not found")
		})
		invokeArgs := prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

		return "" , errors.New(ERRORnotID)	
	}

	Aval, _ = strconv.Atoi(string(Avalbytes))
	Bvalbytes, err := stub.GetState(B)

	if err != nil {
		log.SetFlags(0)
		uuid = uuidgen()
		TxID = stub.GetTxID()
		timestamp = timeNow()
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][stateIssuer] Failed to get state")
		re = captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][stateIssuer] Failed to get state")
		})
		invokeArgs := prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

		return "" , errors.New(ERRORGetState)
	}

	if Bvalbytes == nil {
		log.SetFlags(0)
		uuid = uuidgen()
		TxID = stub.GetTxID()
		timestamp = timeNow()
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][idIssuer] Entity not found")
		re = captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][idIssuer] Entity not found")
		})
		invokeArgs := prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

		return "" , errors.New(ERRORnotID)
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		log.SetFlags(0)
		uuid = uuidgen()
		TxID = stub.GetTxID()
		timestamp = timeNow()
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][valueIssuer] Invalid transaction amount, expecting a integer value")
		re = captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][valueIssuer] Invalid transaction amount, expecting a integer value")
		})
		invokeArgs = prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

		return "" , errors.New(ERRORParsingData)	
	}

	Aval = Aval + X
	Bval = Bval + X
	A = strconv.Itoa(Aval)
	B = strconv.Itoa(Bval)
	log.SetFlags(0)
	uuid = uuidgen()
	TxID = stub.GetTxID()
	timestamp = timeNow()
	log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][Transaction] Aval ="+A+" Bval ="+B+" after performing the transaction")
	re = captureOutput(func(){
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][Transaction] Aval ="+A+" Bval ="+B+" after performing the transaction")
	})
	invokeArgs = prepareToInvoke(uuid, re)
	stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		log.SetFlags(0)
		uuid = uuidgen()
		TxID = stub.GetTxID()
		timestamp = timeNow()
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][stateIssuer] Failed to write the state back to the ledger")
		re := captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][stateIssuer] Failed to write the state back to the ledger")
		})
		invokeArgs := prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

		return "" , errors.New(ERRORPutState)	
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		log.SetFlags(0)
		uuid = uuidgen()
		TxID = stub.GetTxID()
		timestamp = timeNow()
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][stateIssuer] Failed to write the state back to the ledger")
		re = captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][stateIssuer] Failed to write the state back to the ledger")
		})
		invokeArgs := prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		
		return "" , errors.New(ERRORPutState)
	}

	log.SetFlags(0)
	uuid = uuidgen()
	TxID = stub.GetTxID()
	timestamp = timeNow()
	log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][Transaction] Transaction makes payment of X units from A to B")
	re = captureOutput(func(){
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][Transaction] Transaction makes payment of X units from A to B")
	})
	invokeArgs = prepareToInvoke(uuid, re)
	stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

	return string(X) , errors.New("")
}

// query callback representing the query of a chaincode
func (cc *Chaincode) get(stub shim.ChaincodeStubInterface, args []string) (string, error){
	var A string // Entities
	var err error

	if len(args) != 1 {
		log.SetFlags(0)
		uuid := uuidgen()
		TxID = stub.GetTxID()
		timestamp := timeNow()
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc] Incorrect number of arguments. Expecting name of the person to query")
		re := captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc] Incorrect number of arguments. Expecting name of the person to query")
		})		
		invokeArgs := prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		return "" , errors.New(ERRORWrongNumberArgs)
	}

	A = args[0]
	Avalbytes, err := stub.GetState(A) // Get the state from the ledger
	if err != nil {
		log.SetFlags(0)
		uuid := uuidgen()
		TxID = stub.GetTxID()
		timestamp := timeNow()
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][get] "+jsonResp)
		re := captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][get] "+jsonResp)
		})
		invokeArgs := prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		
		return "" , errors.New(ERRORGetState)	
	}

	if Avalbytes == nil {
		log.SetFlags(0)
		uuid := uuidgen()
		TxID = stub.GetTxID()
		timestamp := timeNow()
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][get] "+jsonResp)
		re := captureOutput(func(){
			log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][usecase_cc][get] "+jsonResp)
		})
		invokeArgs := prepareToInvoke(uuid, re)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

		return "" , errors.New(ERRORParsingData)
	}

	log.SetFlags(0)
	uuid := uuidgen()
	TxID = stub.GetTxID()
	timestamp := timeNow()
	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][Get] Query Response: "+jsonResp)
	re := captureOutput(func(){
		log.Println("["+timestamp+"]["+uuid+"]["+CHANNEL_ENV+"]["+TxID+"][Get] Query Response: "+jsonResp)
	})
	invokeArgs := prepareToInvoke(uuid, re)
	stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

	return string(Avalbytes) , errors.New(ERRORParsingData)
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		panic(err)
	}
}