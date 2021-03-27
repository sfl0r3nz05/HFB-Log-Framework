package main

import (
	"os"
	"errors"
	"strconv"
	"io/ioutil"
	"crypto/sha256"
	log "github.com/log"
	b64 "encoding/base64"
	//log "github.com/sirupsen/logrus"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Chaincode example simple Chaincode implementation
type Chaincode struct {
}

const logLevel string = "DEBUG"
var CHANNEL_ENV string
var outC string
var sum string
var out string

func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	log.Init("DEBUG")
	CHANNEL_ENV = stub.GetChannelID()

	log.Infof("[%s][%s][usecase_cc][Init] ex02 Init",uuidgen(), CHANNEL_ENV)

	_, args := stub.GetFunctionAndParameters()
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		//return shim.Error("Incorrect argument numbers. Expecting 4")
		log.Errorf("[%s][%s][usecase_cc][valueIssuer] Incorrect argument numbers. Expecting 4: %v",uuidgen(), CHANNEL_ENV, err.Error())
		return shim.Error(err.Error())
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		//return shim.Error("Expecting integer value for asset holding")
		log.Errorf("[%s][%s][usecase_cc][valueIssuer] Expecting integer value for asset holding: %v",uuidgen(), CHANNEL_ENV, err.Error())
		return shim.Error(err.Error())
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		//return shim.Error("Expecting integer value for asset holding")
		log.Errorf("[%s][%s][usecase_cc][valueIssuer] Expecting integer value for asset holding: %v",uuidgen(), CHANNEL_ENV, err.Error())
		return shim.Error(err.Error())
	}
	//fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)
	log.Infof("[%s][%s][usecase_cc][Init] Initialize the chaincode with Aval = %d, Bval = %d",uuidgen(), CHANNEL_ENV, Aval, Bval)

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		//return shim.Error(err.Error())
		log.Errorf("[%s][%s][usecase_cc][stateIssuer] Error in writing the state to the ledger: %v",uuidgen(), CHANNEL_ENV, err.Error())
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		//return shim.Error(err.Error())
		log.Errorf("[%s][%s][usecase_cc][stateIssuer] Error in writing the state to the ledger: %v",uuidgen(), CHANNEL_ENV, err.Error())
		return shim.Error(err.Error())
	}

	log.Infof("[%s][%s][usecase_cc][PutState] Succeed to write the state to the ledger",uuidgen(), CHANNEL_ENV)
	return shim.Success(nil)
}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	//////////////////////////////LOG1///////////////////////////////////////
	uuid := uuidgen()
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	log.Infof("[%s][%s][usecase_cc][Invoke] ex02 Invoke", uuid, CHANNEL_ENV)
	h := sha256.New()
	h.Write([]byte(out))
	b := h.Sum(nil)
	sEnc := b64.StdEncoding.EncodeToString(b)
	params := []string{"set", uuid, sEnc}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {invokeArgs[i] = []byte(arg)}
	stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
	//////////////////////////////LOG1///////////////////////////////////////

	var err error
	var result string
	function, args := stub.GetFunctionAndParameters()
	if function == "set" {
		// Make payment of X units from A to B
		result, err = cc.set(stub, args)
		if err != nil {
			log.Errorf("[%s][%s][set] Error %v",uuidgen(), CHANNEL_ENV, err)
		}
	} else if function == "delete" {
		// Deletes an entity from its state
		result, err = cc.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in set
		result, err = cc.query(stub, args)
	}
	return shim.Success([]byte(result))
}

