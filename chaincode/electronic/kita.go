package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

type StaffChaincode struct {
}

type Staff struct {
	StaffId   string `json:"staff_id"`
	StaffName string `json:"staff_name"`
	StaffTel  string `json:"staff_tel"`
	StaffType string `json:"staff_type"`
}

type Record struct {
	RecordId          string `json:"record_id"`
	RecordType        string `json:"record_type"`
	RecordValue       string `json:"record_value"`
	RecordStaffId     string `json:"record_staff_id"`
	RecordKitaId      string `json:"record_kita_id"`
	RecordDescription string `json:"record_description"`
	RecordTime        string `json:"record_time"`
}

func (t *StaffChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *StaffChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + fn)

	if fn == "initStaff" {
		return t.initStaff(stub, args)
	} else if fn == "addRecord" {
		return t.addRecord(stub, args)
	} else if fn == "queryRecordByID" {

	} else if fn == "deleteStaff" {

	} else if fn == "updateRecord" {

	}

	fmt.Println("invoke did not find func : " + fn)
	return shim.Error("received unknown function invocation")
}

func (t *StaffChaincode) initStaff(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("输入的参数个有误：请输入4个参数(StaffId,StaffName,StaffTel,StaffType)")
	}

	staffId := args[0]
	staffName := args[1]
	staffTel := args[2]
	staffType := args[3]

	staff := &Staff{staffId, staffName, staffTel, staffType}

	staffKey, err := stub.CreateCompositeKey("Staff", []string{"staff", staffId})

	if err != nil {
		return shim.Error(err.Error())
	}

	staffJSONasBytes, err := json.Marshal(staff)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(staffKey, staffJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(staffJSONasBytes)

}

func (t *StaffChaincode) addRecord(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	st, err := recordByArgs(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	staffs := queryStaffIds(stub)
	if len(staffs) > 0 {
		for _, staffId := range staffs {
			if staffId == st.RecordStaffId {
				goto StaffExists
			}
		}
		fmt.Println("staff" + st.RecordStaffId + "does not exist")
		return shim.Error("staff" + st.RecordStaffId + "does not exist")
	} else {
		fmt.Println("staff" + st.RecordStaffId + "does not exist")
		return shim.Error("staff" + st.RecordStaffId + "does not exist")
	}

StaffExists:

	recordAsBytes, err := stub.GetState(st.RecordId)
	if err != nil {
		return shim.Error(err.Error())
	} else if recordAsBytes != nil {
		fmt.Println("This record already exists" + st.RecordId)
		return shim.Error("This record already exists" + st.RecordId)
	}

	recordJSONasBytes, err := json.Marshal(st)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(st.RecordId, recordJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(recordJSONasBytes)

}

func (t *StaffChaincode) deleteRecord(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1(recordId)")
	}

	recordId := args[0]

	recordAsBytes, err := stub.GetState(recordId)
	if err != nil {
		return shim.Error("Failed to get record:" + err.Error())
	} else if recordAsBytes == nil {
		return shim.Error("Record does not exist")
	}

	err = stub.DelState(recordId)
	if err != nil {
		return shim.Error("Failed to delete record :" + recordId + err.Error())
	}
	return shim.Success(nil)
}

func (t *StaffChaincode) updateRecord(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	st, err := recordByArgs(args)
	if err != nil {
		return shim.Error(err.Error())
	}
	recordId := st.RecordId

	staffs := queryStaffIds(stub)

	if len(staffs) > 0 {
		for _, staffId := range staffs {
			if staffId == st.RecordStaffId {
				goto StaffExists
			}
		}
		fmt.Println("staff " + st.RecordStaffId + " does not exist")
		return shim.Error("school " + st.RecordStaffId + " does not exist")
	} else {
		fmt.Println("staff " + st.RecordStaffId + " does not exist")
		return shim.Error("school " + st.RecordStaffId + " does not exist")
	}

StaffExists:

	recordJSONasBytes, err := json.Marshal(st)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(recordId, recordJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(recordJSONasBytes)
}

func (t *StaffChaincode) queryRecordByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1(userid)")
	}

	recordId := args[0]
	recordAsBytes, err := stub.GetState(recordId)
	if err != nil {
		return shim.Error("Failed to get student:" + err.Error())
	} else if recordAsBytes == nil {
		return shim.Error("Student does not exist")
	}
	fmt.Printf("Query Response:%s\n", string(recordAsBytes))
	return shim.Success(recordAsBytes)

}

func recordByArgs(args []string) (*Record, error) {
	if len(args) != 4 {
		return nil, errors.New("参数错误：输入7个参数（RecordId,RecordType,RecordValue,RecordStaff,RecordKita,RecordDescription,RecordTime）")
	}

	recordId := args[0]
	recordType := args[1]
	recordValue := args[2]
	recordStaffId := args[3]
	recordKitaId := args[4]
	recordDescription := args[5]
	recordTime := args[6]

	st := &Record{recordId, recordType, recordValue, recordStaffId, recordKitaId, recordDescription, recordTime}

	return st, nil
}

func queryStaffIds(stub shim.ChaincodeStubInterface) []string {
	resultsIterator, err := stub.GetStateByPartialCompositeKey("staff", []string{"staff"})
	if err != nil {
		return nil
	}

	defer resultsIterator.Close()

	scIds := make([]string, 0)

	for i := 0; resultsIterator.HasNext(); i++ {
		responseRange, err := resultsIterator.Next()
		if err != nil {
			return nil
		}

		_, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil
		}

		returnedStaffId := compositeKeyParts[1]
		scIds = append(scIds, returnedStaffId)
	}
	return scIds
}

func main() {
	if err := shim.Start(new(StaffChaincode)); err != nil {
		fmt.Println("Error staring staff chaincode: %s", err)
	}
}
