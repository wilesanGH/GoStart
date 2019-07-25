package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func mockInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func initStaff(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInit(
		"1",
		[][]byte{
			[]byte("initStaff"),
			[]byte(args[0]),
			[]byte(args[1]),
			[]byte(args[2]),
			[]byte(args[3])})

	if res.Status != shim.OK {
		fmt.Println("InitStaff failed:", args[0], string(res.Message))
		t.Fail()
	}
}

func addRecord(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke(
		"1",
		[][]byte{
			[]byte("addRecord"),
			[]byte(args[0]),
			[]byte(args[1]),
			[]byte(args[2]),
			[]byte(args[3]),
			[]byte(args[4]),
			[]byte(args[5]),
			[]byte(args[6])})

	if res.Status != shim.OK {
		fmt.Println("addRecord failed:", args[0], string(res.Message))
		t.FailNow()
	}
}

func queryRecordByID(t *testing.T, stub *shim.MockStub, recordId string) {
	res := stub.MockInvoke(
		"1",
		[][]byte{
			[]byte("queryRecordByID"),
			[]byte(recordId)})

	if res.Status != shim.OK {
		fmt.Println("queryRecordByID failed:", recordId, string(res.Message))
		t.FailNow()
	}

	if res.Payload == nil {
		fmt.Println("queryRecordByID", recordId, "failed to get value")
		t.FailNow()
	}
}

func TestInitStaff(t *testing.T) {
	scc := new(StaffChaincode)
	stub := shim.NewMockStub("StaffChaincode", scc)
	mockInit(t, stub, nil)
	initStaff(t, stub, []string{"001", "liyugang", "13916526330", "normal"})
	initStaff(t, stub, []string{"002", "lishizhen", "13615256844", "normal"})
}

func TestAddRecord(t *testing.T) {
	scc := new(StaffChaincode)
	stub := shim.NewMockStub("StaffChaincode", scc)
	mockInit(t, stub, nil)
	initStaff(t, stub, []string{"001", "liyugang", "13916526330", "normal"})
	addRecord(t, stub, []string{"re001", "add", "5", "001", "kt001", "response right", "20190723"})
	//addRecord(t,stub,[]string{"re001","add","5","004","kt001","response right","20190723"})

}
