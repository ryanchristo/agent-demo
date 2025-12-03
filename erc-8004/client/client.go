package client

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CreateContractClient(client *ethclient.Client, address common.Address, abiJSON []byte) (*bind.BoundContract, error) {

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
