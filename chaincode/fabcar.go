package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"

	// "encoding/gob"
	// "encoding/hex"
	"os"
	"fmt"
	"reflect"

	// "log"
	// "strings"
	// "math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"

	//"github.com/hyperledger/fabric-contract-api-go/contractapi"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

type serverConfig struct {
	CCID    string
	Address string
}

// SmartContract Define the Smart Contract structure
type SmartContract struct {
}

// SmartContract provides functions for authentication in a Hyperledger Fabric network
// type SmartContract struct {
// 	contractapi.Contract
// }

// User :  Define the user structure, with 4 properties.  Structure tags are used by encoding/json library
type User struct {
	VkycId        string        `json:"VkycId"`
	AadhaarNumber string        `json:"AadhaarNumber"`
	PanNumber     string        `json:"PanNumber"`
	Title         string        `json:"Title"`
	FirstName     string        `json:"FirstName"`
	MiddleName    string        `json:"MiddleName"`
	LastName      string        `json:"LastName"`
	Gender        string        `json:"Gender"`
	DateOfBirth   string        `json:"DateOfBirth"`
	Address       string        `json:"Address"`
	State         string        `json:"State"`
	District      string        `json:"District"`
	City          string        `json:"City"`
	Pincode       string        `json:"Pincode"`
	Mobilenumber  string        `json:"Mobilenumber"`
	Email         string        `json:"Email"`
	Language      string        `json:"Language"`
	AadhaarCard   string        `json:"AadhaarCard"`
	PanCard       string        `json:"PanCard"`
	Image         string        `json:"Image"`
	Audio_file    string        `json:"Audio_file"`
	Video_file    string        `json:"Video_file"`
	BankDetails   []BankDetails `json:"BankDetails"`
}

type BankDetails struct {
	BankName      string `json:"BankName"`
	BankCode      string `json:"BankCode"`
	Consentstatus string `json:"Consentstatus"`
	Validity      string `json:"validity"`
	BankEmail     string `json:"BankEmail"`
	Time          string `json:"Time"`
	Date          string `json:"Date"`
}

// Init ;  Method for initializing smart contract
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

var logger = flogging.MustGetLogger("fabcar_cc")

// Invoke :  Method for INVOKING smart contract
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	logger.Infof("Function name is:  %d", function)
	logger.Infof("Args length is : %d", len(args), args)

	switch function {
	case "queryUser":
		return s.queryUser(APIstub, args)
	case "createUser":
		return s.createUser(APIstub, args)
	case "sendConsent":
		return s.sendConsent(APIstub, args)
	case "queryAllUsers":
		return s.queryAllUsers(APIstub)
	case "updateUserData":
		return s.updateUserData(APIstub, args)
	case "getHistoryForUser":
		return s.getHistoryForUser(APIstub, args)
	case "queryUsersByVkycId":
		return s.queryUsersByVkycId(APIstub, args)
	case "test":
		return s.test(APIstub, args)
	default:
		return shim.Error("Invalid Smart Contract function name.")
	}

	// return shim.Error("Invalid Smart Contract function name.")
}

// func main() {
//     t:=test()
//     fmt.Println(t)
// }

func calculateChecksum(id []byte) byte {
	var sum byte
	for i := 0; i < len(id); i++ {
		sum ^= id[i]
	}
	return sum
}

