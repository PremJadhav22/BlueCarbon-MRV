package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Project struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Area string `json:"area"`
}

type SimpleContract struct {
    contractapi.Contract
}

func (s *SimpleContract) Init(ctx contractapi.TransactionContextInterface) error {
    fmt.Println("Chaincode initialized")
    return nil
}

func (s *SimpleContract) CreateProject(ctx contractapi.TransactionContextInterface, id string, name string, area string) error {
    project := Project{
        ID:   id,
        Name: name,
        Area: area,
    }
    
    projectJSON, _ := json.Marshal(project)
    return ctx.GetStub().PutState(id, projectJSON)
}

func (s *SimpleContract) GetProject(ctx contractapi.TransactionContextInterface, id string) (*Project, error) {
    projectJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("failed to read: %v", err)
    }
    if projectJSON == nil {
        return nil, fmt.Errorf("project not found")
    }
    
    var project Project
    json.Unmarshal(projectJSON, &project)
    return &project, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(&SimpleContract{})
    if err != nil {
        panic(err)
    }
    
    if err := chaincode.Start(); err != nil {
        panic(err)
    }
}