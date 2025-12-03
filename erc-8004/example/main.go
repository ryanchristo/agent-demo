package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ryanchristo/agentic/erc-8004/client/identity"
	"github.com/ryanchristo/agentic/erc-8004/client/reputation"
	"github.com/ryanchristo/agentic/erc-8004/client/validation"
)

func main() {

	// Load environment variables.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Get RPC environment variable.
	ethRpcUrl := os.Getenv("ETH_RPC_URL")
	if ethRpcUrl == "" {
		log.Fatal("ETH_RPC_URL environment variable not set")
	}

	// Log connecting to RPC endpoint.
	fmt.Println("Connecting to node RPC endpoint...")

	// Connect ethereum client to RPC endpoint.
	client, err := ethclient.Dial(ethRpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to RPC endpoint: %v", err)
	}

	// Close ethereum client when program exits.
	defer client.Close()

	// Log creating contract clients.
	fmt.Println("Creating ERC-8004 contract clients...")

	// Create identity contract client.
	identityClient, err := identity.NewClient(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to create identity client: %v", err)
	}

	// Create reputation contract client.
	reputationClient, err := reputation.NewClient(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to create reputation client: %v", err)
	}

	// Create validation contract client.
	validationClient, err := validation.NewClient(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to create validation client: %v", err)
	}

	// Log calling contract clients.
	fmt.Println("Calling ERC-8004 contract clients...")

	// Get identity contract version.
	identityVersion, err := readContractMethod(identityClient.Contract, "getVersion")
	if err != nil {
		log.Fatalf("Failed to call contract method getVersion: %v", err)
	} else {
		fmt.Printf("Version: %s\n", identityVersion)
	}

	// Get reputation contract version.
	reputationVersion, err := readContractMethod(reputationClient.Contract, "getVersion")
	if err != nil {
		log.Fatalf("Failed to call contract method getVersion: %v", err)
	} else {
		fmt.Printf("Version: %s\n", reputationVersion)
	}

	// Get validation contract version.
	validationVersion, err := readContractMethod(validationClient.Contract, "getVersion")
	if err != nil {
		log.Fatalf("Failed to call contract method getVersion: %v", err)
	} else {
		fmt.Printf("Version: %s\n", validationVersion)
	}
}

// Call a contract method that is read-only.
func readContractMethod(contract *bind.BoundContract, methodName string, args ...any) ([]any, error) {
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
