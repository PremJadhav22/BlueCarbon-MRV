package main

import (
    "encoding/json"
    "fmt"
    "log"
    
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Project represents a blue carbon project
type Project struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Area     float64 `json:"area"`
    Location string  `json:"location"`
    Owner    string  `json:"owner"`
}

// BlueCarbonContract chaincode
type BlueCarbonContract struct {
    contractapi.Contract
}

// InitLedger initializes the ledger with sample data
func (cc *BlueCarbonContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
    projects := []Project{
        {ID: "project1", Name: "Sundarbans Restoration", Area: 100.5, Location: "West Bengal", Owner: "NGO1"},
        {ID: "project2", Name: "Goa Mangrove Conservation", Area: 50.2, Location: "Goa", Owner: "Panchayat1"},
    }
    
    for _, project := range projects {
        projectJSON, err := json.Marshal(project)
        if err != nil {
            return err
        }
        
        err = ctx.GetStub().PutState(project.ID, projectJSON)
        if err != nil {
            return fmt.Errorf("failed to put to world state: %v", err)
        }
    }
    
    return nil
}

// CreateProject creates a new project
func (cc *BlueCarbonContract) CreateProject(ctx contractapi.TransactionContextInterface, 
    id string, name string, area float64, location string, owner string) error {
    
    exists, err := cc.ProjectExists(ctx, id)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("project %s already exists", id)
    }
    
    project := Project{
        ID:       id,
        Name:     name,
        Area:     area,
        Location: location,
        Owner:    owner,
    }
    
    projectJSON, err := json.Marshal(project)
    if err != nil {
        return err
    }
    
    return ctx.GetStub().PutState(id, projectJSON)
}

// GetProject returns the project with given ID
func (cc *BlueCarbonContract) GetProject(ctx contractapi.TransactionContextInterface, id string) (*Project, error) {
    projectJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if projectJSON == nil {
        return nil, fmt.Errorf("project %s does not exist", id)
    }
    
    var project Project
    err = json.Unmarshal(projectJSON, &project)
    if err != nil {
        return nil, err
    }
    
    return &project, nil
}

// ProjectExists checks if project exists
func (cc *BlueCarbonContract) ProjectExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
    projectJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return false, fmt.Errorf("failed to read from world state: %v", err)
    }
    
    return projectJSON != nil, nil
}

// GetAllProjects returns all projects
func (cc *BlueCarbonContract) GetAllProjects(ctx contractapi.TransactionContextInterface) ([]*Project, error) {
    resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()
    
    var projects []*Project
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }
        
        var project Project
        err = json.Unmarshal(queryResponse.Value, &project)
        if err != nil {
            return nil, err
        }
        projects = append(projects, &project)
    }
    
    return projects, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(&BlueCarbonContract{})
    if err != nil {
        log.Panicf("Error creating bluecarbon chaincode: %v", err)
    }
    
    if err := chaincode.Start(); err != nil {
        log.Panicf("Error starting bluecarbon chaincode: %v", err)
    }
}