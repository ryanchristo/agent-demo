package contracts

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient(client *ethclient.Client, address common.Address, abiJSON []byte) (*bind.BoundContract, error) {

	// parse contract ABI
	contractABI, err := abi.JSON(bytes.NewReader(abiJSON))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse ABI: %w", err)
	}

	// create bound contract
	contract := bind.NewBoundContract(address, contractABI, client, client, client)

	// return bound contract
	return contract, nil
}

// Call a contract method that is read-only.
func Read(contract *bind.BoundContract, methodName string, args ...any) ([]any, error) {
	ctx := context.Background()

	var result []any

	// Create call options
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// Call the contract method
	err := contract.Call(opts, &result, methodName, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to call %s: %w", methodName, err)
	}

	return result, nil
}