func test1(n int) string {

	stringPermutation1 := "EZNDUAVXHTKCBQFYROPLJIGMWS"
	stringPermutation2 := "GLQPEZCIXASBTNHFUMDROWKYVJ"
	stringPermutation3 := "PZUKCRQLNMYFXIASEVBWHTJDOG"
	stringPermutation4 := "XNAEBJVTQCDGOMWSZLUHFIRYPK"
	stringPermutation5 := "VNIBWHQMSDYRPKCAFZGXJUOLTE"

	numberPermutation1 := "7960143582"
	numberPermutation2 := "6021974358"
	numberPermutation3 := "6240718359"
	numberPermutation4 := "7420891653"

	uniqueIDLength := 10

	// Map to store the generated IDs
	uniqueIDs := make(map[string]bool)

	for i := n; i < n+1; i++ {
		uniqueID := make([]byte, uniqueIDLength)

		stringIndex1 := i % len(stringPermutation1)
		stringIndex2 := (i / len(stringPermutation1)) % len(stringPermutation2)
		numberIndex1 := (i / len(stringPermutation1) / len(stringPermutation2)) % len(numberPermutation1)
		numberIndex2 := (i / len(stringPermutation1) / len(stringPermutation2) / len(numberPermutation1)) % len(numberPermutation2)
		stringIndex3 := (i / len(stringPermutation1) / len(stringPermutation2) / len(numberPermutation1) / len(numberPermutation2)) % len(stringPermutation3)
		stringIndex4 := (i / len(stringPermutation1) / len(stringPermutation2) / len(numberPermutation1) / len(numberPermutation2) / len(stringPermutation3)) % len(stringPermutation4)
		numberIndex3 := (i / len(stringPermutation1) / len(stringPermutation2) / len(numberPermutation1) / len(numberPermutation2) / len(stringPermutation3) / len(stringPermutation4)) % len(numberPermutation3)
		numberIndex4 := (i / len(stringPermutation1) / len(stringPermutation2) / len(numberPermutation1) / len(numberPermutation2) / len(stringPermutation3) / len(stringPermutation4) / len(numberPermutation3)) % len(numberPermutation4)

		// Loop to generate the unique ID
		for i := 0; i < uniqueIDLength-1; i++ {
			switch i {
			case 0:
				uniqueID[i] = stringPermutation1[stringIndex1]
			case 1:
				uniqueID[i] = stringPermutation2[stringIndex2]
			case 2:
				uniqueID[i] = numberPermutation1[numberIndex1]
			case 3:
				uniqueID[i] = numberPermutation2[numberIndex2]
			case 4:
				uniqueID[i] = stringPermutation3[stringIndex3]
			case 5:
				uniqueID[i] = stringPermutation4[stringIndex4]
			case 6:
				uniqueID[i] = numberPermutation3[numberIndex3]
			case 7:
				uniqueID[i] = numberPermutation4[numberIndex4]
			case 8:
				uniqueID[i] = stringPermutation5[0]
			}
		}

		checksum := calculateChecksum(uniqueID[:uniqueIDLength-1]) % 10
		uniqueID[9] = checksum + '0'

		// Convert the uniqueID to a string to store it in the map
		idString := string(uniqueID)

		// Check if the ID already exists in the map
		if _, exists := uniqueIDs[idString]; exists {
			fmt.Printf("Duplicate ID found: %s\n", idString)
			return "duplicate id"
		} else {
			uniqueIDs[idString] = true
			fmt.Println("Unique ID:", idString)

			return idString
		}
		// return "gujg"

	}
	return "VKYCID"

}

// Function to add bank details to the User
func AddBankDetails(user *User, BankName, BankCode, Consentstatus, Validity, BankEmail, Time, Date string) {
	newBank := BankDetails{
		BankName:      BankName,
		BankCode:      BankCode,
		Consentstatus: Consentstatus,
		Validity:      Validity,
		BankEmail:     BankEmail,
		Time:          Time,
		Date:          Date,
	}
	user.BankDetails = append(user.BankDetails, newBank)
}

func (s *SmartContract) sendConsent(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	// VkycId := args[0]
	// BankName := args[1]
	// Consent := args[2]

	// var user = User{VkycId: data2, AadhaarNumber: args[0], PanNumber: args[1], Title: args[2], FirstName: args[3], MiddleName: args[4], LastName: args[5], Gender: args[6], DateOfBirth: args[7], Address: args[8], State: args[9],
	// 	District: args[10], City: args[11], Pincode: args[12], Mobilenumber: args[13], Email: args[14], Language: args[15], AadhaarCard: sEnc_AadhaarCard, PanCard: sEnc_PanCard, Image: sEnc_Image, Audio_file: sEnc_audio_file, Video_file: sEnc_Video_file, BankName: args[21], Consent: args[22]}

	// fmt.Println(2)

	userAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error("Error getting user data from the ledger")
	}

	if userAsBytes == nil {
		return shim.Error("User does not exist")
	}

	user := User{}
	err = json.Unmarshal(userAsBytes, &user)
	if err != nil {
		return shim.Error("Error unmarshalling user data")
	}

	// Add new bank details to the User struct
	AddBankDetails(&user, args[1], args[2], args[3], args[4], args[5], args[6], args[7])

	updatedUserAsBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error("Error marshalling updated user data")
	}

	// Put updated user data on the ledger
	err = APIstub.PutState(args[0], updatedUserAsBytes)
	if err != nil {
		return shim.Error("Error putting updated user data on the ledger")
	}

	return shim.Success(updatedUserAsBytes)
}