// Transaction makes payment of X units from A to B
func (cc *Chaincode) set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error

	//////////////////////////////LOG3///////////////////////////////////////
	uuid := uuidgen()
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	log.Infof("[%s][%s][usecase_cc][set] ex02 set", uuid, CHANNEL_ENV)
	h := sha256.New()
	h.Write([]byte(out))
	b := h.Sum(nil)
	sEnc := b64.StdEncoding.EncodeToString(b)
	params := []string{"set", uuid, sEnc}
	invokeArgs := make([][]byte, len(params))
	for i, arg := range params {invokeArgs[i] = []byte(arg)}
	stub.InvokeChaincode("base_cc", invokeArgs, CHANNEL_ENV)
	//////////////////////////////LOG3///////////////////////////////////////

	if len(args) != 3 {
		//return shim.Error("Incorrect number of arguments. Expecting 3")
		log.Errorf("[%s][%s][usecase_cc][valueIssuer] Incorrect number of arguments. Expecting 3",uuidgen(), CHANNEL_ENV)
		return "" , errors.New(ERRORWrongNumberArgs)
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		//return shim.Error("Failed to get state")
		log.Errorf("[%s][%s][usecase_cc][stateIssuer] Failed to get state",uuidgen(), CHANNEL_ENV)
		return "" , errors.New(ERRORGetState)
	}
	if Avalbytes == nil {
		//return shim.Error("Entity not found")	
		log.Errorf("[%s][%s][usecase_cc][idIssuer] Entity not found",uuidgen(), CHANNEL_ENV)	
		return "" , errors.New(ERRORnotID)	
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		//return shim.Error("Failed to get state")
		log.Errorf("[%s][%s][usecase_cc][stateIssuer] Failed to get state",uuidgen(), CHANNEL_ENV)
		return "" , errors.New(ERRORGetState)
	}
	if Bvalbytes == nil {
		//return shim.Error("Entity not found")
		log.Errorf("[%s][%s][usecase_cc][idIssuer] Entity not found",uuidgen(), CHANNEL_ENV)	
		return "" , errors.New(ERRORnotID)
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		//return shim.Error("Invalid transaction amount, expecting a integer value")
		log.Errorf("[%s][%s][usecase_cc][valueIssuer] Invalid transaction amount, expecting a integer value",uuidgen(), CHANNEL_ENV)
		return "" , errors.New(ERRORParsingData)	
	}
	Aval = Aval - X
	Bval = Bval + X
	//fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)
	log.Infof("[%s][%s][usecase_cc][Transaction] Aval = %d, Bval = %d after performing the transaction",uuidgen(), CHANNEL_ENV, Aval, Bval)	

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		//return shim.Error(err.Error())
		log.Errorf("[%s][%s][usecase_cc][stateIssuer] Failed to write the state back to the ledger",uuidgen(), CHANNEL_ENV)
		return "" , errors.New(ERRORPutState)	
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		//return shim.Error(err.Error())
		log.Errorf("[%s][%s][usecase_cc][stateIssuer] Failed to write the state back to the ledger",uuidgen(), CHANNEL_ENV)
		return "" , errors.New(ERRORPutState)
	}

	payloadAsBytes := []byte(strconv.Itoa(Bval))	
	log.Infof("[%s][%s][usecase_cc][Transaction] Transaction makes payment of X units from A to B",uuidgen(), CHANNEL_ENV)
	return string(payloadAsBytes) , errors.New("")
}

// Deletes an entity from state
func (cc *Chaincode) delete(stub shim.ChaincodeStubInterface, args []string) (string, error){
	if len(args) != 1 {
		log.Errorf("[%s][%s][usecase_cc][valueIssuer] Incorrect number of arguments. Expecting 1",uuidgen(), CHANNEL_ENV)
		return "" , errors.New(ERRORWrongNumberArgs)
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		log.Errorf("[%s][%s][usecase_cc][stateIssuer] Failed to delete state",uuidgen(), CHANNEL_ENV)
		return "" , errors.New(ERRORDelState)
	}

	log.Infof("[%s][%s][usecase_cc][DelState] Succeed to delete an entity from state",uuidgen(), CHANNEL_ENV)
	return "", errors.New("")
}

// query callback representing the query of a chaincode
func (cc *Chaincode) query(stub shim.ChaincodeStubInterface, args []string) (string, error){
	var A string // Entities
	var err error

	if len(args) != 1 {
		//return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
		log.Errorf("[%s][%s][usecase_cc][valueIssuer] Incorrect number of arguments. Expecting name of the person to query",uuidgen(), CHANNEL_ENV)
		return "" , errors.New(ERRORWrongNumberArgs)
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		//return shim.Error(jsonResp)
		log.Errorf("[%s][%s][usecase_cc][stateIssuer] %s",uuidgen(), CHANNEL_ENV,jsonResp)
		return "" , errors.New(ERRORGetState)	
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		//return shim.Error(jsonResp)
		log.Errorf("[%s][%s][usecase_cc][valueIssuer] %s",uuidgen(), CHANNEL_ENV,jsonResp)	
		return "" , errors.New(ERRORParsingData)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	//fmt.Printf("Query Response:%s\n", jsonResp)
	log.Infof("[%s][%s][usecase_cc][Query] Query Response: %s",uuidgen(), CHANNEL_ENV,jsonResp)
	return string(Avalbytes) , errors.New(ERRORParsingData)
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		panic(err)
	}
}
