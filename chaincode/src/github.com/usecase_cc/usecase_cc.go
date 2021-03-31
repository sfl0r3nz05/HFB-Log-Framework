package main

import (
	"errors"
	"strconv"
	log "github.com/log"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Chaincode example simple Chaincode implementation
type Chaincode struct {
}

const logLevel string = "DEBUG"
var CHANNEL_ENV string
var outC string
var TxID string
var sum string
var out string

func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	log.Init("DEBUG")
	CHANNEL_ENV = stub.GetChannelID()
	TxID = stub.GetTxID()

	log.Infof("[%s][%s][%s][usecase_cc][Init] ex02 Init",uuidgen(), CHANNEL_ENV, TxID)

	_, args := stub.GetFunctionAndParameters()
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		//return shim.Error("Incorrect argument numbers. Expecting 4")
		log.Errorf("[%s][%s][%s][usecase_cc][valueIssuer] Incorrect argument numbers. Expecting 4: %v",uuidgen(), CHANNEL_ENV, TxID, err.Error())
		return shim.Error(err.Error())
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		//return shim.Error("Expecting integer value for asset holding")
		log.Errorf("[%s][%s][%s][usecase_cc][valueIssuer] Expecting integer value for asset holding: %v",uuidgen(), CHANNEL_ENV, TxID, err.Error())
		return shim.Error(err.Error())
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		//return shim.Error("Expecting integer value for asset holding")
		log.Errorf("[%s][%s][%s][usecase_cc][valueIssuer] Expecting integer value for asset holding: %v",uuidgen(), CHANNEL_ENV, TxID, err.Error())
		return shim.Error(err.Error())
	}
	//fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)
	log.Infof("[%s][%s][%s][usecase_cc][Init] Initialize the chaincode with Aval = %d, Bval = %d",uuidgen(), CHANNEL_ENV, TxID, Aval, Bval)

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		//return shim.Error(err.Error())
		log.Errorf("[%s][%s][%s][usecase_cc][stateIssuer] Error in writing the state to the ledger: %v",uuidgen(), CHANNEL_ENV, TxID, err.Error())
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		//return shim.Error(err.Error())
		log.Errorf("[%s][%s][%s][usecase_cc][stateIssuer] Error in writing the state to the ledger: %v",uuidgen(), CHANNEL_ENV, TxID, err.Error())
		return shim.Error(err.Error())
	}

	log.Infof("[%s][%s][%s][usecase_cc][PutState] Succeed to write the state to the ledger",uuidgen(), CHANNEL_ENV, TxID)
	return shim.Success(nil)
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	var result string

	uuid := uuidgen()
	log.Infof("[%s][%s][%s][usecase_cc][Invoke] ex02 Invoke", uuid, CHANNEL_ENV, TxID)
	function, args := stub.GetFunctionAndParameters()

	if function == "set" {
		// Make payment of X units from A to B
		cc.set(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		cc.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in set
		cc.query(stub, args)
	}
	return shim.Success([]byte(result))
}