func (s *SmartContract) queryUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(userAsBytes)
}

func (s *SmartContract) test(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 16 {
		return shim.Error("Incorrect number of arguments. Expecting 16")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(userAsBytes)
}

type QueryResults struct {
	Key    string `json:"Key"`
	Record *User
}

func (s *SmartContract) createUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 21 {
		return shim.Error("Incorrect number of arguments. Expecting 21")
	}
	fmt.Println(0)

	startKey := ""
	endKey := ""

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)

	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	results := []QueryResults{}

	user1 := new(User)
	// user2 := new(User)

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		_ = json.Unmarshal(queryResponse.Value, user1)
		queryResult := QueryResults{Key: queryResponse.Key, Record: user1}
		results = append(results, queryResult)

		_ = json.Unmarshal(queryResponse.Value, user1)
		if user1.AadhaarNumber == args[0] {
			return shim.Error("This User Aadhar already exists")
		}
	}

	fmt.Println("this is result data", results, len(results), reflect.TypeOf(results))

	data1 := len(results)

	data2 := test1(data1)

	sEnc_AadhaarCard := b64.StdEncoding.EncodeToString([]byte(args[16]))
	sEnc_PanCard := b64.StdEncoding.EncodeToString([]byte(args[17]))
	sEnc_Image := b64.StdEncoding.EncodeToString([]byte(args[18]))
	sEnc_audio_file := b64.StdEncoding.EncodeToString([]byte(args[19]))
	sEnc_Video_file := b64.StdEncoding.EncodeToString([]byte(args[20]))

	var user = User{VkycId: data2, AadhaarNumber: args[0], PanNumber: args[1], Title: args[2], FirstName: args[3], MiddleName: args[4], LastName: args[5], Gender: args[6], DateOfBirth: args[7], Address: args[8], State: args[9],
		District: args[10], City: args[11], Pincode: args[12], Mobilenumber: args[13], Email: args[14], Language: args[15], AadhaarCard: sEnc_AadhaarCard, PanCard: sEnc_PanCard, Image: sEnc_Image, Audio_file: sEnc_audio_file, Video_file: sEnc_Video_file}

	fmt.Println(2)

	userAsBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error from marshalling", err.Error())
	}
	err = APIstub.PutState(data2, userAsBytes)
	if err != nil {
		fmt.Println("Error from putting state", err.Error())
	}
	fmt.Println(1)

	indexName := "VkycId~key"
	FirstNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{user.VkycId, data2})
	if err != nil {
		fmt.Println("Error from composite key", err.Error())
	}
	fmt.Println(3)
	if err != nil {
		return shim.Error(err.Error())
	}

	value := []byte{0x00}
	err = APIstub.PutState(FirstNameIndexKey, value)
	if err != nil {
		fmt.Println("Error from composite key", err.Error())
	}
	fmt.Println(4)

	return shim.Success(userAsBytes)
}

// 	return shim.Success(userAsBytes)
// }

func (S *SmartContract) queryUsersByVkycId(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}
	VkycId := args[0]
	fmt.Println(0)

	vKYC_IDAndIdResultIterator, err := APIstub.GetStateByPartialCompositeKey("VkycId~key", []string{VkycId})
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(1)

	defer vKYC_IDAndIdResultIterator.Close()

	var i int
	var id string

	var users []byte
	bArrayMemberAlreadyWritten := false

	users = append([]byte("["))
	fmt.Println(2)

	for i = 0; vKYC_IDAndIdResultIterator.HasNext(); i++ {
		responseRange, err := vKYC_IDAndIdResultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println(3, i)

		objectType, compositeKeyParts, err := APIstub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println(4, i)

		id = compositeKeyParts[1]
		userAsBytes, err := APIstub.GetState(id)
		fmt.Println(5, i)

		if bArrayMemberAlreadyWritten == true {
			newBytes := append([]byte(","), userAsBytes...)
			users = append(users, newBytes...)

		} else {
			// newBytes := append([]byte(","), usersAsBytes...)
			users = append(users, userAsBytes...)
		}
		fmt.Println(6, i)

		fmt.Printf("Found a user for index : %s user id : ", objectType, compositeKeyParts[0], compositeKeyParts[1])
		bArrayMemberAlreadyWritten = true

	}

	users = append(users, []byte("]")...)

	return shim.Success(users)
}

