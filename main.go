package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (e *SmartContract) Init(ctx contractapi.TransactionContext) error {
	// TODO: insert key
	return ctx.GetStub().PutState("mykey", []byte("myvalure"))
}

func (s *SmartContract) Invoke(ctx contractapi.TransactionContext) error {
	function, args := ctx.GetStub().GetFunctionAndParameters()

	switch function {
	case "write":
		return s.write(ctx, args)
	case "read":
		return s.read(ctx, args)
	default:
		return fmt.Errorf("Unknown function %s", function)
	}
}

func (s *SmartContract) write(ctx contractapi.TransactionContext, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("incorrect number of arguments. Expecting 2")
	}

	key := args[0]
	value := args[1]
	return ctx.GetStub().PutState(key, []byte(value))
}

func (s *SmartContract) read(ctx contractapi.TransactionContext, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("incorrect number of arguments. Expecting 1")
	}

	key := args[0]
	value, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}

	if value == nil {
		return fmt.Errorf("key %s does not exist", key)
	}

	fmt.Printf("Query Result: Key=%s, Value=%s\n", key, value)
	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		panic(err)
	}

	if err := chaincode.Start(); err != nil {
		panic(err)
	}
}