// Transaction makes payment of X units from A to B
func (cc *Chaincode) set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value

	uuid := uuidgen()
	log.Infof("[%s][%s][%s][usecase_cc][set] ex02 set", uuid, CHANNEL_ENV, TxID)
	invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
	stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)


	if len(args) != 3 {
		uuid := uuidgen()
		log.Errorf("[%s][%s][%s][usecase_cc][valueIssuer] Incorrect number of arguments. Expecting 3", uuid, CHANNEL_ENV, TxID)
		invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		uuid := uuidgen()
		log.Errorf("[%s][%s][%s][usecase_cc][stateIssuer] Failed to get state", uuid, CHANNEL_ENV, TxID)
		invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		return "" , errors.New(ERRORGetState)
	}
	if Avalbytes == nil {
		uuid := uuidgen()	
		log.Errorf("[%s][%s][%s][usecase_cc][idIssuer] Entity not found", uuid, CHANNEL_ENV, TxID)
		invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		return "" , errors.New(ERRORnotID)	
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		uuid := uuidgen()
		log.Errorf("[%s][%s][%s][usecase_cc][stateIssuer] Failed to get state",	uuid, CHANNEL_ENV, TxID)
		invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		return "" , errors.New(ERRORGetState)
	}
	if Bvalbytes == nil {
		//return shim.Error("Entity not found")
		uuid := uuidgen()
		log.Errorf("[%s][%s][%s][usecase_cc][idIssuer] Entity not found", uuid, CHANNEL_ENV, TxID)
		invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		return "" , errors.New(ERRORnotID)
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		uuid := uuidgen()
		log.Errorf("[%s][%s][%s][usecase_cc][valueIssuer] Invalid transaction amount, expecting a integer value", uuid, CHANNEL_ENV, TxID)
		invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		return "" , errors.New(ERRORParsingData)	
	}
	Aval = Aval + X
	Bval = Bval + X
	uuid = uuidgen()
	log.Infof("[%s][%s][%s][usecase_cc][Transaction] Aval = %d, Bval = %d after performing the transaction", uuid, CHANNEL_ENV, TxID, Aval, Bval)	
	invokeArgs = prepareToInvoke(uuid, TxID, CHANNEL_ENV)
	stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		uuid := uuidgen()
		log.Errorf("[%s][%s][%s][usecase_cc][stateIssuer] Failed to write the state back to the ledger", uuid, CHANNEL_ENV, TxID)
		invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		return "" , errors.New(ERRORPutState)	
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		uuid := uuidgen()
		log.Errorf("[%s][%s][%s][usecase_cc][stateIssuer] Failed to write the state back to the ledger", uuid, CHANNEL_ENV, TxID)
		invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		return "" , errors.New(ERRORPutState)
	}

	uuid = uuidgen()
	log.Infof("[%s][%s][%s][usecase_cc][Transaction] Transaction makes payment of X units from A to B", uuid, CHANNEL_ENV, TxID)
	invokeArgs = prepareToInvoke(uuid, TxID, CHANNEL_ENV)
	stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)

	return string(X) , errors.New("")
}

// Deletes an entity from state
func (cc *Chaincode) delete(stub shim.ChaincodeStubInterface, args []string) (string, error){
	if len(args) != 1 {
		uuid := uuidgen()
		log.Errorf("[%s][%s][%s][usecase_cc][valueIssuer] Incorrect number of arguments. Expecting 1", uuid, CHANNEL_ENV, TxID)
		invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		return "" , errors.New(ERRORWrongNumberArgs)
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		uuid := uuidgen()
		log.Errorf("[%s][%s][%s][usecase_cc][stateIssuer] Failed to delete state", uuid, CHANNEL_ENV, TxID)
		invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
		stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
		return "" , errors.New(ERRORDelState)
	}

	uuid := uuidgen()
	log.Infof("[%s][%s][%s][usecase_cc][DelState] Succeed to delete an entity from state", uuid, CHANNEL_ENV, TxID)
	invokeArgs := prepareToInvoke(uuid, TxID, CHANNEL_ENV)
	stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
	return "", errors.New("")
}

// query callback representing the query of a chaincode
func (cc *Chaincode) query(stub shim.ChaincodeStubInterface, args []string) (string, error){
	var A string // Entities
	var err error

	if len(args) != 1 {
		//return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
		log.Errorf("[%s][%s][%s][usecase_cc][valueIssuer] Incorrect number of arguments. Expecting name of the person to query",uuidgen(), CHANNEL_ENV, TxID)
		return "" , errors.New(ERRORWrongNumberArgs)
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		//return shim.Error(jsonResp)
		log.Errorf("[%s][%s][%s][usecase_cc][stateIssuer] %s",uuidgen(), CHANNEL_ENV, TxID, jsonResp)
		return "" , errors.New(ERRORGetState)	
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		//return shim.Error(jsonResp)
		log.Errorf("[%s][%s][%s][usecase_cc][valueIssuer] %s",uuidgen(), CHANNEL_ENV, TxID, jsonResp)	
		return "" , errors.New(ERRORParsingData)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	log.Infof("[%s][%s][%s][usecase_cc][Query] Query Response: %s",uuidgen(), CHANNEL_ENV, TxID, jsonResp)
	return string(Avalbytes) , errors.New(ERRORParsingData)
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		panic(err)
	}
}