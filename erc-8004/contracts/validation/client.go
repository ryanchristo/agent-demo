package validation

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ryanchristo/agent-demo/erc-8004/contracts"
)

// ValidationClient wraps a bound contract for validation operations.
type ValidationClient struct {
	Contract *bind.BoundContract
}

// NewClient instantiates a new validation client.
func NewClient(ctx context.Context, client *ethclient.Client) (*ValidationClient, error) {

	// Get the identity contract address from environment.
	envContractAddress := os.Getenv("ETH_CONTRACT_VALIDATION")
	if envContractAddress == "" {
		log.Fatalf("ETH_CONTRACT_VALIDATION environment variable is not set")
	}

	// Get the contract address from the contract environment variable.
	contractAddress := common.HexToAddress(envContractAddress)

	// Get the contract ABI from the JSON file.
	contractABI, err := os.ReadFile("contracts/validation/abi.json")
	if err != nil {
		return nil, fmt.Errorf("Failed to read ABI: %w", err)
	}

	// Create a new contract client.
	contractClient, err := contracts.NewClient(client, contractAddress, contractABI)
	if err != nil {
		return nil, fmt.Errorf("Failed to create contract client: %w", err)
	}

	return &ValidationClient{Contract: contractClient}, nil
}