func (s *SmartContract) queryAllUsers(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := ""
	endKey := ""

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)

	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllUsers:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) restictedMethod(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	// get an ID for the client which is guaranteed to be unique within the MSP
	//id, err := cid.GetID(APIstub) -

	// get the MSP ID of the client's identity
	//mspid, err := cid.GetMSPID(APIstub) -

	// get the value of the attribute
	//val, ok, err := cid.GetAttributeValue(APIstub, "attr1") -

	// get the X509 certificate of the client, or nil if the client's identity was not based on an X509 certificate
	//cert, err := cid.GetX509Certificate(APIstub) -

	val, ok, err := cid.GetAttributeValue(APIstub, "role")
	if err != nil {
		// There was an error trying to retrieve the attribute
		shim.Error("Error while retriving attributes")
	}
	if !ok {
		// The client identity does not possess the attribute
		shim.Error("Client identity doesnot posses the attribute")
	}
	// Do something with the value of 'val'
	if val != "approver" {
		fmt.Println("Attribute role: " + val)
		return shim.Error("Only user with role as APPROVER have access this method!")
	} else {
		if len(args) != 1 {
			return shim.Error("Incorrect number of arguments. Expecting 1")
		}

		userAsBytes, _ := APIstub.GetState(args[0])
		return shim.Success(userAsBytes)
	}

}

func (s *SmartContract) updateUserData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 20 {
		return shim.Error("Incorrect number of arguments. Expecting 20")
	}

	userAsBytes, err := APIstub.GetState(args[0])
	if userAsBytes != nil {
		fmt.Println("This user already exists")

		// aadhaarCard := "/home/vkyc25/Downloads/vkyc record/shinde.jpeg"
		// panCard := "/home/vkyc25/Downloads/vkyc record/user1doc.pdf"
		// audio_file := "/home/vkyc25/Downloads/vkyc record/shinde.jpeg"

		// // Read entire JPG into byte slice.
		// reader := bufio.NewReader(aadhaarCard)
		// reader := bufio.NewReader(panCard)
		// reader := bufio.NewReader(audio_file)
		// sEnc_image, _ := ioutil.ReadAll(reader)
		// sEnc_document, _ := ioutil.ReadAll(reader)
		// sEnc_audio, _ := ioutil.ReadAll(reader)

		sEnc_AadhaarCard := b64.StdEncoding.EncodeToString([]byte(args[15]))
		sEnc_PanCard := b64.StdEncoding.EncodeToString([]byte(args[16]))
		sEnc_Image := b64.StdEncoding.EncodeToString([]byte(args[17]))
		sEnc_Audio_file := b64.StdEncoding.EncodeToString([]byte(args[18]))
		sEnc_Video_file := b64.StdEncoding.EncodeToString([]byte(args[19]))

		// userAsBytes, := APIstub.GetState(args[0])
		user := User{}

		json.Unmarshal(userAsBytes, &user)
		user.Title = args[1]
		user.FirstName = args[2]
		user.MiddleName = args[3]
		user.LastName = args[4]
		user.Gender = args[5]
		user.DateOfBirth = args[6]
		user.Address = args[7]
		user.State = args[8]
		user.District = args[9]
		user.City = args[10]
		user.Pincode = args[11]
		user.Mobilenumber = args[12]
		user.Email = args[13]
		user.Language = args[14]
		user.AadhaarCard = sEnc_AadhaarCard
		user.PanCard = sEnc_PanCard
		user.Image = sEnc_Image
		user.Audio_file = sEnc_Audio_file
		user.Video_file = sEnc_Video_file

		userAsBytes, _ = json.Marshal(user)
		APIstub.PutState(args[0], userAsBytes)

		return shim.Success(userAsBytes)
	} else if userAsBytes != nil {
		fmt.Println("This user does not exists", err.Error())
		return shim.Error("This user does not exists")
	}

	return shim.Success(userAsBytes)
}

func (t *SmartContract) getHistoryForUser(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	VkycId := args[0]

	resultsIterator, err := stub.GetHistoryForKey(VkycId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForUser returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

// See chaincode.env.example
	config := serverConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}
	
        smartContract := new(SmartContract)

	server := &shim.ChaincodeServer{
		CCID:    config.CCID,
		Address: config.Address,
		CC:      smartContract,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}
	
	if err := server.Start(); err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
	
}
