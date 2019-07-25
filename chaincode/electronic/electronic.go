package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ElectronicChaincode struct {
}

type Proposal struct {
	ProposalId      string `json:"proposal_id"`
	ProposalStatus  string `json:"proposal_status"`
	ProposalAddress string `json:"proposal_address"`
	ProposalGps     string `json:"proposal_gps"`
	ProposalAccount string `json:"proposal_user"`
	ProposalPicture string `json:"proposal_picture"`
	ProposalKita    string `json:"proposal_kita"`
	ProposalComment string `json:"proposal_comment"`
	ProposalTime    string `json:"proposal_time"`
}

type Credit struct {
	CreditId      string `json:"credit_id"`
	CreditAccount string `json:"credit_account"`
	CreditValue   string `json:"credit_value"`
	Type          string `json:"type"`
	CurrentValue  string `json:"current_value"`
}

type Score struct {
	ScoreId      string `json:"score_id"`
	ScoreAccount string `json:"credit_account"`
	ScoreValue   string `json:"score_value"`
	Type         string `json:"type"`
	CurrentValue string `json:"current_value"`
}

func (t *ElectronicChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *ElectronicChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + fn)

	switch fn {
	case "addProposal":
		return t.addProposal(stub, args)
	case "updateProposal":
		return t.updateProposal(stub, args)
	case "updateCredit":
		return t.updateCredit(stub, args)
	case "updateSocre":
		return t.updateSocre(stub, args)
	case "queryProposal":
		return t.queryProposal(stub, args)
	case "queryCredit":
		return t.queryCredit(stub, args)
	case "queryScore":
		return t.queryScore(stub, args)

	}
}

func (t *ElectronicChaincode) addProposal(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

func (t *ElectronicChaincode) updateProposal(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

func (t *ElectronicChaincode) updateCredit(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

func (t *ElectronicChaincode) updateSocre(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

func (t *ElectronicChaincode) queryProposal(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

func (t *ElectronicChaincode) queryCredit(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}

func (t *ElectronicChaincode) queryScore(stub shim.ChaincodeStubInterface, args []string) peer.Response {

}
