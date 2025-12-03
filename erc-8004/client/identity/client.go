package identity

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	c "github.com/ryanchristo/agentic/erc-8004/client"
)

// IdentityClient wraps a bound contract for identity operations.
type IdentityClient struct {
	Contract *bind.BoundContract
}

// NewClient instantiates a new identity client.
func NewClient(ctx context.Context, client *ethclient.Client) (*IdentityClient, error) {

	// Get the identity contract address from environment.
	envContractAddress := os.Getenv("ETH_CONTRACT_IDENTITY")
	if envContractAddress == "" {
		log.Fatalf("ETH_CONTRACT_IDENTITY environment variable is not set")
	}

	// Get the contract address from the contract environment variable.
	contractAddress := common.HexToAddress(envContractAddress)

	// Get the contract ABI from the JSON file.
	contractABI, err := os.ReadFile("client/identity/abi.json")
	if err != nil {
		return nil, fmt.Errorf("Failed to read ABI: %w", err)
	}

	// Create a new contract client.
	contractClient, err := c.CreateContractClient(client, contractAddress, contractABI)
	if err != nil {
		return nil, fmt.Errorf("Failed to create contract client: %w", err)
	}

	return &IdentityClient{Contract: contractClient}, nil
}
